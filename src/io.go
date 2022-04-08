package gb

//SDL
const btn_up = 1073741906
const btn_down = 1073741905
const btn_left = 1073741904
const btn_right = 1073741903
const btn_a = 121
const btn_b = 120
const btn_select = 8
const btn_start = 13

const ioRegister = 0xff00

type io_controller struct {
  up_state bool
  down_state bool
  left_state bool
  right_state bool
  a_state bool
  b_state bool
  select_state bool
  start_state bool

  mode uint8
}

func NewIoController() io_controller {
  return io_controller {
    up_state: false,
    down_state: false,
    left_state: false,
    right_state: false,
    a_state: false,
    b_state: false,
    select_state: false,
    start_state: false,

    mode: 0,
  }
}

func (this io_controller) ChangeMode(x uint8) {
  this.mode = x
}

func (this io_controller) SetBtn(btn int, state bool) {
  switch(btn) {
  case btn_up:
    this.up_state = state
  case btn_down:
    this.down_state = state
  case btn_left:
    this.left_state = state
  case btn_right:
    this.right_state = state
  case btn_a:
    this.a_state = state
  case btn_b:
    this.b_state = state
  case btn_select:
    this.select_state = state
  case btn_start:
    this.start_state = state
  }
}

func (this io_controller) Get() uint8 {
  x := this.mode
  if(0 == (x & (1 << 5))) {
    //action
    if(!this.start_state) { x |= (1 << 3) } else { x &= ^uint8(1 << 3) }
    if(!this.select_state) { x |= (1 << 2) } else { x &= ^uint8(1 << 2) }
    if(!this.b_state) { x |= (1 << 1) } else { x &= ^uint8(1 << 1) }
    if(!this.a_state) { x |= (1 << 0) } else { x &= ^uint8(1 << 0) }
  } else {
    //direction
    if(!this.down_state) { x |= (1 << 3) } else { x &= ^uint8(1 << 3) }
    if(!this.up_state) { x |= (1 << 2) } else { x &= ^uint8(1 << 2) }
    if(!this.left_state) { x |= (1 << 1) } else { x &= ^uint8(1 << 1) }
    if(!this.right_state) { x |= (1 << 0) } else { x &= ^uint8(1 << 0) }
  }
  return x
}
