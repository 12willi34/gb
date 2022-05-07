package gb

import "os"
import "time"

//todo: polymorphism

type Cartridge struct {
  mode int

  rom []uint8
  romBank uint32
  romBanking bool

  ram []byte
  ramBank uint32
  ramEnabled bool

  rtc []uint8
  latchedRtc []uint8
  latched bool
}

func NewCartridge(game string) (Cartridge, []byte) {
  cartBytes, err := os.ReadFile(game)
  if(err != nil) { panic(err) }
  rom := make([]uint8, len(cartBytes))
  for i := 0; i < len(cartBytes); i++ {
    rom[i] = uint8(cartBytes[i])
  }
  var mode int
  battery := false
  switch rom[0x147] {
  case 0x00, 0x08, 0x0b, 0x0c:
    mode = 0
  case 0x09, 0x0d:
    mode = 0
    battery = true
  case 0x01, 0x02:
    mode = 1
  case 0x03:
    mode = 1
    battery = true
  case 0x05:
    mode = 2
  case 0x06:
    mode = 2
    battery = true
  case 0x11, 0x12:
    mode = 3
  case 0xf, 0x10, 0x13:
    mode = 3
    battery = true
  default:
    println("cartridge not supported")
    mode = 0
  }
  var cart Cartridge
  switch mode {
    case 0:
      cart = Cartridge {
        mode: 0,
        rom: rom,
        ram: make([]uint8, 0x2000),
      }
    case 2:
      cart = Cartridge {
        mode: 2,
        rom: rom,
        romBank: 1,
        ram: make([]uint8, 0x2000),
      }
    case 1:
      cart = Cartridge {
        mode: 1,
        rom: rom,
        romBank: 1,
        ram: make([]uint8, 0x8000),
      }
    case 3:
      cart = Cartridge {
        mode: 3,
        rom: rom,
        romBank: 1,
        ram: make([]uint8, 0x8000),
        rtc: make([]uint8, 0x10),
        latchedRtc: make([]uint8, 0x10),
      }
  }
  if(battery) {
    cart.Load(game)
    ticker := time.NewTicker(time.Second)
    go func() {
      for range ticker.C { cart.Save(game) }
    }()
  }
  return cart, cartBytes[:256]
}

func (this *Cartridge) Read(i uint16) uint8 {
  switch this.mode {
  case 0:
    return this.Read_mode0(i)
  case 1:
    return this.Read_mode1(i)
  case 2:
    return this.Read_mode2(i)
  case 3:
    return this.Read_mode3(i)
  default:
    return 0
  }
}

func (this *Cartridge) Write(i uint16, d uint8) {
  switch this.mode {
  case 0:
    this.Write_mode0(i, d)
  case 1:
    this.Write_mode1(i, d)
  case 2:
    this.Write_mode2(i, d)
  case 3:
    this.Write_mode3(i, d)
  }
}

//mode 0

func (this *Cartridge) Read_mode0(i uint16) uint8 {
  if(i < 0x8000) {
    return this.rom[i]
  }
  return this.ram[i]
}

func (this *Cartridge) Write_mode0(i uint16, d uint8) {
  if(i > 0x8000) {
    this.ram[i - 0xa000] = d
  }
}

//mode 1

func (this *Cartridge) Read_mode1(i uint16) uint8 {
  if(i < 0x4000) {
    return this.rom[i]
  }
  if(i < 0x8000) {
    return this.rom[uint32(i) + (this.romBank - 1)*0x4000]
  }
  return this.ram[this.ramBank*0x2000 + uint32(i - 0xa000)]
}

