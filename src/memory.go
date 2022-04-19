package gb

type memoryunit struct {
	addr []uint8
  Processor *cpu
  Io *io_controller
  cart *Cartridge
}

func NewMemoryUnit(boot [256]byte, rom []byte) memoryunit {
  io := NewIoController()
	mu := memoryunit {
		addr: make([]uint8, 0x10000),
    Io: &io,
	}
  for i := 0; i < 0x100 && i < len(boot); i++ {
    mu.addr[i] = boot[i]
  }
  for i := 0x100; i < 0x4000 && i < len(rom); i++ {
    mu.addr[i] = uint8(rom[i])
  }
  cartRom := [16384]uint8{}
  for i := 0x4000; i < 0x7ffff && i < len(rom); i++ {
    cartRom[i - 0x4000] = uint8(rom[i])
  }
  cart := NewCartridge(cartRom)
  mu.cart = &cart
  return mu
}

func (this *memoryunit) Read_8(i uint16) uint8 {
  if(i >= 0 && i < 0x4000) {
	  return this.addr[i]
  }
  if(i >= 0x4000 && i <= 0x7fff) {
    return (*this.cart).Read(i)
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

func (this *memoryunit) Write_8(i uint16, data uint8) {
  if(i < 0x8000) {
    (*this.cart).Write(i, data)
    return
  } else if(i < 0xe000) {
	  this.addr[i] = data
    return
  } else if((i >= 0xe000) && (i <= 0xfdff)) {
    this.Write_8(i - 0x2000, data)
    return
  } else if((i >= 0xfea0) && (i <= 0xfeff)) {
    return
  } else if(i == 0xff00) {
    this.Io.ChangeMode(data)
    return
  } else if(i == 0xff02) {
    if(data == 0x81) {
      //serial transfer requested
      // -> always rejected
      this.addr[0xff01] = 0xff
      this.addr[0xff02] = 0x01
      this.addr[0xff0f] |= (1 << 3)
      return
    }
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

func (this *memoryunit) Read_16(i uint16) uint16 {
	return uint16((uint16(this.addr[i])) | uint16(this.addr[i + 1]) << 8)
}

func (this *memoryunit) Write_16(i uint16, data uint16) {
	this.addr[i] = uint8(data & 0xFF)
	this.addr[i + 1] = uint8(data >> 8)
}

func (this *memoryunit) read_io(i uint16) uint8 {
  switch i {
  case 0xff00:
    return this.Io.Get()
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

func (this *memoryunit) dma(data uint8) {
  var src uint16 = uint16(data)*uint16(0x100)
  for i := uint8(0); i < 0x9f; i++ {
    val := this.Read_8(src + uint16(i))
    this.Write_8(0xfe00 + uint16(i), val)
  }
}

func (this *memoryunit) SetBtn(keyCode int, state bool) bool {
  res := this.Io.Set(keyCode, state)
  this.addr[0xff00] = this.Io.Get()
  return res
}
