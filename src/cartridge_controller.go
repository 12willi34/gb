package gb

type Cartridge interface {
  Read(i uint16) uint8
  Write(i uint16, d uint8)
}

type ZeroCartridge struct {
  rom [16384]uint8
}

func NewCartridge(rom [16384]uint8) Cartridge {
  return ZeroCartridge{
    rom: rom,
  }
}

func (this ZeroCartridge) Read(i uint16) uint8 {
  return this.rom[i - 0x4000]
}

func (this ZeroCartridge) Write(i uint16, d uint8) {
  return
}