func (this *Cartridge) Write_mode1(i uint16, d uint8) {
  switch {
  case i < 0x2000:
    t := d & 0xf
    if(t == 0xa) {
      this.ramEnabled = true
    } else if(t == 0) {
      this.ramEnabled = false
    }
  case i < 0x4000:
    this.romBank = (this.romBank & 0xe0) | uint32(d & 0x1f)
    this.updateRomBankIfZero()
  case i < 0x6000:
    if this.romBanking {
      this.romBank = (this.romBank & 0x1f) | uint32(d & 0xe0)
      this.updateRomBankIfZero()
    } else {
      this.ramBank = uint32(d & 0x3)
    }
  case i < 0x8000:
    this.romBanking = (d & 0x1 == 0)
    if this.romBanking {
      this.ramBank = 0
    } else {
      this.romBank = this.romBank & 0x1f
    }
  case i >= 0xa000 && i <= 0xbfff:
    if this.ramEnabled {
      this.ram[(0x2000)*this.ramBank + uint32(i - 0xa000)] = d
    }
  }
}

func (this *Cartridge) updateRomBankIfZero() {
  if this.romBank == 0 || this.romBank == 0x20 || this.romBank == 0x40 || this.romBank == 0x60 {
    this.romBank++
  }
}

//mode 2

func (this *Cartridge) Read_mode2(i uint16) uint8 {
  if(i < 0x4000) {
    return this.rom[i]
  }
  if(i < 0x8000) {
    return this.rom[uint32(i) + (this.romBank - 1)*0x4000]
  }
  return this.ram[i - 0xa000]
}

func (this *Cartridge) Write_mode2(i uint16, d uint8) {
  switch {
  case i < 0x2000:
    if i & 0x100 == 0 {
      if d & 0xf == 0xa {
        this.ramEnabled = true
      } else if d & 0xf == 0 {
        this.ramEnabled = false
      }
    }
  case i < 0x4000:
    if i & 0x100 != 0x100 { return }
    this.romBank = uint32(d & 0xf)
    this.updateRomBankIfZero()
  case i >= 0xa000 && i <= 0xbfff:
    if this.ramEnabled {
      this.ram[i - 0xa000] = d & 0xf
    }
  }
}

//mode 3

func (this *Cartridge) Read_mode3(i uint16) uint8 {
  if(i < 0x4000) {
    return this.rom[i]
  }
  if(i < 0x8000) {
    return this.rom[uint32(i) + (this.romBank - 1)*0x4000]
  }
  if(this.ramBank >= 0x4) {
    if(this.latched) {
      return this.latchedRtc[this.ramBank]
    }
    return this.rtc[this.ramBank]
  }
  return this.ram[this.ramBank*0x2000 + uint32(i - 0xa000)]
}

func (this *Cartridge) Write_mode3(i uint16, d uint8) {
  switch {
  case i < 0x2000:
    this.ramEnabled = d & 0xa != 0
  case i < 0x4000:
    this.romBank = uint32(d & 0x7f)
    if(this.romBank == 0) { this.romBank++ }
  case i < 0x6000:
    this.ramBank = uint32(d)
  case i < 0x8000:
    if(d == 1) {
      this.latched = false
    } else if(d == 0) {
      this.latched = true
      copy(this.rtc, this.latchedRtc)
    }
  case i >= 0xa000 && i <= 0xbfff:
    if(!this.ramEnabled) { return }
    if(this.ramBank >= 0x4) {
      this.rtc[this.ramBank] = d
    } else {
      this.ram[this.ramBank*0x2000 + uint32(i - 0xa000)] = d
    }
  }
}

func (this *Cartridge) Save(filename string) {
  data := make([]byte, len(this.ram))
  for i := 0; i < len(this.ram); i++ {
    data[i] = byte(this.ram[i])
  }
  f, err := os.OpenFile(filename + ".dump", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
  if(err != nil) { panic(err) }
  defer f.Close()
  _, err = f.Write(data)
  if(err != nil) { panic(err) }
}

func (this *Cartridge) Load(filename string) {
  ram, err := os.ReadFile(filename + ".dump")
  if(err != nil) {
    println("no data loaded")
  } else {
    this.ram = ram
  }
}
