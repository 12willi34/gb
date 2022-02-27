package gb

import (
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
  buffer [width][height]uint8

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
    buffer: [width][height]uint8{},

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

  tile_start := uint16(0x8800)
  unsigned_tile_address := false
  if(((*this).lcdc.get() & (1 << 4)) > 0) {
    tile_start = 0x8000
    unsigned_tile_address = true
  }

  var background_start uint16
  var ctrl_bit uint8
  if(in_window) {
    ctrl_bit = 3
  } else {
    ctrl_bit = 6
  }
  if(((*this).lcdc.get() & (1 << ctrl_bit)) > 0) {
    background_start = 0x9c00
  } else {
    background_start = 0x9800
  }

  yPos := (*this).line.get()
  if(in_window) {
    yPos -= wy
  } else {
    yPos += sy
  }
  row := uint16(uint8(yPos/8)*32)
  for x := uint8(0); x < 160; x++ {
    xPos := sx + x
    if(in_window && x >= wx) {
      xPos = x - wx
    }
    col := uint16(xPos/8)
    var tile_addr uint16
    if(unsigned_tile_address) {
      tile_addr = tile_start + (uint16(background_start + row + col)*16)
    } else {
      tile_addr = tile_start + ((uint16(int16(background_start) + int16(row) + int16(col)) + uint16(128))*16)
    }
    line := uint16((yPos % 8)*2) //nicht sicher ob das richtig ist
    val0 := (*(*this).mu).Read_8(tile_addr + line)
    val1 := (*(*this).mu).Read_8(tile_addr + line + 1)
    colour_bit := (int(xPos % 8) - 7)*(- 1) //nicht sicher ob das richtig ist v2
    colour := uint8(0)
    if((val0 & (1 << colour_bit)) > 0) {
      colour |= 1 << 1
    }
    if((val1 & (1 << colour_bit)) > 0) {
      colour |= 1
    }
    (*this).buffer[x][(*this).line.get()] = colour
  }
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
          /*
          for y := 0; y < 144; y++ {
            for x := 0; x < 160; x++ {
              print((*this).buffer[x][y])
            }
            println()
          }
          println()
          println()
          */
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
