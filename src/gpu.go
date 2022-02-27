package gb

import (
  "fmt"
)

const width = 160
const height = 144
const hblank_mode = uint8(0)
const vblank_mode = uint8(1)
const oam_mode = uint8(2)
const data_mode = uint8(3)

type Gpu struct {
  mu *memoryunit
  ir interrupter
  clock int
  buffer [height][width]uint8
  vblank bool

  //registers
  line Line
  lcdc Lcdc
  stat Stat
  scroll Scroll
  window Window
  palette Palette
  palette0 ObjPalette0
  palette1 ObjPalette1
}

func NewGpu(mu memoryunit, interrupter interrupter) Gpu {
  x := Gpu {
    mu: &mu,
    ir: interrupter,
    clock: 0,
    buffer: [height][width]uint8{},
    vblank: false,

    //registers
    lcdc: NewLcdc(mu),
    stat: NewStat(mu),
    scroll: NewScroll(mu),
    window: NewWindow(mu),
    palette: NewPalette(mu),
    palette0: NewObjPalette0(mu),
    palette1: NewObjPalette1(mu),
    line: NewLine(mu),
  }
  x.lcdc.set(x.lcdc.get() | (1 << 7)) //enable lcd
  return x
}

func (this *Gpu) clearScreen() {
  //todo
}

func (this *Gpu) scanline(line uint8) {
  ctrl := (*this).lcdc.get()
  if((ctrl & (1 << 0)) > 0) {
    this.renderTiles()
  }
  if((ctrl & (1 << 1)) > 0) {
    this.renderSprites()
  }
}

/**
 * bit 3 - background tile map area
 *  0: 0x9800 - 0x9bff
 *  1: 0x9c00 - 0x9fff
 * bit 4 - tile data area
 *  0: 0x8800 - 0x97ff
 *  1: 0x8000 - 0x8fff
 * bit 6 - window tile map area
 *  0: 0x9800 - 0x9bff
 *  1: 0x9c00 - 0x9fff
 **/
func (this *Gpu) renderTiles() {
  sx := (*this).scroll.get_x()
  sy := (*this).scroll.get_y()
  wx := (*this).window.get_x() - 7
  wy := (*this).window.get_y()
  in_window := (((*this).lcdc.get() & (1 << 5)) > 0) && (wy <= (*this).line.get())

  var tile_start uint16 = 0x8800
  unsigned_tile_address := false
  if(((*this).lcdc.get() & (1 << 4)) > 0) {
    tile_start = 0x8000
    unsigned_tile_address = true
  }

  var background_start uint16
  var ctrl_bit uint8
  if(!in_window) {
    ctrl_bit = 3
  } else {
    ctrl_bit = 6
  }
  if(((*this).lcdc.get() & (1 << ctrl_bit)) > 0) {
    background_start = 0x9c00
  } else {
    background_start = 0x9800
  }

  var yPos uint8
  var xPos uint8
  if(!in_window) {
    yPos = sy + (*this).line.get()
  } else {
    yPos = (*this).line.get() - wy
  }
  row := uint16(yPos/8)*32
  for x := uint8(0); x < 160; x++ {
    xPos = x + sx
    if(in_window && x >= wx) {
      xPos = x - wx
    }
    col := uint16(xPos/8)
    var tile_num int16
    var tile_adr uint16 = background_start + row + col
    var tile_loc uint16 = tile_start
    if(unsigned_tile_address) {
      tile_num = int16(uint8((*(*this).mu).Read_8(tile_adr)))
      tile_loc += uint16(tile_num*16)
    } else {
      tile_num = int16(int8((*(*this).mu).Read_8(tile_adr)))
      tile_loc = uint16(int32(tile_loc) + int32((tile_num + 128)*16))
    }

    line := (yPos % 8)*2
    val0 := (*(*this).mu).Read_8(tile_loc + uint16(line))
    val1 := (*(*this).mu).Read_8(tile_loc + uint16(line) + uint16(1))

    //nicht sicher ob das richtig ist v2
    var colour_bit uint8 = uint8(int8((xPos % 8) - 7)*-1)

    colour := uint8(0)
    if((val0 & (1 << colour_bit)) > 0) {
      colour |= 1 << 1
    }
    if((val1 & (1 << colour_bit)) > 0) {
      colour |= 1
    }
    (*this).buffer[(*this).line.get()][x] = this.getColourFromPalette(colour)
    fmt.Printf("(%d, %d): %x\n", x, (*this).line.get(), (*this).buffer[(*this).line.get()][x])
  }
}

func (this *Gpu) getColourFromPalette(colour uint8) uint8 {
  //pal := (*this).palette.get()
  //todo
  return colour
}

func (this *Gpu) renderSprites() {
  //todo
}

func (this *Gpu) hdma_transfer() {
  if((*(*this).mu).dma_status) {
    source := uint16((*(*this).mu).dma_val) << 8
    for i := uint16(0); i <= 0xa0; i++ {
      val := (*(*this).mu).Read_8(source + i)
      (*(*this).mu).Write_8(0xfe00 + i, val)
    }
    (*(*this).mu).dma_status = false
  }
}

func (this *Gpu) updateRegisters() {
  if(!(*this).lcdc.get_lcd_enabled()) {
    this.clearScreen()
    (*this).line.set(0)
    (*this).stat.set_mode(hblank_mode)
    return
  }

  interrupt := false

  switch((*this).stat.get_mode()) {
    case hblank_mode: //= 0
      if((*this).clock >= 204) {
        (*this).clock = 0
        line := (*this).line.inc()
        if(line >= height) {
          interrupt = (*this).stat.set_mode(vblank_mode)
          (*this).vblank = true
        } else {
          this.scanline(line)
          interrupt = (*this).stat.set_mode(oam_mode)
        }
      }
    case vblank_mode: //= 1
      if((*this).clock >= 456) {
        (*this).clock = 0
        line := (*this).line.inc()
        if(line > 153) {
          (*this).ir.Request(0)
          interrupt = (*this).stat.set_mode(oam_mode)
          (*this).line.set(0)
        }
      }
    case oam_mode: //= 2
      if((*this).clock >= 80) {
        (*this).clock = 0
        (*this).stat.set_mode(data_mode)
        this.hdma_transfer()
      }
    case data_mode: //= 3
      if((*this).clock >= 172) {
        (*this).clock = 0
        interrupt = (*this).stat.set_mode(hblank_mode)
      }
  }

  if(interrupt) { (*this).ir.Request(1) }

  if((*this).line.get() == (*this).line.get_c()) {
    (*this).stat.set((*this).stat.get() | (1 << 2))
    if(((*this).stat.get() & (1 << 6)) > 0) {
      (*this).ir.Request(1)
    }
  } else {
    (*this).stat.set((*this).stat.get() & uint8(^(uint8(1) << 2)))
  }
}

func (this *Gpu) Step(cycles int) {
  (*this).clock += cycles
  this.updateRegisters()
}
