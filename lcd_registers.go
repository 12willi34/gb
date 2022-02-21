package gb

const lcdc = 0xff40
const stat = 0xff41
const scy = 0xff42
const scx = 0xff43
const ly = 0xff44
const lyc = 0xff45
const wy = 0xff4a
const wx = 0xff4b
const dma = 0xff46
const bgp = 0xff47
const obp0 = 0xff48
const obp1 = 0xff49

/* lcdc */

type Lcdc struct { mu memoryunit }

func NewLcdc(mu memoryunit) Lcdc {
  return Lcdc { mu: mu, }
}

func (this Lcdc) get() uint8 {
  return this.mu.Read_8(lcdc)
}

func (this Lcdc) set(val uint8 ) {
  this.mu.Write_8(lcdc, val)
}

func (this Lcdc) get_lcd_enabled() bool {
  return 0 > this.mu.Read_8(lcdc) & (1 << 7)
}

/* Stat */

type Stat struct { mu memoryunit }

func NewStat(mu memoryunit) Stat {
  return Stat { mu: mu, }
}

func (this Stat) get() uint8 {
  return this.mu.Read_8(stat)
}

func (this Stat) set(val uint8) {
  this.mu.Write_8(stat, val)
}

func (this Stat) get_mode() uint8 {
  return this.get() & 0x03
}

func (this Stat) set_mode(val uint8) bool {
  status := this.get()
  this.set(status | val)
  if(val != 3) {
    return 0 < (status & (1 << (3 + val)))
  }
  return false
}

func (this Stat) get_vblank() bool {
  return ((this.mu.Read_8(stat) & (1 << 4)) > 0)
}

/* Line */

type Line struct { mu memoryunit }

func NewLine(mu memoryunit) Line {
  return Line { mu: mu, }
}

func (this Line) inc() uint8 {
  this.set(this.get() + 1)
  return this.get()
}

func (this Line) get() uint8 {
  return this.mu.Read_8(ly)
}

func (this Line) set(val uint8) {
  this.mu.Write_8(ly, val)
}

func (this Line) get_c() uint8 {
  return this.mu.Read_8(lyc)
}

func (this Line) set_c(val uint8) {
  this.mu.Write_8(lyc, val)
}

/* Scroll */

type Scroll struct { mu memoryunit }

func NewScroll(mu memoryunit) Scroll {
  return Scroll { mu: mu, }
}

func (this Scroll) get_x() uint8 {
  return this.mu.Read_8(scx)
}

func (this Scroll) set_x(val uint8) {
  this.mu.Write_8(scx, val)
}

func (this Scroll) get_y() uint8 {
  return this.mu.Read_8(scy)
}

func (this Scroll) set_y(val uint8) {
  this.mu.Write_8(scy, val)
}

/* Window */

type Window struct { mu memoryunit }

func NewWindow(mu memoryunit) Window {
  return Window { mu: mu, }
}

func (this Window) get_x() uint8 {
  return this.mu.Read_8(wx)
}

func (this Window) set_x(val uint8) {
  this.mu.Write_8(wx, val)
}

func (this Window) get_y() uint8 {
  return this.mu.Read_8(wy)
}

func (this Window) set_y(val uint8) {
  this.mu.Write_8(wy, val)
}

/* Palette */

type Palette struct { mu memoryunit }

func NewPalette(mu memoryunit) Palette {
  return Palette { mu: mu, }
}

func (this Palette) get() uint8 {
  return this.mu.Read_8(bgp)
}

func (this Palette) set(val uint8) {
  this.mu.Write_8(bgp, val)
}

/* ObjPalette0 */

type ObjPalette0 struct { mu memoryunit }

func NewObjPalette0(mu memoryunit) ObjPalette0 {
  return ObjPalette0 { mu: mu, }
}

func (this ObjPalette0) get() uint8 {
  return this.mu.Read_8(obp0)
}

func (this ObjPalette0) set(val uint8) {
  this.mu.Write_8(obp0, val)
}

/* ObjPalette1 */

type ObjPalette1 struct { mu memoryunit }

func NewObjPalette1(mu memoryunit) ObjPalette1 {
  return ObjPalette1 { mu: mu, }
}

func (this ObjPalette1) get() uint8 {
  return this.mu.Read_8(obp1)
}

func (this ObjPalette1) set(val uint8) {
  this.mu.Write_8(obp1, val)
}
