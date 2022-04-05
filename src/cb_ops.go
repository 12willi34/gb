package gb

import (
  //"fmt"
)

func SWAP_B(this *cpu) int {
  (*this).bc.w_high(this.swap((*this).bc.r_high()))
  return 8
}

func SWAP_A(this *cpu) int {
  (*this).af.w_high(this.swap((*this).af.r_high()))
  return 8
}

func SWAP_C(this *cpu) int {
  (*this).bc.w_low(this.swap((*this).bc.r_low()))
  return 8
}

func SWAP_D(this *cpu) int {
  (*this).de.w_high(this.swap((*this).de.r_high()))
  return 8
}

func SWAP_E(this *cpu) int {
  (*this).de.w_low(this.swap((*this).de.r_low()))
  return 8
}

func SWAP_H(this *cpu) int {
  (*this).hl.w_high(this.swap((*this).hl.r_high()))
  return 8
}

func SWAP_L(this *cpu) int {
  (*this).hl.w_low(this.swap((*this).hl.r_low()))
  return 8
}

func SWAP_HL(this *cpu) int {
  addr := (*this).hl.value
  (*(*this).mu).Write_8(addr, this.swap((*(*this).mu).Read_8(addr)))
  return 16
}

func BIT_6_B(this *cpu) int {
  this.bit(6, (*this).bc.r_high())
  return 8
}

func BIT_7_H(this *cpu) int {
  this.bit(7, (*this).hl.r_high())
  return 8
}

func BIT_7_A(this *cpu) int {
  this.bit(7, this.af.r_high())
  return 8
}

func RL_A(this *cpu) int {
  a := this.rotate_left((*this).af.r_high())
  (*this).af.w_high(a)
  return 8
}

func RL_B(this *cpu) int {
  a := this.rotate_left((*this).bc.r_high())
  (*this).bc.w_high(a)
  return 8
}

func RL_C(this *cpu) int {
  a := this.rotate_left((*this).bc.r_low())
  (*this).bc.w_low(a)
  return 8
}

func RL_D(this *cpu) int {
  a := this.rotate_left((*this).de.r_high())
  (*this).de.w_high(a)
  return 8
}

func RL_E(this *cpu) int {
  a := this.rotate_left((*this).de.r_low())
  (*this).de.w_low(a)
  return 8
}

func RL_H(this *cpu) int {
  a := this.rotate_left((*this).hl.r_high())
  (*this).hl.w_high(a)
  return 8
}

func RL_L(this *cpu) int {
  a := this.rotate_left((*this).hl.r_low())
  (*this).hl.w_low(a)
  return 8
}

func SRL_A(this *cpu) int {
  (*this).af.w_high(this.srl((*this).af.r_high()))
  return 8
}

func SRL_B(this *cpu) int {
  (*this).bc.w_high(this.srl((*this).bc.r_high()))
  return 8
}

func SRL_C(this *cpu) int {
  (*this).bc.w_low(this.srl((*this).bc.r_low()))
  return 8
}

func SRL_D(this *cpu) int {
  (*this).de.w_high(this.srl((*this).de.r_high()))
  return 8
}

func SRL_E(this *cpu) int {
  (*this).de.w_low(this.srl((*this).de.r_low()))
  return 8
}

func SRL_H(this *cpu) int {
  (*this).hl.w_high(this.srl((*this).hl.r_high()))
  return 8
}

func SRL_L(this *cpu) int {
  (*this).hl.w_low(this.srl((*this).hl.r_low()))
  return 8
}

func SRL_HL(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  val = this.srl(val)
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 16
}

func RES_1_H(this *cpu) int {
  val := (*this).hl.r_high() & ^uint8(1 << 1)
  (*this).hl.w_high(val)
  return 8
}

func RES_0_HL(this *cpu) int {
  val := this.mu.Read_8(this.hl.value) & ^uint8(1 << 0)
  this.mu.Write_8(this.hl.value, val)
  return 16
}

func RES_0_A(this *cpu) int {
  val := (*this).af.r_high() & ^uint8(1 << 0)
  (*this).af.w_high(val)
  return 8
}

func SLA_A(this *cpu) int {
  this.af.w_high(this.shift_left_carry(this.af.r_high()))
  return 8
}

func (this *cpu) do_cb_op(op uint8) int {
  switch(op) {
  case 0x10:
    return RL_B(this)
  case 0x11:
    return RL_C(this)
  case 0x12:
    return RL_D(this)
  case 0x13:
    return RL_E(this)
  case 0x14:
    return RL_H(this)
  case 0x15:
    return RL_L(this)
  case 0x17:
    return RL_A(this)
  case 0x27:
    return SLA_A(this)
  case 0x30:
    return SWAP_B(this)
  case 0x31:
    return SWAP_C(this)
  case 0x37:
    return SWAP_A(this)
  case 0x32:
    return SWAP_D(this)
  case 0x33:
    return SWAP_E(this)
  case 0x34:
    return SWAP_H(this)
  case 0x35:
    return SWAP_L(this)
  case 0x36:
    return SWAP_HL(this)
  case 0x38:
    return SRL_B(this)
  case 0x39:
    return SRL_C(this)
  case 0x3a:
    return SRL_D(this)
  case 0x3b:
    return SRL_E(this)
  case 0x3c:
    return SRL_H(this)
  case 0x3d:
    return SRL_L(this)
  case 0x3e:
    return SRL_HL(this)
  case 0x3f:
    return SRL_A(this)
  case 0x7c:
    return BIT_7_H(this)
  case 0x7f:
    return BIT_7_A(this)
  case 0x86:
    return RES_0_HL(this)
  case 0x8c:
    return RES_1_H(this)
  case 0x87:
    return RES_0_A(this)
  default:
    return -1
  }
}
