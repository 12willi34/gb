package gb

import ()

const width = 160
const height = 144
const hblank_mode = uint8(0)
const vblank_mode = uint8(1)
const oam_mode = uint8(2)
const data_mode = uint8(3)

type Gpu struct {
  mu *memoryunit
  ir interrupter
  lcdc Lcdc
  stat Stat
  scroll Scroll
  window Window
  palette Palette
  palette0 ObjPalette0
  palette1 ObjPalette1
  line Line
  clock int
}

func NewGpu(mu *memoryunit, interrupter interrupter) *Gpu {
  x := &(Gpu {
    mu: mu,
    ir: interrupter,
    lcdc: NewLcdc(*mu),
    stat: NewStat(*mu),
    scroll: NewScroll(*mu),
    window: NewWindow(*mu),
    palette: NewPalette(*mu),
    palette0: NewObjPalette0(*mu),
    palette1: NewObjPalette1(*mu),
    line: NewLine(*mu),
    clock: 0,
  })
  (*x).lcdc.set((*x).lcdc.get() | (1 << 7)) //enable lcd
  return x
}

func (this *Gpu) clearScreen() {
  //todo
}

func (this *Gpu) scanline(line uint8) {
  //todo
}

func (this *Gpu) hdma_transfer() {
  //todo
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

  if(interrupt) {
    (*this).ir.Request(1)
  }

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
