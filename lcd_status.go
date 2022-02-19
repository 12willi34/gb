package gb

const lcdc = 0xff40
const lcdc_status = 0xff41

type Lcdc struct {
  mu memoryunit
}

func NewLcdc(mu memoryunit) Lcdc {
  return Lcdc {
    mu: mu,
  }
}

func (this Lcdc) lcd_enabled() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 7)) == 1
}

func (this Lcdc) tile_map_select() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 6)) == 1
}

func (this Lcdc) window_enabled() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 5)) == 1
}

func (this Lcdc) bg_tile_data_select() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 4)) == 1
}

func (this Lcdc) bg_tile_map_select() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 3)) == 1
}

func (this Lcdc) obj_size() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 2)) == 1
}

func (this Lcdc) obj_enabled() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 1)) == 1
}

func (this Lcdc) bg() bool {
  val := this.mu.Read_8(lcdc)
  return (val & (1 << 0)) == 1
}

type Lcdc_status struct {
  mu memoryunit
}

func NewLcdc_status(mu memoryunit) Lcdc_status {
  return Lcdc_status {
    mu: mu,
  }
}

func (this Lcdc_status) lyc_interrupt() bool {
  val := this.mu.Read_8(lcdc_status)
  return (val & (1 << 6)) == 1
}

func (this Lcdc_status) oam_interrupt() bool {
  val := this.mu.Read_8(lcdc_status)
  return (val & (1 << 5)) == 1
}

func (this Lcdc_status) vblank_interrupt() bool {
  val := this.mu.Read_8(lcdc_status)
  return (val & (1 << 4)) == 1
}

func (this Lcdc_status) hblank_interrupt() bool {
  val := this.mu.Read_8(lcdc_status)
  return (val & (1 << 3)) == 1
}

func (this Lcdc_status) coincidence_flag() bool {
  val := this.mu.Read_8(lcdc_status)
  return (val & (1 << 2)) == 1
}

func (this Lcdc_status) mode_flag() uint8 {
  val := this.mu.Read_8(lcdc_status)
  return val & uint8(0x03)
}
