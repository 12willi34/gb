package gb

import "fmt"

type Cartridge interface {
  Read(i uint16) uint8
  Write(i uint16, d uint8)
}

type Cartridge_0 struct {
  rom [16384]uint8
}

type Cartridge_1 struct {
  rom [16384]uint8
}

type Cartridge_2 struct {
  rom [16384]uint8
}

func NewCartridge(rom [16384]uint8, mode uint8) Cartridge {
  println(mode)
  switch mode {
  case 0x00, 0x08, 0x09, 0x0b, 0x0c, 0x0d:
    return Cartridge_0 {
      rom: rom,
    }
  case 0x01, 0x02, 0x03:
    return Cartridge_1 {
      rom: rom,
    }
  case 0x04, 0x05, 0x06:
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

func (this Cartridge_0) Read(i uint16) uint8 {
  return this.rom[i - 0x4000]
}

func (this Cartridge_0) Write(i uint16, d uint8) {
  return
}

func (this Cartridge_1) Read(i uint16) uint8 {
  return this.rom[i - 0x4000]
}

func (this Cartridge_1) Write(i uint16, d uint8) {
  return
}

func (this Cartridge_2) Read(i uint16) uint8 {
  return this.rom[i - 0x4000]
}

func (this Cartridge_2) Write(i uint16, d uint8) {
  return
}
