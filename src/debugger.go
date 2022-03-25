package gb

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)

var reader = bufio.NewReader(os.Stdin)

type Debugger struct {
  to_ignore []uint8
  history []uint8
  ignore_history bool
}

func NewDebugger() Debugger {
  return Debugger {
    to_ignore: make([]uint8, 0),
    history: make([]uint8, 0),
    ignore_history: false,
  }
}

func in(a uint8, b []uint8) bool {
  for i := 0; i < len(b); i++ {
    if(b[i] == a) {
      return true
    }
  }
  return false
}

func (this *cpu) Step_debug() int {
  if((*this).Halt) { return 4 }
  var cycles = -1
  var cb_op = uint8(0)
  op := this.fetch()
  if(op == 0xcb) {
    cb_op = this.fetch()
    cycles = this.do_cb_op(cb_op)
  } else {
    cycles = this.do_op(op)
  }
  if(cycles == -1) {
    fmt.Printf("opcode not implemented")
    return -1
  }
  fmt.Printf("\n\nop: %02x\n\n", op)
  if(op == 0xcb) { fmt.Printf("cb_op: %02x\n", cb_op) }
  fmt.Printf("af: %04x\n", (*this).af.value)
  fmt.Printf("bc: %04x\n", (*this).bc.value)
  fmt.Printf("de: %04x\n", (*this).de.value)
  fmt.Printf("hl: %04x\n", (*this).hl.value)
  fmt.Printf("sp: %04x\n", (*this).sp.value)
  fmt.Printf("pc: %04x\n", (*this).pc.value)

  ignore := false

  if(len(this.debug.to_ignore) > 0) {
    if(in(op, this.debug.to_ignore)) {
      ignore = true
    } else {
      this.debug.to_ignore = make([]uint8, 0)
    }
  } else if(this.debug.ignore_history) {
    if(in(op, this.debug.history)) {
      ignore = true
    } else {
      this.debug.history = append(this.debug.history, op)
    }
  }

  if(ignore) { return cycles }

  fmt.Printf("\nOptions\n")
  fmt.Printf("\ti = ignore mode\n")
  fmt.Printf("\th = ignore history mode\n")
  fmt.Printf("new mode: ")
  mode, _ := reader.ReadString('\n')
  switch {
  case string(mode[0]) == "i":
    fmt.Printf("to ignore (comma separated): ")
    to_ignore_raw, _ := reader.ReadString('\n')
    to_ignore_splice := strings.Split(to_ignore_raw[:len(to_ignore_raw) - 1], ",")
    this.debug.to_ignore = make([]uint8, len(to_ignore_splice))
    for i := 0; i < len(to_ignore_splice); i++ {
      x, _ := strconv.ParseUint(to_ignore_splice[i], 0, 8)
      this.debug.to_ignore[i] = uint8(x)
    }
  case string(mode[0]) == "h":
    this.debug.ignore_history = true
    this.debug.history = append(this.debug.history, op)
  }
  return cycles
}
