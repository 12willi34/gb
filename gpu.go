package gb

import ()

const width = 160
const height = 144

type Gpu struct {
  mu *memoryunit
  lcdc Lcdc
  lcdc_status Lcdc_status
}

func NewGpu(mu *memoryunit) *Gpu {
  return &(Gpu {
    mu: mu,
    lcdc: NewLcdc(*mu),
    lcdc_status: NewLcdc_status(*mu),
  })
}

func (this *Gpu) Step(cycles int) {
  if(!(*this).lcdc.lcd_enabled()) {
    return
  }
}
