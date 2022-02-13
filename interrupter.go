package gb

import ()

const enable_register = 0xffff
const flag_register = 0xff0f

type interrupter struct {
  bus memoryunit
  processor cpu
}

func Interrupter(mu memoryunit, processor cpu) *interrupter {
  return &interrupter {
    bus: mu,
    processor: processor,
  }
}

func (this *interrupter) check_interrupt(inter_f uint8, inter_e uint8, i int) bool {
  return (1 & (inter_f << i) == 1) && (1 & (inter_e << i) == 1)
}

func (this *interrupter) do_interrupt(i int) {
  //ToDo
}

func (this *interrupter) handle() {
  if((*this).processor.Interrupt) {
    inter_f := (*this).bus.Read_8(flag_register)
    if(inter_f > 0) {
      inter_e := (*this).bus.Read_8(enable_register)
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
