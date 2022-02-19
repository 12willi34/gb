package gb

import ("fmt")

const width = 160
const height = 144
const hblank_mode = uint8(0)
const vblank_mode = uint8(1)
const oam_mode = uint8(2)
const data_mode = uint8(3)

type Gpu struct {
  mu *memoryunit
  lcdc Lcdc
  lcdc_status Lcdc_status
  screen []uint8
  mode uint8
  clock int
  line int
}

func NewGpu(mu *memoryunit) *Gpu {
  return &(Gpu {
    mu: mu,
    lcdc: NewLcdc(*mu),
    lcdc_status: NewLcdc_status(*mu),
    screen: make([]uint8, width*height),
    mode: 0,
    clock: 0,
    line: 0,
  })
}

func (this *Gpu) updateScreen() {
  fmt.Println("new screen")
  //todo
}

func (this *Gpu) renderLine() {
  //todo
}

func (this *Gpu) Step(cycles int) {
  (*this).clock += cycles

  /*
  if(!t.lcdc.lcd_enabled()) {
    return
  }
  */

  switch((*this).mode) {
    case hblank_mode:
      if((*this).clock >= 204) {
        (*this).clock = 0
        (*this).line++
        if((*this).line == 143) {
          (*this).mode = vblank_mode
          this.updateScreen()
        } else {
          (*this).mode = oam_mode
        }
      }
      break
    case vblank_mode:
      if((*this).clock >= 456) {
        (*this).clock = 0
        (*this).line++
        if((*this).line > 153) {
          (*this).mode = oam_mode
          (*this).line = 0
        }
      }
      break
    case oam_mode:
      if((*this).clock >= 80) {
        (*this).clock = 0
        (*this).mode = data_mode
      }
      break
    case data_mode:
      if((*this).clock >= 172) {
        (*this).clock = 0
        (*this).mode = hblank_mode
        this.renderLine()
      }
      break
  }
}
