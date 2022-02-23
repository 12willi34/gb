package gb

import (
  //"fmt"
)

const enable_register = 0xffff
const flag_register = 0xff0f
const vblank_flag = 0
const lcd_start_flag = 1
const timer_flag = 2
const serial_flag = 3
const joypad_flag = 4
var flag_addr_table = []uint16{0x40, 0x48, 0x50, 0x58, 0x60}

type interrupter struct {
  mu memoryunit
  processor cpu
}

func NewInterrupter(mu memoryunit, processor cpu) interrupter {
  return interrupter {
    mu: mu,
    processor: processor,
  }
}

func (this *interrupter) set_flag(i int) {
  new_flag_val := (*this).mu.Read_8(flag_register) | (1 << i)
  (*this).mu.Write_8(flag_register, new_flag_val)
}

func (this *interrupter) Request(flag int) {
  this.set_flag(flag)
}

func (this *interrupter) check_interrupt(inter_f uint8, inter_e uint8, i int) bool {
  return (1 & (inter_f << i) == 1) && (1 & (inter_e << i) == 1)
}

func (this *interrupter) do_interrupt(i int) {
  (*this).processor.Interrupt = false
  (*this).processor.pushStack((*this).processor.pc.value)
  (*this).processor.pc.value = flag_addr_table[i]
}

func (this *interrupter) handle() {
  if((*this).processor.Interrupt) {
    inter_f := (*this).mu.Read_8(flag_register)
    if(inter_f > 0) {
      inter_e := (*this).mu.Read_8(enable_register)
      for i := 0; i < 5; i++ {
        if(this.check_interrupt(inter_f, inter_e, i)) {
          this.do_interrupt(i)
          this.processor.Interrupt = false
          return
        }
      }
    }
  }
}
