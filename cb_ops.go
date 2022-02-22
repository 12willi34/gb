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
  (*(*this).memory).Write_8(addr, this.swap((*(*this).memory).Read_8(addr)))
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

func (this *cpu) init_cb_ops() [0x100]func(*cpu) int {
  var cb_ops [0x100]func(*cpu) int
  cb_ops[0x10] = RL_B
  cb_ops[0x11] = RL_C
  cb_ops[0x12] = RL_D
  cb_ops[0x13] = RL_E
  cb_ops[0x14] = RL_H
  cb_ops[0x15] = RL_L
  cb_ops[0x17] = RL_A
  cb_ops[0x30] = SWAP_B
  cb_ops[0x31] = SWAP_C
  cb_ops[0x37] = SWAP_A
  cb_ops[0x32] = SWAP_D
  cb_ops[0x33] = SWAP_E
  cb_ops[0x34] = SWAP_H
  cb_ops[0x35] = SWAP_L
  cb_ops[0x36] = SWAP_HL
  cb_ops[0x7c] = BIT_7_H
  return cb_ops
}
