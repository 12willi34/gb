package gb

const (
  width = 160
  height = 144
  hblank_mode = uint8(0)
  vblank_mode = uint8(1)
  oam_mode = uint8(2)
  data_mode = uint8(3)
  col_white = 0xff
  col_light_gray = 191
  col_dark_gray = 63
  col_black = 0x00
)

type Gpu struct {
  mu *memoryunit
  ir *interrupter
  clock int

  PubClock int
  PubMode int
  mode uint8

  Buffer [height][width]uint8
  Vblank bool

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

func NewGpu(mu *memoryunit, interrupter *interrupter) Gpu {
  x := Gpu {
    mu: mu,
    ir: interrupter,
    clock: 0,

    PubClock: 0,
    PubMode: int(oam_mode),
    mode: oam_mode,

    Buffer: [height][width]uint8{},
    Vblank: false,

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
  x.stat.set_mode(oam_mode)
  return x
}

func (this *Gpu) clearScreen() {
  for i := 0; i < len((*this).Buffer); i++ {
    for j := 0; j < len((*this).Buffer[0]); j++ {
      (*this).Buffer[i][j] = col_white
    }
  }
}

func (this *Gpu) scanline(line uint8) {
  if(!(*this).lcdc.get_lcd_enabled()) { return }
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
  in_window := (wy <= (*this).line.get()) && (((*this).lcdc.get() & (1 << 5)) > 0)

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
  ly := (*this).line.get()
  if(!in_window) {
    yPos = ly + sy
  } else {
    yPos = ly - wy
  }
  line := (yPos % 8)*2
  row := uint16(yPos/8)*32
  pal := this.mu.Read_8(0xff47)
  for x := uint8(0); x < 160; x++ {
    xPos = x + sx
    if(x >= wx && in_window) {
      xPos = x - wx
    }
    col := uint16(xPos/8)
    var tile_num int16
    tile_adr := background_start + row + col
    tile_loc := tile_start
    if(unsigned_tile_address) {
      tile_num = int16(this.mu.Read_8(tile_adr))
      tile_loc += uint16(tile_num)*16
    } else {
      tile_num = int16(int8((*(*this).mu).Read_8(tile_adr)))
      tile_loc = uint16(int32(tile_loc) + (int32(tile_num) + 128)*16)
    }
    temp := tile_loc + uint16(line)
    val0 := this.mu.Read_8(temp)
    val1 := this.mu.Read_8(temp + 1)

    colour_bit := (((int(xPos) % 8) - 7)*-1)

    colour := uint8(0)
    if((val0 & (1 << colour_bit)) > 0) {
      colour |= 1
    }
    if((val1 & (1 << colour_bit)) > 0) {
      colour |= 2
    }
    this.Buffer[ly][x] = this.getColour(colour, pal)
  }
}

func (this *Gpu) getColour(colour uint8, pal uint8) uint8 {
  l := colour << 1
  h := l | 1
  res := uint8(0)
  if((pal & (1 << h)) > 0) { res += 2 }
  if((pal & (1 << l)) > 0) { res += 1 }
  switch(res) {
    case 0:
      return col_white
    case 1:
      return col_light_gray
    case 2:
      return col_dark_gray
    case 3:
      return col_black
  }
  return res
}

func (this *Gpu) renderSprites() {
  use8x16 := bool((this.lcdc.get() & (1 << 2)) > 0)
  line := this.line.get()
  var ySize int = 8
  if(use8x16) { ySize = 16 }
  pal1, pal2 := this.mu.Read_8(0xff48), this.mu.Read_8(0xff49)
  var pal uint8
  for ind := uint16(0); ind < 40; ind++ {
    sprite_ind := uint8(ind*4)
    yPos := this.mu.Read_8(0xfe00 + uint16(sprite_ind)) - 16
    if line < yPos || line >= (yPos + uint8(ySize)) {
      continue
    }
    xPos := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind) + 1) - 8)
    tileLocation := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind) + 2))
    attributes := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind) + 3))
    yFlip := (attributes & (1 << 6)) > 0
    xFlip := (attributes & (1 << 5)) > 0
    behindBg := 0 < attributes & (1 << 7)
    l := int(line - yPos)
    if(yFlip) {
      l = -1*(l - int(ySize))
    }
    l *= 2
    dataAddr := 0x8000 + uint16(tileLocation)*16 + uint16(l)
    data1 := this.mu.Read_8(dataAddr)
    data2 := this.mu.Read_8(dataAddr + 1)
    for pixel := 7; pixel >= 0; pixel-- {
      x := xPos + uint8(7 - pixel)
      if x < 0 || x >= width { continue }
      colour := pixel
      if(xFlip) { colour = -1*(colour - 7) }
      colourNum := 0
      if((data2 & (1 << colour)) > 0) { colourNum += 2 }
      if((data1 & (1 << colour)) > 0) { colourNum += 1 }
      if colourNum == 0 { continue }
      pal = pal1
      if((attributes & (1 << 4)) > 0) { pal = pal2 }
      if(behindBg && this.Buffer[line][x] != col_white) { continue }
      this.Buffer[line][x] = this.getColour(uint8(colourNum), pal)
    }
  }
}

func (this *Gpu) setMode(mode uint8) {
  this.mode = mode
  this.stat.set_mode(mode)
}

func (this *Gpu) Step(cycles int) {
  this.clock += cycles/4
  switch(this.mode) {
    case hblank_mode: //= 0
      if((*this).clock >= 204) {
        (*this).clock %= 204
        this.scanline(this.line.get())
        line := this.line.inc()
        if(line == 144) {
          (*this).setMode(vblank_mode)
          this.ir.Request(0)
        } else {
          this.setMode(oam_mode)
        }
      }
    case vblank_mode: //= 1
      if((*this).clock >= 456) {
        (*this).clock %= 456
        line := (*this).line.inc()
        if(line == 154) {
          this.setMode(oam_mode)
          (*this).line.set(0)
          (*this).Vblank = true
        }
      }
    case oam_mode: //= 2
      if((*this).clock >= 80) {
        (*this).clock %= 80
        this.setMode(data_mode)
      }
    case data_mode: //= 3
      if((*this).clock >= 172) {
        (*this).clock %= 172
        this.setMode(hblank_mode)
        if(0 < (this.stat.get() & (1 << 3))) {
          this.ir.Request(1)
        }
        coincidence := bool((*this).line.get() == (*this).line.get_c())
        if(coincidence && ((this.stat.get() & (1 << 6)) > 0)) {
          this.ir.Request(1)
        }
        if(coincidence) {
          this.stat.set(this.stat.get() | uint8(1 << 2))
        } else {
          this.stat.set(this.stat.get() & ^uint8(1 << 2))
        }
      }
  }

  //temp
  this.PubClock = this.clock
  this.PubMode = int(this.mode)
}
