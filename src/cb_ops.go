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

func BIT_6_C(this *cpu) int {
  this.bit(6, this.bc.r_low())
  return 8
}

func BIT_6_D(this *cpu) int {
  this.bit(6, this.de.r_high())
  return 8
}

func BIT_6_E(this *cpu) int {
  this.bit(6, this.de.r_low())
  return 8
}

func BIT_6_H(this *cpu) int {
  this.bit(6, this.hl.r_high())
  return 8
}

func BIT_6_L(this *cpu) int {
  this.bit(6, this.hl.r_low())
  return 8
}

func BIT_6_HL(this *cpu) int {
  this.bit(6, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_6_A(this *cpu) int {
  this.bit(6, this.af.r_high())
  return 8
}

func BIT_0_B(this *cpu) int {
  this.bit(0, this.bc.r_high())
  return 8
}

func BIT_0_C(this *cpu) int {
  this.bit(0, this.bc.r_low())
  return 8
}

func BIT_0_D(this *cpu) int {
  this.bit(0, this.de.r_high())
  return 8
}

func BIT_0_E(this *cpu) int {
  this.bit(0, this.de.r_low())
  return 8
}

func BIT_0_H(this *cpu) int {
  this.bit(0, this.hl.r_high())
  return 8
}

func BIT_0_L(this *cpu) int {
  this.bit(0, this.hl.r_low())
  return 8
}

func BIT_0_HL(this *cpu) int {
  this.bit(0, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_0_A(this *cpu) int {
  this.bit(0, this.af.r_high())
  return 8
}

func BIT_7_B(this *cpu) int {
  this.bit(7, this.bc.r_high())
  return 8
}

func BIT_7_C(this *cpu) int {
  this.bit(7, this.bc.r_low())
  return 8
}

func BIT_7_D(this *cpu) int {
  this.bit(7, this.de.r_high())
  return 8
}

func BIT_7_E(this *cpu) int {
  this.bit(7, this.de.r_low())
  return 8
}

func BIT_7_H(this *cpu) int {
  this.bit(7, this.hl.r_high())
  return 8
}

func BIT_7_L(this *cpu) int {
  this.bit(7, this.hl.r_low())
  return 8
}

func BIT_7_HL(this *cpu) int {
  this.bit(7, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_7_A(this *cpu) int {
  this.bit(7, this.af.r_high())
  return 8
}

func BIT_3_B(this *cpu) int {
  this.bit(3, this.bc.r_high())
  return 8
}

func BIT_3_C(this *cpu) int {
  this.bit(3, this.bc.r_low())
  return 8
}

func BIT_3_D(this *cpu) int {
  this.bit(3, this.de.r_high())
  return 8
}

func BIT_3_E(this *cpu) int {
  this.bit(3, this.de.r_low())
  return 8
}

func BIT_3_H(this *cpu) int {
  this.bit(3, this.hl.r_high())
  return 8
}

func BIT_3_L(this *cpu) int {
  this.bit(3, this.hl.r_low())
  return 8
}

func BIT_3_HL(this *cpu) int {
  this.bit(3, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_3_A(this *cpu) int {
  this.bit(3, this.af.r_high())
  return 8
}

func BIT_5_B(this *cpu) int {
  this.bit(5, this.bc.r_high())
  return 8
}

func BIT_5_C(this *cpu) int {
  this.bit(5, this.bc.r_low())
  return 8
}

func BIT_5_D(this *cpu) int {
  this.bit(5, this.de.r_high())
  return 8
}

func BIT_5_E(this *cpu) int {
  this.bit(5, this.de.r_low())
  return 8
}

func BIT_5_H(this *cpu) int {
  this.bit(5, this.hl.r_high())
  return 8
}

func BIT_5_L(this *cpu) int {
  this.bit(5, this.hl.r_low())
  return 8
}

func BIT_5_HL(this *cpu) int {
  this.bit(5, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_5_A(this *cpu) int {
  this.bit(5, this.af.r_high())
  return 8
}

func BIT_4_B(this *cpu) int {
  this.bit(4, this.bc.r_high())
  return 8
}

func BIT_4_C(this *cpu) int {
  this.bit(4, this.bc.r_low())
  return 8
}

func BIT_4_D(this *cpu) int {
  this.bit(4, this.de.r_high())
  return 8
}

func BIT_4_E(this *cpu) int {
  this.bit(4, this.de.r_low())
  return 8
}

func BIT_4_H(this *cpu) int {
  this.bit(4, this.hl.r_high())
  return 8
}

func BIT_4_L(this *cpu) int {
  this.bit(4, this.hl.r_low())
  return 8
}

func BIT_4_A(this *cpu) int {
  this.bit(4, this.af.r_high())
  return 8
}

func BIT_4_HL(this *cpu) int {
  this.bit(4, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_2_B(this *cpu) int {
  this.bit(2, this.bc.r_high())
  return 8
}

func BIT_2_C(this *cpu) int {
  this.bit(2, this.bc.r_low())
  return 8
}

func BIT_2_D(this *cpu) int {
  this.bit(2, this.de.r_high())
  return 8
}

func BIT_2_E(this *cpu) int {
  this.bit(2, this.de.r_low())
  return 8
}

func BIT_2_H(this *cpu) int {
  this.bit(2, this.hl.r_high())
  return 8
}

func BIT_2_L(this *cpu) int {
  this.bit(2, this.hl.r_low())
  return 8
}

func BIT_2_A(this *cpu) int {
  this.bit(2, this.af.r_high())
  return 8
}

func BIT_2_HL(this *cpu) int {
  this.bit(2, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_1_B(this *cpu) int {
  this.bit(1, (*this).bc.r_high())
  return 8
}

func BIT_1_C(this *cpu) int {
  this.bit(1, this.bc.r_low())
  return 8
}

func BIT_1_D(this *cpu) int {
  this.bit(1, this.de.r_high())
  return 8
}

func BIT_1_E(this *cpu) int {
  this.bit(1, this.de.r_low())
  return 8
}

func BIT_1_H(this *cpu) int {
  this.bit(1, this.hl.r_high())
  return 8
}

func BIT_1_L(this *cpu) int {
  this.bit(1, this.hl.r_low())
  return 8
}

func BIT_1_HL(this *cpu) int {
  this.bit(1, this.mu.Read_8(this.hl.value))
  return 16
}

func BIT_1_A(this *cpu) int {
  this.bit(1, this.af.r_high())
  return 8
}

func RRC_B(this *cpu) int {
  this.bc.w_high(this.rotate_right_carry(this.bc.r_high()))
  return 8
}

func RRC_C(this *cpu) int {
  this.bc.w_low(this.rotate_right_carry(this.bc.r_low()))
  return 8
}

func RRC_D(this *cpu) int {
  this.de.w_high(this.rotate_right_carry(this.de.r_high()))
  return 8
}

func RRC_E(this *cpu) int {
  this.de.w_low(this.rotate_right_carry(this.de.r_low()))
  return 8
}

func RRC_H(this *cpu) int {
  this.hl.w_high(this.rotate_right_carry(this.hl.r_high()))
  return 8
}

func RRC_L(this *cpu) int {
  this.hl.w_low(this.rotate_right_carry(this.hl.r_low()))
  return 8
}

func RRC_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.rotate_right_carry(this.mu.Read_8(this.hl.value)))
  return 16
}

func RRC_A(this *cpu) int {
  this.af.w_high(this.rotate_right_carry(this.af.r_high()))
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

func RES_3_B(this *cpu) int {
  this.bc.w_high(this.res(3, this.bc.r_high()))
  return 8
}

func RES_3_C(this *cpu) int {
  this.bc.w_low(this.res(3, this.bc.r_low()))
  return 8
}

func RES_3_D(this *cpu) int {
  this.de.w_high(this.res(3, this.de.r_high()))
  return 8
}

func RES_3_E(this *cpu) int {
  this.de.w_low(this.res(3, this.de.r_low()))
  return 8
}

func RES_3_H(this *cpu) int {
  this.hl.w_high(this.res(3, this.hl.r_high()))
  return 8
}

func RES_3_L(this *cpu) int {
  this.hl.w_low(this.res(3, this.hl.r_low()))
  return 8
}

func RES_3_HL(this *cpu) int {
  hl := this.mu.Read_8(this.hl.value)
  this.mu.Write_8(this.hl.value, this.res(3, hl))
  return 16
}

func RES_3_A(this *cpu) int {
  this.af.w_high(this.res(3, this.af.r_high()))
  return 8
}

func RES_7_B(this *cpu) int {
  this.bc.w_high(this.res(7, this.bc.r_high()))
  return 8
}

func RES_7_C(this *cpu) int {
  this.bc.w_low(this.res(7, this.bc.r_low()))
  return 8
}

func RES_7_D(this *cpu) int {
  this.de.w_high(this.res(7, this.de.r_high()))
  return 8
}

func RES_7_E(this *cpu) int {
  this.de.w_low(this.res(7, this.de.r_low()))
  return 8
}

func RES_7_H(this *cpu) int {
  this.hl.w_high(this.res(7, this.hl.r_high()))
  return 8
}

func RES_7_L(this *cpu) int {
  this.hl.w_low(this.res(7, this.hl.r_low()))
  return 8
}

func RES_7_HL(this *cpu) int {
  hl := this.mu.Read_8(this.hl.value)
  this.mu.Write_8(this.hl.value, this.res(7, hl))
  return 16
}

func RES_7_A(this *cpu) int {
  this.af.w_high(this.res(7, this.af.r_high()))
  return 8
}

func RES_2_B(this *cpu) int {
  this.bc.w_high(this.res(2, this.bc.r_high()))
  return 8
}

func RES_2_C(this *cpu) int {
  this.bc.w_low(this.res(2, this.bc.r_low()))
  return 8
}

func RES_2_D(this *cpu) int {
  this.de.w_high(this.res(2, this.de.r_high()))
  return 8
}

func RES_2_E(this *cpu) int {
  this.de.w_low(this.res(2, this.de.r_low()))
  return 8
}

func RES_2_H(this *cpu) int {
  this.hl.w_high(this.res(2, this.hl.r_high()))
  return 8
}

func RES_2_L(this *cpu) int {
  this.hl.w_low(this.res(2, this.hl.r_low()))
  return 8
}

func RES_2_HL(this *cpu) int {
  hl := this.mu.Read_8(this.hl.value)
  this.mu.Write_8(this.hl.value, this.res(2, hl))
  return 16
}

func RES_2_A(this *cpu) int {
  this.af.w_high(this.res(2, this.af.r_high()))
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

func RES_5_B(this *cpu) int {
  this.bc.w_high(this.res(5, this.bc.r_high()))
  return 8
}

func RES_5_C(this *cpu) int {
  this.bc.w_low(this.res(5, this.bc.r_low()))
  return 8
}

func RES_5_D(this *cpu) int {
  this.de.w_high(this.res(5, this.de.r_high()))
  return 8
}

func RES_5_E(this *cpu) int {
  this.de.w_low(this.res(5, this.de.r_low()))
  return 8
}

func RES_5_H(this *cpu) int {
  this.hl.w_high(this.res(5, this.hl.r_high()))
  return 8
}

func RES_5_L(this *cpu) int {
  this.hl.w_low(this.res(5, this.hl.r_low()))
  return 8
}

func RES_5_HL(this *cpu) int {
  hl := this.mu.Read_8(this.hl.value)
  this.mu.Write_8(this.hl.value, this.res(5, hl))
  return 16
}

func RES_5_A(this *cpu) int {
  this.af.w_high(this.res(5, this.af.r_high()))
  return 8
}

func RES_1_B(this *cpu) int {
  this.bc.w_high(this.res(1, this.bc.r_high()))
  return 8
}

func RES_1_C(this *cpu) int {
  this.bc.w_low(this.res(1, this.bc.r_low()))
  return 8
}

func RES_1_D(this *cpu) int {
  this.de.w_high(this.res(1, this.de.r_high()))
  return 8
}

func RES_1_E(this *cpu) int {
  this.de.w_low(this.res(1, this.de.r_low()))
  return 8
}

func RES_1_H(this *cpu) int {
  this.hl.w_high(this.res(1, this.hl.r_high()))
  return 8
}

func RES_1_L(this *cpu) int {
  this.hl.w_low(this.res(1, this.hl.r_low()))
  return 8
}

func RES_1_HL(this *cpu) int {
  hl := this.mu.Read_8(this.hl.value)
  this.mu.Write_8(this.hl.value, this.res(1, hl))
  return 16
}

func RES_1_A(this *cpu) int {
  this.af.w_high(this.res(1, this.af.r_high()))
  return 8
}

func SET_0_B(this *cpu) int {
  this.bc.w_high(this.set(0, this.bc.r_high()))
  return 8
}

func SET_0_C(this *cpu) int {
  this.bc.w_low(this.set(0, this.bc.r_low()))
  return 8
}

func SET_0_D(this *cpu) int {
  this.de.w_high(this.set(0, this.de.r_high()))
  return 8
}

func SET_0_E(this *cpu) int {
  this.de.w_low(this.set(0, this.de.r_low()))
  return 8
}

func SET_0_H(this *cpu) int {
  this.hl.w_high(this.set(0, this.hl.r_high()))
  return 8
}

func SET_0_L(this *cpu) int {
  this.hl.w_low(this.set(0, this.hl.r_low()))
  return 8
}

func SET_0_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.set(0, this.mu.Read_8(this.hl.value)))
  return 16
}

func SET_0_A(this *cpu) int {
  this.af.w_high(this.set(0, this.af.r_high()))
  return 8
}

func SET_5_B(this *cpu) int {
  this.bc.w_high(this.set(5, this.bc.r_high()))
  return 8
}

func SET_5_C(this *cpu) int {
  this.bc.w_low(this.set(5, this.bc.r_low()))
  return 8
}

func SET_5_D(this *cpu) int {
  this.de.w_high(this.set(5, this.de.r_high()))
  return 8
}

func SET_5_E(this *cpu) int {
  this.de.w_low(this.set(5, this.de.r_low()))
  return 8
}

func SET_5_H(this *cpu) int {
  this.hl.w_high(this.set(5, this.hl.r_high()))
  return 8
}

func SET_5_L(this *cpu) int {
  this.hl.w_low(this.set(5, this.hl.r_low()))
  return 8
}

func SET_5_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.set(5, this.mu.Read_8(this.hl.value)))
  return 16
}

func SET_5_A(this *cpu) int {
  this.af.w_high(this.set(5, this.af.r_high()))
  return 8
}

func SET_1_B(this *cpu) int {
  this.bc.w_high(this.set(1, this.bc.r_high()))
  return 8
}

func SET_1_C(this *cpu) int {
  this.bc.w_low(this.set(1, this.bc.r_low()))
  return 8
}

func SET_1_D(this *cpu) int {
  this.de.w_high(this.set(1, this.de.r_high()))
  return 8
}

func SET_1_E(this *cpu) int {
  this.de.w_low(this.set(1, this.de.r_low()))
  return 8
}

func SET_1_H(this *cpu) int {
  this.hl.w_high(this.set(1, this.hl.r_high()))
  return 8
}

func SET_1_L(this *cpu) int {
  this.hl.w_low(this.set(1, this.hl.r_low()))
  return 8
}

func SET_1_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.set(1, this.mu.Read_8(this.hl.value)))
  return 16
}

func SET_1_A(this *cpu) int {
  this.af.w_high(this.set(1, this.af.r_high()))
  return 8
}

func SET_3_B(this *cpu) int {
  this.bc.w_high(this.set(3, this.bc.r_high()))
  return 8
}

func SET_3_C(this *cpu) int {
  this.bc.w_low(this.set(3, this.bc.r_low()))
  return 8
}

func SET_3_D(this *cpu) int {
  this.de.w_high(this.set(3, this.de.r_high()))
  return 8
}

func SET_3_E(this *cpu) int {
  this.de.w_low(this.set(3, this.de.r_low()))
  return 8
}

func SET_3_H(this *cpu) int {
  this.hl.w_high(this.set(3, this.hl.r_high()))
  return 8
}

func SET_3_L(this *cpu) int {
  this.hl.w_low(this.set(3, this.hl.r_low()))
  return 8
}

func SET_3_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.set(3, this.mu.Read_8(this.hl.value)))
  return 16
}

func SET_3_A(this *cpu) int {
  this.af.w_high(this.set(3, this.af.r_high()))
  return 8
}

func SET_7_B(this *cpu) int {
  this.bc.w_high(this.set(7, this.bc.r_high()))
  return 8
}

func SET_7_C(this *cpu) int {
  this.bc.w_low(this.set(7, this.bc.r_low()))
  return 8
}

func SET_7_D(this *cpu) int {
  this.de.w_high(this.set(7, this.de.r_high()))
  return 8
}

func SET_7_E(this *cpu) int {
  this.de.w_low(this.set(7, this.de.r_low()))
  return 8
}

func SET_7_H(this *cpu) int {
  this.hl.w_high(this.set(7, this.hl.r_high()))
  return 8
}

func SET_7_L(this *cpu) int {
  this.hl.w_low(this.set(7, this.hl.r_low()))
  return 8
}

func SET_7_HL(this *cpu) int {
  val := this.set(7, this.mu.Read_8(this.hl.value))
  this.mu.Write_8(this.hl.value, val)
  return 16
}

func SET_7_A(this *cpu) int {
  this.af.w_high(this.set(7, this.af.r_high()))
  return 8
}

func SLA_B(this *cpu) int {
  this.bc.w_high(this.shift_left_carry(this.bc.r_high()))
  return 8
}

func SLA_C(this *cpu) int {
  this.bc.w_low(this.shift_left_carry(this.bc.r_low()))
  return 8
}

func SLA_D(this *cpu) int {
  this.de.w_high(this.shift_left_carry(this.de.r_high()))
  return 8
}

func SLA_E(this *cpu) int {
  this.de.w_low(this.shift_left_carry(this.de.r_low()))
  return 8
}

func SLA_H(this *cpu) int {
  this.hl.w_high(this.shift_left_carry(this.hl.r_high()))
  return 8
}

func SLA_L(this *cpu) int {
  this.hl.w_low(this.shift_left_carry(this.hl.r_low()))
  return 8
}

func SLA_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.shift_left_carry(this.mu.Read_8(this.hl.value)))
  return 16
}

func SET_2_B(this *cpu) int {
  this.bc.w_high(this.set(2, this.bc.r_high()))
  return 8
}

func SET_2_C(this *cpu) int {
  this.bc.w_low(this.set(2, this.bc.r_low()))
  return 8
}

func SET_2_D(this *cpu) int {
  this.de.w_high(this.set(2, this.de.r_high()))
  return 8
}

func SET_2_E(this *cpu) int {
  this.de.w_low(this.set(2, this.de.r_low()))
  return 8
}

func SET_2_H(this *cpu) int {
  this.hl.w_high(this.set(2, this.hl.r_high()))
  return 8
}

func SET_2_L(this *cpu) int {
  this.hl.w_low(this.set(2, this.hl.r_low()))
  return 8
}

func SET_2_HL(this *cpu) int {
  this.mu.Write_8(this.hl.value, this.set(2, this.mu.Read_8(this.hl.value)))
  return 16
}

func SET_2_A(this *cpu) int {
  this.af.w_high(this.set(2, this.af.r_high()))
  return 8
}

func SLA_A(this *cpu) int {
  this.af.w_high(this.shift_left_carry(this.af.r_high()))
  return 8
}

func (this *cpu) do_cb_op(op uint8) int {
  switch(op) {
  case 0x08:
    return RRC_B(this)
  case 0x09:
    return RRC_C(this)
  case 0x0a:
    return RRC_D(this)
  case 0x0b:
    return RRC_E(this)
  case 0x0c:
    return RRC_H(this)
  case 0x0d:
    return RRC_L(this)
  case 0x0e:
    return RRC_HL(this)
  case 0x0f:
    return RRC_A(this)
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
  case 0x20:
    return SLA_B(this)
  case 0x21:
    return SLA_C(this)
  case 0x22:
    return SLA_D(this)
  case 0x23:
    return SLA_E(this)
  case 0x24:
    return SLA_H(this)
  case 0x25:
    return SLA_L(this)
  case 0x26:
    return SLA_HL(this)
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
  case 0x40:
    return BIT_0_B(this)
  case 0x41:
    return BIT_0_C(this)
  case 0x42:
    return BIT_0_D(this)
  case 0x43:
    return BIT_0_E(this)
  case 0x44:
    return BIT_0_H(this)
  case 0x45:
    return BIT_0_L(this)
  case 0x46:
    return BIT_0_HL(this)
  case 0x47:
    return BIT_0_A(this)
  case 0x48:
    return BIT_1_B(this)
  case 0x49:
    return BIT_1_C(this)
  case 0x4a:
    return BIT_1_D(this)
  case 0x4b:
    return BIT_1_E(this)
  case 0x4c:
    return BIT_1_H(this)
  case 0x4d:
    return BIT_1_L(this)
  case 0x4e:
    return BIT_1_HL(this)
  case 0x4f:
    return BIT_1_A(this)
  case 0x50:
    return BIT_2_B(this)
  case 0x51:
    return BIT_2_C(this)
  case 0x52:
    return BIT_2_D(this)
  case 0x53:
    return BIT_2_E(this)
  case 0x54:
    return BIT_2_H(this)
  case 0x55:
    return BIT_2_L(this)
  case 0x56:
    return BIT_2_HL(this)
  case 0x57:
    return BIT_2_A(this)
  case 0x58:
    return BIT_3_B(this)
  case 0x59:
    return BIT_3_C(this)
  case 0x5a:
    return BIT_3_D(this)
  case 0x5b:
    return BIT_3_E(this)
  case 0x5c:
    return BIT_3_H(this)
  case 0x5d:
    return BIT_3_L(this)
  case 0x5e:
    return BIT_3_HL(this)
  case 0x5f:
    return BIT_3_A(this)
  case 0x60:
    return BIT_4_B(this)
  case 0x61:
    return BIT_4_C(this)
  case 0x62:
    return BIT_4_D(this)
  case 0x63:
    return BIT_4_E(this)
  case 0x64:
    return BIT_4_H(this)
  case 0x65:
    return BIT_4_L(this)
  case 0x66:
    return BIT_4_HL(this)
  case 0x67:
    return BIT_4_A(this)
  case 0x68:
    return BIT_5_B(this)
  case 0x69:
    return BIT_5_C(this)
  case 0x6a:
    return BIT_5_D(this)
  case 0x6b:
    return BIT_5_E(this)
  case 0x6c:
    return BIT_5_H(this)
  case 0x6d:
    return BIT_5_L(this)
  case 0x6e:
    return BIT_5_HL(this)
  case 0x6f:
    return BIT_5_A(this)
  case 0x70:
    return BIT_6_B(this)
  case 0x71:
    return BIT_6_C(this)
  case 0x72:
    return BIT_6_D(this)
  case 0x73:
    return BIT_6_E(this)
  case 0x74:
    return BIT_6_H(this)
  case 0x75:
    return BIT_6_L(this)
  case 0x76:
    return BIT_6_HL(this)
  case 0x77:
    return BIT_6_A(this)
  case 0x78:
    return BIT_7_B(this)
  case 0x79:
    return BIT_7_C(this)
  case 0x7a:
    return BIT_7_D(this)
  case 0x7b:
    return BIT_7_E(this)
  case 0x7c:
    return BIT_7_H(this)
  case 0x7d:
    return BIT_7_L(this)
  case 0x7e:
    return BIT_7_HL(this)
  case 0x7f:
    return BIT_7_A(this)
  case 0x86:
    return RES_0_HL(this)
  case 0x87:
    return RES_0_A(this)
  case 0x88:
    return RES_1_B(this)
  case 0x89:
    return RES_1_C(this)
  case 0x8a:
    return RES_1_D(this)
  case 0x8b:
    return RES_1_E(this)
  case 0x8c:
    return RES_1_H(this)
  case 0x8d:
    return RES_1_L(this)
  case 0x8e:
    return RES_1_HL(this)
  case 0x8f:
    return RES_1_A(this)
  case 0x90:
    return RES_2_B(this)
  case 0x91:
    return RES_2_C(this)
  case 0x92:
    return RES_2_D(this)
  case 0x93:
    return RES_2_E(this)
  case 0x94:
    return RES_2_H(this)
  case 0x95:
    return RES_2_L(this)
  case 0x96:
    return RES_2_HL(this)
  case 0x97:
    return RES_2_A(this)
  case 0x98:
    return RES_3_B(this)
  case 0x99:
    return RES_3_C(this)
  case 0x9a:
    return RES_3_D(this)
  case 0x9b:
    return RES_3_E(this)
  case 0x9c:
    return RES_3_H(this)
  case 0x9d:
    return RES_3_L(this)
  case 0x9e:
    return RES_3_HL(this)
  case 0x9f:
    return RES_3_A(this)
  case 0xa8:
    return RES_5_B(this)
  case 0xa9:
    return RES_5_C(this)
  case 0xaa:
    return RES_5_D(this)
  case 0xab:
    return RES_5_E(this)
  case 0xac:
    return RES_5_H(this)
  case 0xad:
    return RES_5_L(this)
  case 0xae:
    return RES_5_HL(this)
  case 0xaf:
    return RES_5_A(this)
  case 0xb8:
    return RES_7_B(this)
  case 0xb9:
    return RES_7_C(this)
  case 0xba:
    return RES_7_D(this)
  case 0xbb:
    return RES_7_E(this)
  case 0xbc:
    return RES_7_H(this)
  case 0xbd:
    return RES_7_L(this)
  case 0xbe:
    return RES_7_HL(this)
  case 0xbf:
    return RES_7_A(this)
  case 0xc0:
    return SET_0_B(this)
  case 0xc1:
    return SET_0_C(this)
  case 0xc2:
    return SET_0_D(this)
  case 0xc3:
    return SET_0_E(this)
  case 0xc4:
    return SET_0_H(this)
  case 0xc5:
    return SET_0_L(this)
  case 0xc6:
    return SET_0_HL(this)
  case 0xc7:
    return SET_0_A(this)
  case 0xc8:
    return SET_1_B(this)
  case 0xc9:
    return SET_1_C(this)
  case 0xca:
    return SET_1_D(this)
  case 0xcb:
    return SET_1_E(this)
  case 0xcc:
    return SET_1_H(this)
  case 0xcd:
    return SET_1_L(this)
  case 0xce:
    return SET_1_HL(this)
  case 0xcf:
    return SET_1_A(this)
  case 0xd0:
    return SET_2_B(this)
  case 0xd1:
    return SET_2_C(this)
  case 0xd2:
    return SET_2_D(this)
  case 0xd3:
    return SET_2_E(this)
  case 0xd4:
    return SET_2_H(this)
  case 0xd5:
    return SET_2_L(this)
  case 0xd6:
    return SET_2_HL(this)
  case 0xd7:
    return SET_2_A(this)
  case 0xd8:
    return SET_3_B(this)
  case 0xd9:
    return SET_3_C(this)
  case 0xda:
    return SET_3_D(this)
  case 0xdb:
    return SET_3_E(this)
  case 0xdc:
    return SET_3_H(this)
  case 0xdd:
    return SET_3_L(this)
  case 0xde:
    return SET_3_HL(this)
  case 0xdf:
    return SET_3_A(this)
  case 0xe8:
    return SET_5_B(this)
  case 0xe9:
    return SET_5_C(this)
  case 0xea:
    return SET_5_D(this)
  case 0xeb:
    return SET_5_E(this)
  case 0xec:
    return SET_5_H(this)
  case 0xed:
    return SET_5_L(this)
  case 0xee:
    return SET_5_HL(this)
  case 0xef:
    return SET_5_A(this)
  case 0xf8:
    return SET_7_B(this)
  case 0xf9:
    return SET_7_C(this)
  case 0xfa:
    return SET_7_D(this)
  case 0xfb:
    return SET_7_E(this)
  case 0xfc:
    return SET_7_H(this)
  case 0xfd:
    return SET_7_L(this)
  case 0xfe:
    return SET_7_HL(this)
  case 0xff:
    return SET_7_A(this)
  default:
    return -1
  }
}
