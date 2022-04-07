package gb

//import "fmt"

type memoryunit struct {
	addr []uint8
  Processor *cpu
  io io_controller
}

func NewMemoryUnit(boot [0x100]byte, rom []byte, io io_controller) memoryunit {
	mu := memoryunit {
		addr: make([]uint8, 0x10000),
    io: io,
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
  for i := 0; i < 0x8000 && i < len(rom); i++ {
    mu.addr[i] = rom[i]
  }
  for i := 0; i < 0x100 && i < len(boot); i++ {
    mu.addr[i] = boot[i]
  }
  return mu
}

func (this memoryunit) Read_8(i uint16) uint8 {
  if(i >= 0 && i < 0x8000) {
	  return this.addr[i]
  }
  if(i >= 0xe000 && i <= 0xfdff) {
	  return this.Read_8(i - 0x2000)
  }
  if(i >= 0xfea0 && i <= 0xfeff) {
    return 0xff
  }
  if(i >= 0xff00 && i <= 0xff7f) {
    return this.read_io(i)
  }
	return this.addr[i]
}

func (this memoryunit) Write_8(i uint16, data uint8) {
  if(i < 0x8000) {
    return
  } else if((i >= 0xe000) && (i <= 0xfdff)) {
    this.Write_8(i - 0x2000, data)
    return
  } else if((i >= 0xfea0) && (i <= 0xfeff)) {
    return
  } else if(i == 0xff00) {
    this.addr[i] = this.io.ChangeMode(data & 0b00110000)
    return
  } else if(i == 0xff04) {
    this.addr[0xff04] = 0
  } else if(i == 0xff05) {
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

func (this memoryunit) read_io(i uint16) uint8 {
  switch i {
  case 0xff00:
    x := this.io.Get()
    //fmt.Printf("%b\n", x)
    return x
  case 0xff02:
    return 0xff
  case 0xff10:
    fallthrough
  case 0xff11:
    fallthrough
  case 0xff12:
    fallthrough
  case 0xff13:
    fallthrough
  case 0xff14:
    fallthrough
  case 0xff16:
    fallthrough
  case 0xff17:
    fallthrough
  case 0xff18:
    fallthrough
  case 0xff19:
    fallthrough
  case 0xff1a:
    fallthrough
  case 0xff1b:
    fallthrough
  case 0xff1c:
    fallthrough
  case 0xff1d:
    fallthrough
  case 0xff1e:
    fallthrough
  case 0xff20:
    fallthrough
  case 0xff21:
    fallthrough
  case 0xff22:
    fallthrough
  case 0xff23:
    fallthrough
  case 0xff24:
    fallthrough
  case 0xff25:
    fallthrough
  case 0xff26:
    return 0xff
  case 0xff4d:
    return 0
  default:
    return this.addr[i]
  }
}

func (this memoryunit) dma(data uint8) {
  var src uint16 = uint16(data)*uint16(0x100)
  for i := uint8(0); i < 0x9f; i++ {
    val := this.Read_8(src + uint16(i))
    this.Write_8(0xfe00 + uint16(i), val)
  }
}
