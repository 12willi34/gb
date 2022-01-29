package gb

//sources
//https://www.linkedin.com/pulse/creating-gameboy-emulator-part-1-bruno-croci
//http://gameboy.mongenel.com/dmg/asmmemmap.html

type memoryunit struct {
	bios []uint8
	addr [0x10000]uint8
}

func NewMemoryUnit() memoryunit {
	mmu := memoryunit {
		bios: make([]uint8, 0x100, 0x100),
		addr: make([]uint8, 0x10000, 0x10000),
	}
}

func (mmu memoryunit) finished_bios() bool {
	return mmu.addr[0xFF50] == 1
}

func (mmu memoryunit) read_8(uint16 i) uint8 {
	if(!mmu.finished_bios() && i <= 0xFF) {
		return mmu.bios[i]
	}
	return mmu.addr[i]
}

func (mmu memoryunit) write_8(uint16 i, uint8 data) {
	mmu.addr[i] = data
}

func (mmu memoryunit) read_16(uint16 i) uint16 {
	if(!mmu.finished_bios() && i <= 0xFF) {
		return (mmu.bios[i] << 8) | mmu.bios[i + 1]
	}
	return (mmu.addr[i] << 8) & mmu.addr[i + 1]
}

func (mmu memoryunit) write_16(uint16 i, uint16 data) {
	mmu.addr[i] = 0 | (data >> 8)
	mmu.addr[i + 1] = data &^ 0
}
