package gb

//sources
//https://www.linkedin.com/pulse/creating-gameboy-emulator-part-1-bruno-croci
//http://gameboy.mongenel.com/dmg/asmmemmap.html

import (
	//"strconv";
	//"fmt";
)

type memoryunit struct {
	bios []uint8
	addr []uint8
}

func NewMemoryUnit() memoryunit {
	mmu := memoryunit {
		bios: make([]uint8, 0x100, 0x100),
		addr: make([]uint8, 0x10000, 0x10000),
	}
	return mmu
}

func (mmu memoryunit) Finished_bios() bool {
	return mmu.addr[0xFF50] == 1
}

func (mmu memoryunit) Read_8(i uint16) uint8 {
	if(!mmu.Finished_bios() && i <= 0xFF) {
		return mmu.bios[i]
	}
	return mmu.addr[i]
}

func (mmu memoryunit) Write_8(i uint16, data uint8) {
	mmu.addr[i] = data
}

func (mmu memoryunit) Read_16(i uint16) uint16 {
	var x []uint8 = mmu.addr
	if(!mmu.Finished_bios() && i <= 0xFF) { x = mmu.bios }
	return uint16((uint16(x[i])) | uint16(x[i + 1]) << 8)
}

func (mmu memoryunit) Write_16(i uint16, data uint16) {
	mmu.addr[i] = uint8(data & 0xFF)
	mmu.addr[i + 1] = uint8(data >> 8)
}
