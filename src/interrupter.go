package gb

const enable_register = 0xffff
const flag_register = 0xff0f
const vblank_flag = 0
const lcd_start_flag = 1
const timer_flag = 2
const serial_flag = 3
const joypad_flag = 4
var flag_addr_table = []uint16{0x40, 0x48, 0x50, 0x58, 0x60}

type interrupter struct {
  mu *memoryunit
  processor *cpu
}

func NewInterrupter(mu *memoryunit, processor *cpu) interrupter {
  return interrupter {
    mu: mu,
    processor: processor,
  }
}

func (this *interrupter) Request(flag int) {
  x := this.mu.Read_8(flag_register) | uint8(1 << flag)
  this.mu.Write_8(flag_register, x)
}

func (this *interrupter) handle() {
  if(this.processor.EnInterrupt) {
    this.processor.EnInterrupt = false
    this.processor.Interrupt = true
    return
  }
  if(this.processor.Interrupt) {
    flag := this.mu.Read_8(flag_register)
    enable := this.mu.Read_8(enable_register)
    x := flag & enable
    if(x == 0) { return }
    this.processor.Halt = false
    this.processor.Interrupt = false
    for i := 0; i < 5; i++ {
      if 0 < (x & (1 << i)) {
        this.mu.Write_8(flag_register, flag & ^uint8(1 << i))
        this.processor.pushStack(this.processor.pc.value)
        this.processor.pc.value = flag_addr_table[i]
        return
      }
    }
  }
}
