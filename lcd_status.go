package gb

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
