package gb

//SDL
const btn_up = 1073741906
const btn_down = 1073741905
const btn_left = 1073741904
const btn_right = 1073741903
const btn_a = 121
const btn_b = 120
const btn_select = 8
var btn_start int = 13

type key struct {
  name string
  code int
  regPlace int
  is_direction bool
  state bool
}

func NewKey(n string, code, rp int) key {
  return key {
    name: n,
    code: code,
    regPlace: rp,
    state: false,
  }
}

type io_controller struct {
  dirKeys []key
  actionKeys []key
  Direction bool
  Action bool
}

func NewIoController() io_controller {
  return io_controller {
    dirKeys: []key {
      NewKey("right", 1073741903, 0),
      NewKey("left", 1073741904, 1),
      NewKey("up", 1073741906, 2),
      NewKey("down", 1073741905, 3),
    },
    actionKeys: []key {
      NewKey("A", 121, 0),
      NewKey("B", 120, 1),
      NewKey("select", 8, 2),
      NewKey("start", 13, 3),
    },
    Direction: false,
    Action: false,
  }
}

func (this *io_controller) ChangeMode(x uint8) {
  this.Direction = bool(0 == (x & (1 << 4)))
  this.Action = bool(0 == (x & (1 << 5)))
}

func (this *io_controller) Set(keyCode int, state bool) {
  for i := 0; i < len(this.dirKeys); i++ {
    if this.dirKeys[i].code == keyCode {
      this.dirKeys[i].state = state
      return
    }
  }
  for i := 0; i < len(this.actionKeys); i++ {
    if this.actionKeys[i].code == keyCode {
      this.actionKeys[i].state = state
      return
    }
  }
}

func (this *io_controller) Get() uint8 {
  x := uint8(0b00111111)
  if(this.Direction) {
    x &= ^uint8(1 << 4)
    for i := 0; i < len(this.dirKeys); i++ {
      if this.dirKeys[i].state {
        x &= ^uint8(1 << this.dirKeys[i].regPlace)
      } else {
        x |= uint8(1 << this.dirKeys[i].regPlace)
      }
    }
  }
  if(this.Action) {
    x &= ^uint8(1 << 5)
    for i := 0; i < len(this.actionKeys); i++ {
      if this.actionKeys[i].state {
        x &= ^uint8(1 << this.actionKeys[i].regPlace)
      } else {
        x |= uint8(1 << this.actionKeys[i].regPlace)
      }
    }
  }
  return x
}
