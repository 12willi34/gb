package gb

import ()

type memoryunit struct {
	addr []uint8
  Processor cpu
}

func NewMemoryUnit(boot []byte, rom []byte) memoryunit {
	mu := memoryunit {
		addr: make([]uint8, 0x10000),
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
  for i := 0; i < len(rom); i++ {
    mu.addr[i] = rom[i]
  }
  for i := 0; i < len(boot); i++ {
    mu.addr[i] = boot[i]
  }
  return mu
}

func (this memoryunit) Read_8(i uint16) uint8 {
  if(i >= 0 && i < 0x8000) {
	  return this.addr[i]
  }
  if(i >= 0xe000 && i <= 0xfdff) {
    //println("reading from mirrored work ram")
	  return this.addr[i - 0x2000]
  }
  if(i >= 0xfea0 && i <= 0xfeff) {
    //println("reading from unusable memory")
    return 0xff
  }
	return this.addr[i]
}

func (this memoryunit) Write_8(i uint16, data uint8) {
  if(i < 0x8000) {
    return
  } else if((i >= 0xe000) && (i <= 0xfdff)) {
    //println("writing to mirrored work ram")
    this.Write_8(i - 0x2000, data)
    return
  } else if((i >= 0xfea0) && (i <= 0xfeff)) {
    //println("writing to unusable memory")
    return
  } else if(i == 0xff46) {
    this.dma(data)
  } else if(i == 0xff44) {
    this.addr[0xff44] = 0
    return
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

func (this memoryunit) dma(data uint8) {
  var src uint16 = uint16(data)*uint16(0x100)
  for i := uint8(0); i < 0x9f; i++ {
    val := this.Read_8(src + uint16(i))
    this.Write_8(0xfe00 + uint16(i), val)
  }
}
