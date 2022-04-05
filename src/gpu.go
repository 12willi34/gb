package gb

//import "fmt"
//import "strconv"

const (
  width = 160
  height = 144
  hblank_mode = uint8(0)//uint8(0)
  vblank_mode = uint8(1)//uint8(1)
  oam_mode = uint8(2)//uint8(2)
  data_mode = uint8(3)//uint8(3)
  col_white = 0xff
  col_light_gray = 0xff*(3/4)
  col_dark_gray = 0xff*(1/4)
  col_black = 0x00
)

type Gpu struct {
  mu *memoryunit
  ir *interrupter
  clock int

  PubClock int
  PubMode int
  mode uint8

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

func NewGpu(mu *memoryunit, interrupter *interrupter) Gpu {
  x := Gpu {
    mu: mu,
    ir: interrupter,
    clock: 0,

    PubClock: 0,
    PubMode: int(oam_mode),
    mode: oam_mode,

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
  x.stat.set_mode(oam_mode)
  return x
}

func (this *Gpu) clearScreen() {
  for i := 0; i < len((*this).buffer); i++ {
    for j := 0; j < len((*this).buffer[0]); j++ {
      (*this).buffer[i][j] = col_white
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

    var colour_bit int = (((int(xPos) % 8) - 7)*-1)

    colour := uint8(0)
    if((val0 & (1 << uint8(colour_bit))) > 0) {
      colour |= 1 << 1
    }
    if((val1 & (1 << uint8(colour_bit))) > 0) {
      colour |= 1
    }
    (*this).buffer[(*this).line.get()][x] = this.getColour(colour, 0xff47)
  }
}

func (this *Gpu) getColour(colour uint8, addr uint16) uint8 {
  var h = 1 + 2*colour
  var l = 2*colour
  palette := (*(*this).mu).Read_8(addr)
  res := uint8(0)
  if((palette & (1 << h)) > 0) { res += 1 << 1 }
  if((palette & (1 << l)) > 0) { res += 1 << 0 }
  switch(res) {
    case 0:
      res = col_white
    case 1:
      res = col_light_gray
    case 2:
      res = col_dark_gray
    case 3:
      res = col_black
  }
  return res
}

func (this *Gpu) renderSprites() {
  lcdc := (*this).lcdc.get()
  use8x16 := bool((lcdc & (1 << 2)) > 0)
  for ind := uint16(0); ind < 40; ind++ {
    sprite_ind := uint8(ind*4)
    yPos := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind)) - 16)
    xPos := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind) + 1) - 8)
    tileLocation := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind) + 2))
    attributes := uint8((*(*this).mu).Read_8(0xfe00 + uint16(sprite_ind) + 3))
    yFlip := bool((attributes & (1 << 6)) > 0)
    xFlip := bool((attributes & (1 << 5)) > 0)
    line := int((*this).line.get())
    var ySize int = 8
    if(use8x16) { ySize = 16 }
    if((line >= int(yPos)) && (line <= (int(yPos) + ySize))) {
      l := line - int(yPos)
      if(yFlip) {
        l -= ySize
        l *= -1
      }
      l *= 2
      dataAddr := uint16(0x8000) + uint16(tileLocation*16) + uint16(l)
      data1 := uint8((*(*this).mu).Read_8(dataAddr))
      data2 := uint8((*(*this).mu).Read_8(dataAddr + 1))
      for pixel := 7; pixel >= 0; pixel-- {
        colour := pixel
        if(xFlip) { colour = (colour - 7)*(-1) }
        colourNum := 0
        if((data2 & (1 << colour)) > 0) { colourNum |= 1 << 1 }
        if((data1 & (1 << colour)) > 0) { colourNum |= 1 << 0 }
        colourAddr := 0xff48
        if((attributes & (1 << 4)) > 0) { colourAddr = 0xff49 }
        res := (*this).getColour(uint8(colourNum), uint16(colourAddr))
        if(res == col_white) { continue }
        x := int(7 - pixel)
        (*this).buffer[int(int(xPos) + x)][line] = res
      }
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
          (*this).vblank = true
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
