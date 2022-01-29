package gb

//sources
//https://www.linkedin.com/pulse/creating-gameboy-emulator-part-1-bruno-croci
//http://gameboy.mongenel.com/dmg/asmmemmap.html

import (
	"strconv"
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

func (mmu memoryunit) finished_bios() bool {
	return mmu.addr[0xFF50] == 1
}

func (mmu memoryunit) read_8(i uint16) uint8 {
	if(!mmu.finished_bios() && i <= 0xFF) {
		return mmu.bios[i]
	}
	return mmu.addr[i]
}

func (mmu memoryunit) write_8(i uint16, data uint8) {
	mmu.addr[i] = data
}

func (mmu memoryunit) read_16(i uint16) uint16 {
	if(!mmu.finished_bios() && i <= 0xFF) {
		var left string = strconv.FormatInt(int64(mmu.bios[i]), 2)
		var right string = strconv.FormatInt(int64(mmu.bios[i + 1]), 2)
		res, _ := strconv.ParseInt(left + right, 2, 16)
		return uint16(res)
	}
	var left string = strconv.FormatInt(int64(mmu.addr[i]), 2)
	var right string = strconv.FormatInt(int64(mmu.addr[i + 1]), 2)
	res, _ := strconv.ParseInt(left + right, 2, 16)
	return uint16(res)
}

func (mmu memoryunit) write_16(i uint16, data uint16) {
	mmu.addr[i] = uint8((0) | (data >> 8))
	mmu.addr[i + 1] = uint8(data & uint16(0xFF))
}
