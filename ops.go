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

func LD_A_A(this *cpu) int {
  (*this).af.w_high((*this).af.r_high())
  return 4
}

func LD_A_B(this *cpu) int {
  (*this).af.w_high((*this).bc.r_high())
  return 4
}

func LD_A_C(this *cpu) int {
  (*this).af.w_high((*this).bc.r_low())
  return 4
}

func LD_A_D(this *cpu) int {
  (*this).af.w_high((*this).de.r_high())
  return 4
}

func LD_A_E(this *cpu) int {
  (*this).af.w_high((*this).de.r_low())
  return 4
}

func LD_A_H(this *cpu) int {
  (*this).af.w_high((*this).hl.r_high())
  return 4
}

func LD_A_L(this *cpu) int {
  (*this).af.w_high((*this).hl.r_low())
  return 4
}

func LD_A_HL(this *cpu) int {
  (*this).af.w_high((*this).memory.Read_8((*this).hl.value))
  return 8
}

func LD_B_B(this *cpu) int {
  (*this).bc.w_high((*this).bc.r_high())
  return 4
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

func LD_A_BC(this *cpu) int {
  val := (*this).bc.value
  (*this).af.w_high((*this).memory.Read_8(val))
  return 8
}

func LD_A_DE(this *cpu) int {
  val := (*this).de.value
  (*this).af.w_high((*this).memory.Read_8(val))
  return 8
}

func LD_A_nn(this *cpu) int {
  val := this.fetch_16()
  (*this).af.w_high((*this).memory.Read_8(val))
  return 16
}

func LD_A_number(this *cpu) int {
  (*this).af.w_high(this.fetch())
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

func RET(this *cpu) int {
  (*this).pc.value = this.popStack()
  return 8
}

func CALL(this *cpu) int {
  (*this).pushStack((*this).pc.value)
  (*this).pc.value = this.fetch_16()
  return 12
}

func EI(this *cpu) int {
  (*this).Interrupt = true
  return 4
}

func DI(this *cpu) int {
  (*this).Interrupt = false
  return 4
}

func CP_n(this *cpu) int {
  this.compare_8(this.af.r_high(), this.fetch())
  return 8
}

func SBC_A_A(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).af.r_high()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_B(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).bc.r_high()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_C(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).bc.r_low()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_D(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).de.r_high()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_E(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).de.r_low()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_H(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).hl.r_high()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_L(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).hl.r_low()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 4
}

func SBC_A_HL(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).memory.Read_8((*this).hl.value)
  (*this).af.w_high(this.subtract_carry(a, b))
  return 8
}

func (this *cpu) init_ops() [0x100]func(*cpu) int {
  var ops [0x100]func(*cpu) int
  ops[0x00] = NOOP
  ops[0x05] = DEC_B
  ops[0x06] = LD_B_n
  ops[0x0a] = LD_A_BC
  ops[0x1a] = LD_A_DE
  ops[0x20] = JR_nZ
  ops[0x3e] = LD_A_number
  ops[0x40] = LD_B_B
  ops[0x78] = LD_A_B
  ops[0x79] = LD_A_C
  ops[0x7a] = LD_A_D
  ops[0x7b] = LD_A_E
  ops[0x7c] = LD_A_H
  ops[0x7d] = LD_A_L
  ops[0x7e] = LD_A_HL
  ops[0x7f] = LD_A_A
  ops[0x98] = SBC_A_B
  ops[0x99] = SBC_A_C
  ops[0x9a] = SBC_A_D
  ops[0x9b] = SBC_A_E
  ops[0x9c] = SBC_A_H
  ops[0x9d] = SBC_A_L
  ops[0x9e] = SBC_A_HL
  ops[0x9f] = SBC_A_A
  ops[0xc1] = POP_BC
  ops[0xc3] = JP
  ops[0xc5] = PUSH_BC
  ops[0xc8] = RET_Z
  ops[0xc9] = RET
  ops[0xcd] = CALL
  ops[0xe0] = LD_n_A
  ops[0xf0] = LD_A_n
  ops[0xf3] = DI
  ops[0xfa] = LD_A_nn
  ops[0xfb] = EI
  ops[0xfe] = CP_n
  return ops
}
