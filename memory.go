package gb

//https://www.linkedin.com/pulse/creating-gameboy-emulator-part-1-bruno-croci
//http://gameboy.mongenel.com/dmg/asmmemmap.html
//https://raphaelstaebler.medium.com/memory-and-memory-mapped-i-o-of-the-gameboy-part-3-of-a-series-37025b40d89b
//memory layout: http://bgb.bircd.org/pandocs.htm#memorymap

import ()

type memoryunit struct {
	addr []uint8
}

func NewMemoryUnit() memoryunit {
	return memoryunit {
		addr: make([]uint8, 0x10000),
	}
}

func (this memoryunit) Read_8(i uint16) uint8 {
	return this.addr[i]
}

func (this memoryunit) Write_8(i uint16, data uint8) {
	this.addr[i] = data
}

func (this memoryunit) Read_16(i uint16) uint16 {
	return uint16((uint16(this.addr[i])) | uint16(this.addr[i + 1]) << 8)
}

func (this memoryunit) Write_16(i uint16, data uint16) {
	this.addr[i] = uint8(data & 0xFF)
	this.addr[i + 1] = uint8(data >> 8)
}
