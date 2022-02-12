package gb

import ()

func NOOP(this *cpu) int {
  return 4
}

func DEC_B(this *cpu) int {
  (*this).bc.w_high(this.decrement((*this).bc.r_high()))
  return 4
}

func LD_B_n(this *cpu) int {
  (*this).bc.w_high(this.fetch())
  return 8
}

func JR_nZ(this *cpu) int {
  a := int8(this.fetch())
  ticks := 8
  if(!this.get_f_zero()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
    ticks += 4
  }
  return 8
}

func LD_B_B(this *cpu) int {
  (*this).bc.w_high((*this).bc.r_high())
  return 4
}

func POP_BC(this *cpu) int {
  (*this).bc.value = this.popStack()
  return 12
}

func JP(this *cpu) int {
  a := this.fetch()
  b := this.fetch()
  this.pc.value = uint16(a | b<<8)
  return 12
}

func PUSH_BC(this *cpu) int {
  this.pushStack((*this).bc.value)
  return 16
}

func RET_Z(this *cpu) int {
  if(this.get_f_zero()) {
    this.pc.value = this.popStack()
  }
  return 8
}

func CALL(this *cpu) int {
  (*this).pushStack((*this).pc.value)
  (*this).pc.value = this.fetch_16()
  return 12
}

func LD_n_A(this *cpu) int {
  i := 0xff00 + uint16(this.fetch())
  (*this.memory).Write_8(i, this.af.r_high())
  return 12
}

func LD_A_n(this *cpu) int {
  a := (*this.memory).Read_8(0xff00 + uint16(this.fetch()))
  this.af.w_high(a)
  return 12
}

func EI(this *cpu) int {
  (*this).Interrupt = true
  return 4
}

func CP_n(this *cpu) int {
  this.compare_8(this.af.r_high(), this.fetch())
  return 8
}

func (this *cpu) init_ops() [0x100]func(*cpu) int {
  var ops [0x100]func(*cpu) int
  ops[0x00] = NOOP
  ops[0x05] = DEC_B
  ops[0x06] = LD_B_n
  ops[0x20] = JR_nZ
  ops[0xc1] = POP_BC
  ops[0xc3] = JP
  ops[0xc5] = PUSH_BC
  ops[0xc8] = RET_Z
  ops[0xcd] = CALL
  ops[0xe0] = LD_n_A
  ops[0xf0] = LD_A_n
  ops[0xfb] = EI
  ops[0xfe] = CP_n
  return ops
}
