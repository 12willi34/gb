package gb

import "fmt"

type Cartridge interface {
  Read(i uint16) uint8
  Write(i uint16, d uint8)
}

//MBC 0

type Cartridge_0 struct {
  rom []uint8
  ram []uint8
}

func (this Cartridge_0) Read(i uint16) uint8 {
  if(i < 0x8000) {
    return this.rom[i]
  }
  return this.ram[i]
}

func (this Cartridge_0) Write(i uint16, d uint8) {
  if(i > 0x8000) {
    this.ram[i - 0xa000] = d
  }
}

//MBC 1

type Cartridge_1 struct {
  rom []uint8
  romBank uint32
  romBanking bool

  ram []byte
  ramBank uint32
  ramEnabled bool
}


func (this Cartridge_1) Read(i uint16) uint8 {
  if(i < 0x4000) {
    return this.rom[i]
  }
  if(i < 0x8000) {
    return this.rom[uint32(i - 0x4000) + this.romBank*0x4000]
  }
  return this.ram[this.ramBank*0x2000 + uint32(i - 0xa000)]
}

func (this Cartridge_1) Write(i uint16, d uint8) {
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

func (this Cartridge_1) updateRomBankIfZero() {
  if this.romBank == 0 || this.romBank == 0x20 || this.romBank == 0x40 || this.romBank == 0x60 {
    this.romBank++
  }
}

//MBC 2

type Cartridge_2 struct {
  rom []uint8
}

func (this Cartridge_2) Read(i uint16) uint8 {
  return this.rom[i - 0x4000]
}

func (this Cartridge_2) Write(i uint16, d uint8) {
  return
}

func NewCartridge(rom []uint8, mode uint8) Cartridge {
  print("cartridge mode: ")
  switch mode {
  case 0x00, 0x08, 0x09, 0x0b, 0x0c, 0x0d:
    print("0")
    println()
    return Cartridge_0 {
      rom: rom,
      ram: make([]uint8, 0x2000),
    }
  case 0x01, 0x02, 0x03:
    print("1")
    println()
    return Cartridge_1 {
      rom: rom,
      romBank: 1,
      ram: make([]uint8, 0x8000),
    }
  case 0x04, 0x05, 0x06:
    print("2")
    println()
    return Cartridge_2 {
      rom: rom,
    }
  default:
    fmt.Println("cartridge mode not supported:", mode)
    return Cartridge_0 {
      rom: rom,
    }
  }
}
