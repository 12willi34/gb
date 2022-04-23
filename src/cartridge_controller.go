package gb

import "fmt"

//todo: polymorphism

type Cartridge struct {
  mode int

  rom []uint8
  romBank uint32
  romBanking bool

  ram []byte
  ramBank uint32
  ramEnabled bool
}

func NewCartridge(rom []uint8, mode uint8) Cartridge {
  switch mode {
  case 0x00, 0x08, 0x09, 0x0b, 0x0c, 0x0d:
    return Cartridge {
      mode: 0,
      rom: rom,
      ram: make([]uint8, 0x2000),
    }
  case 0x01, 0x02, 0x03:
    return Cartridge {
      mode: 1,
      rom: rom,
      romBank: 1,
      ram: make([]uint8, 0x8000),
    }
  case 0x04, 0x05, 0x06:
    return Cartridge {
      mode: 2,
      rom: rom,
    }
  default:
    fmt.Println("cartridge mode not supported:", mode)
    return Cartridge {
      mode: 0,
      rom: rom,
    }
  }
}

func (this *Cartridge) Read(i uint16) uint8 {
  switch this.mode {
  case 0:
    return this.Read_mode0(i)
  case 1:
    return this.Read_mode1(i)
  case 2:
    return this.Read_mode2(i)
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
  return this.rom[i - 0x4000]
}

func (this *Cartridge) Write_mode2(i uint16, d uint8) {
  return
}
