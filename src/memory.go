package gb

import ()

type memoryunit struct {
	addr []uint8
  dma_val uint8
  dma_status bool
}

func NewMemoryUnit() memoryunit {
	mu := memoryunit {
		addr: make([]uint8, 0x10000),
    dma_val: 0,
    dma_status: false,
	}
  /*
  mu.addr[0x04] = 0x1e
  mu.addr[0x05] = 0x00
  mu.addr[0x06] = 0x00
  mu.addr[0x07] = 0xf8
  mu.addr[0x0F] = 0xe1
  mu.addr[0x10] = 0x80
  mu.addr[0x11] = 0xbf
  mu.addr[0x12] = 0xf3
  mu.addr[0x14] = 0xbf
  mu.addr[0x16] = 0x3f
  mu.addr[0x17] = 0x00
  mu.addr[0x19] = 0xbf
  mu.addr[0x1a] = 0x7f
  mu.addr[0x1b] = 0xff
  mu.addr[0x1c] = 0x9f
  mu.addr[0x1e] = 0xbf
  mu.addr[0x20] = 0xff
  mu.addr[0x21] = 0x00
  mu.addr[0x22] = 0x00
  mu.addr[0x23] = 0xbf
  mu.addr[0x24] = 0x77
  mu.addr[0x25] = 0xf3
  mu.addr[0x26] = 0xf1
  mu.addr[0x40] = 0x91
  mu.addr[0x41] = 0x85
  mu.addr[0x42] = 0x00
  mu.addr[0x43] = 0x00
  mu.addr[0x45] = 0x00
  mu.addr[0x47] = 0xfc
  mu.addr[0x48] = 0xff
  mu.addr[0x49] = 0xff
  mu.addr[0x4a] = 0x00
  mu.addr[0x4b] = 0x00
  mu.addr[0xff] = 0x00
  */
  return mu
}

func (this memoryunit) Read_8(i uint16) uint8 {
	return this.addr[i]
}

func (this memoryunit) Write_8(i uint16, data uint8) {
  if(i == 0xff46) {
    this.dma_val = data
    this.dma_status = true
  }
	this.addr[i] = data
}

func (this memoryunit) Read_16(i uint16) uint16 {
	return uint16((uint16(this.addr[i])) | uint16(this.addr[i + 1]) << 8)
}

func (this memoryunit) Write_16(i uint16, data uint16) {
	this.addr[i] = uint8(data & 0xFF)
	this.addr[i + 1] = uint8(data >> 8)
}
