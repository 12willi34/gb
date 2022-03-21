package gb

import (
  //"fmt"
)

func NOOP(this *cpu) int {
  return 4
}

func DEC_HLadr(this *cpu) int {
  hl := (*this).hl.value
  val := (*(*this).mu).Read_8(hl)
  (*(*this).mu).Write_8(hl, this.decrement(val))
  return 12
}

func DEC_BC(this *cpu) int {
  (*this).bc.value -= 1
  return 8
}

func DEC_DE(this *cpu) int {
  (*this).de.value -= 1
  return 8
}

func DEC_HL(this *cpu) int {
  (*this).hl.value -= 1
  return 8
}

func DEC_SP(this *cpu) int {
  (*this).sp.value -= 1
  return 8
}

func DEC_B(this *cpu) int {
  (*this).bc.w_high(this.decrement((*this).bc.r_high()))
  return 4
}

func DEC_A(this *cpu) int {
  (*this).af.w_high(this.decrement((*this).af.r_high()))
  return 4
}

func DEC_C(this *cpu) int {
  (*this).bc.w_low(this.decrement((*this).bc.r_low()))
  return 4
}

func DEC_D(this *cpu) int {
  (*this).de.w_high(this.decrement((*this).de.r_high()))
  return 4
}

func DEC_E(this *cpu) int {
  (*this).de.w_low(this.decrement((*this).de.r_low()))
  return 4
}

func INC_A(this *cpu) int {
  x := this.increment((*this).af.r_high())
  (*this).af.w_high(x)
  return 4
}

func INC_B(this *cpu) int {
  x := this.increment((*this).bc.r_high())
  (*this).bc.w_high(x)
  return 4
}

func INC_C(this *cpu) int {
  x := this.increment((*this).bc.r_low())
  (*this).bc.w_low(x)
  return 4
}

func INC_D(this *cpu) int {
  x := this.increment((*this).de.r_high())
  (*this).de.w_high(x)
  return 4
}

func INC_E(this *cpu) int {
  x := this.increment((*this).de.r_low())
  (*this).de.w_low(x)
  return 4
}

func INC_H(this *cpu) int {
  x := this.increment((*this).hl.r_high())
  (*this).hl.w_high(x)
  return 4
}

func INC_L(this *cpu) int {
  x := this.increment((*this).hl.r_low())
  (*this).hl.w_low(x)
  return 4
}

func INC_BC(this *cpu) int {
  (*this).bc.value += 1
  return 8
}

func INC_DE(this *cpu) int {
  (*this).de.value += 1
  return 8
}

func INC_HL(this *cpu) int {
  (*this).hl.value += 1
  return 8
}

func INC_SP(this *cpu) int {
  (*this).sp.value += 1
  return 8
}

func INC_HL_mem(this *cpu) int {
  val_before := (*(*this).mu).Read_8((*this).hl.value)
  val_after := this.increment(val_before)
  (*(*this).mu).Write_8((*this).hl.value, val_after)
  return 12
}

func LD_A_HLI(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  (*this).af.w_high(val)
  (*this).hl.value += 1
  return 8
}

func LD_HL_number(this *cpu) int {
  val := this.fetch()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 12
}

func LD_number_SP(this *cpu) int {
  adr := this.fetch_16()
  (*(*this).mu).Write_8(adr, (*this).sp.r_low())
  (*(*this).mu).Write_8(adr + 1, (*this).sp.r_high())
  return 20
}

func LDD_A_HL(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  (*this).af.w_high(val)
  (*this).hl.value -= 1
  return 8
}

func LD_D_HL(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  (*this).de.w_high(val)
  return 8
}

func LD_E_B(this *cpu) int {
  (*this).de.w_low((*this).bc.r_high())
  return 4
}

func LD_E_C(this *cpu) int {
  (*this).de.w_low((*this).bc.r_low())
  return 4
}

func LD_C_L(this *cpu) int {
  (*this).bc.w_low((*this).hl.r_low())
  return 4
}

func LD_C_HL(this *cpu) int {
  (*this).bc.w_low((*(*this).mu).Read_8((*this).hl.value))
  return 8
}

func LD_E_E(this *cpu) int {
  (*this).de.w_low((*this).de.r_low())
  return 4
}

func LD_E_D(this *cpu) int {
  (*this).de.w_low((*this).de.r_high())
  return 4
}

func LD_E_H(this *cpu) int {
  (*this).de.w_low((*this).hl.r_high())
  return 4
}

func LD_E_L(this *cpu) int {
  (*this).de.w_low((*this).hl.r_low())
  return 4
}

func LD_E_HL(this *cpu) int {
  (*this).de.w_low((*(*this).mu).Read_8((*this).hl.value))
  return 8
}

func LD_D_B(this *cpu) int {
  (*this).de.w_high((*this).bc.r_high())
  return 4
}

func LD_SP_HL(this *cpu) int {
  (*this).sp.value = (*this).hl.value
  return 8
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
  (*this).af.w_high((*this).mu.Read_8((*this).hl.value))
  return 8
}

func LD_B_B(this *cpu) int {
  (*this).bc.w_high((*this).bc.r_high())
  return 4
}

func LD_B_A(this *cpu) int {
  (*this).bc.w_high((*this).af.r_high())
  return 4
}

func LD_C_A(this *cpu) int {
  (*this).bc.w_low((*this).af.r_high())
  return 4
}

func LD_D_A(this *cpu) int {
  (*this).de.w_high((*this).af.r_high())
  return 4
}

func LD_E_A(this *cpu) int {
  (*this).de.w_low((*this).af.r_high())
  return 4
}

func LD_L_A(this *cpu) int {
  (*this).hl.w_low((*this).af.r_high())
  return 4
}

func LD_n_A(this *cpu) int {
  i := 0xff00 + uint16(this.fetch())
  (*this.mu).Write_8(i, this.af.r_high())
  return 12
}

func LD_A_n(this *cpu) int {
  a := (*this.mu).Read_8(0xff00 + uint16(this.fetch()))
  this.af.w_high(a)
  return 12
}

func LD_A_BC(this *cpu) int {
  val := (*this).bc.value
  (*this).af.w_high((*this).mu.Read_8(val))
  return 8
}

func LD_A_DE(this *cpu) int {
  val := (*this).de.value
  (*this).af.w_high((*this).mu.Read_8(val))
  return 8
}

func LD_A_nn(this *cpu) int {
  val := this.fetch_16()
  (*this).af.w_high((*this).mu.Read_8(val))
  return 16
}

func LD_A_number(this *cpu) int {
  (*this).af.w_high(this.fetch())
  return 8
}

func LD_SP_nn(this *cpu) int {
  (*this).sp.value = this.fetch_16()
  return 12
}

func LD_HL_nn(this *cpu) int {
  (*this).hl.value = this.fetch_16()
  return 12
}

func LD_DE_nn(this *cpu) int {
  (*this).de.value = this.fetch_16()
  return 12
}

func LD_BC_nn(this *cpu) int {
  (*this).bc.value = this.fetch_16()
  return 12
}

func LD_C_n(this *cpu) int {
  (*this).bc.w_low(this.fetch())
  return 8
}

func LD_D_n(this *cpu) int {
  (*this).de.w_high(this.fetch())
  return 8
}

func LD_E_n(this *cpu) int {
  (*this).de.w_low(this.fetch())
  return 8
}

func LD_H_n(this *cpu) int {
  (*this).hl.w_high(this.fetch())
  return 8
}

func LD_H_A(this *cpu) int {
  (*this).hl.w_high((*this).af.r_high())
  return 4
}

func LD_L_n(this *cpu) int {
  (*this).hl.w_low(this.fetch())
  return 8
}

func LD_A_ff00_C(this *cpu) int {
  i := 0xff00 + uint16((*this).bc.r_low())
  val := (*(*this).mu).Read_8(i)
  (*this).af.w_high(val)
  return 8
}

func LD_B_D(this *cpu) int {
  (*this).bc.w_high((*this).de.r_high())
  return 4
}

func LD_ff00_C_A(this *cpu) int {
  a := (*this).af.r_high()
  i := 0xff00 + uint16((*this).bc.r_low())
  (*(*this).mu).Write_8(i, a)
  return 8
}

func LD_HL_A(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).hl.value
  (*(*this).mu).Write_8(b, a)
  return 8
}

func LD_HL_B(this *cpu) int {
  val := (*this).bc.r_high()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 8
}

func LD_HL_C(this *cpu) int {
  val := (*this).bc.r_low()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 8
}

func LD_HL_D(this *cpu) int {
  val := (*this).de.r_high()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 8
}

func LD_HL_E(this *cpu) int {
  val := (*this).de.r_low()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 8
}

func LD_HL_H(this *cpu) int {
  val := (*this).hl.r_high()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 8
}

func LD_HL_L(this *cpu) int {
  val := (*this).hl.r_low()
  (*(*this).mu).Write_8((*this).hl.value, val)
  return 8
}

func LDI_HL_A(this *cpu) int {
  a := (*this).af.r_high()
  hl := (*this).hl.value
  (*(*this).mu).Write_8(hl, a)
  (*this).hl.value = hl + 1
  return 8
}

func LD_BC_A(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).bc.value
  (*(*this).mu).Write_8(b, a)
  return 8
}

func LD_H_B(this *cpu) int {
  (*this).hl.w_high((*this).bc.r_high())
  return 4
}

func LD_H_C(this *cpu) int {
  (*this).hl.w_high((*this).bc.r_low())
  return 4
}

func LD_H_D(this *cpu) int {
  (*this).hl.w_high((*this).de.r_high())
  return 4
}

func LD_H_E(this *cpu) int {
  (*this).hl.w_high((*this).de.r_low())
  return 4
}

func LD_H_L(this *cpu) int {
  (*this).hl.w_high((*this).hl.r_low())
  return 4
}

func LD_H_H(this *cpu) int {
  (*this).hl.w_high((*this).hl.r_high())
  return 4
}

func LD_H_HL(this *cpu) int {
  (*this).hl.w_high((*(*this).mu).Read_8((*this).hl.value))
  return 8
}

func LD_DE_A(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).de.value
  (*(*this).mu).Write_8(b, a)
  return 8
}

func LD_D_L(this *cpu) int {
  (*this).de.w_high((*this).hl.r_low())
  return 4
}

func LD_nn_A(this *cpu) int {
  a := (*this).af.r_high()
  b := this.fetch_16()
  (*(*this).mu).Write_8(b, a)
  return 16
}

func LD_L_B(this *cpu) int {
  (*this).hl.w_low((*this).bc.r_high())
  return 4
}

func LD_L_C(this *cpu) int {
  (*this).hl.w_low((*this).bc.r_low())
  return 4
}

func LD_L_D(this *cpu) int {
  (*this).hl.w_low((*this).de.r_high())
  return 4
}

func LD_L_E(this *cpu) int {
  (*this).hl.w_low((*this).de.r_low())
  return 4
}

func LD_L_H(this *cpu) int {
  (*this).hl.w_low((*this).hl.r_high())
  return 4
}

func LD_L_L(this *cpu) int {
  (*this).hl.w_low((*this).hl.r_low())
  return 4
}

func LD_L_HL(this *cpu) int {
  (*this).hl.w_low((*(*this).mu).Read_8((*this).hl.value))
  return 8
}

func LDD_HL_A(this *cpu) int {
  hl := (*this).hl.value
  (*(*this).mu).Write_8(hl, (*this).af.r_high())
  (*this).hl.value = hl - 1
  return 8
}

func LD_B_C(this *cpu) int {
  (*this).bc.w_high((*this).bc.r_low())
  return 4
}

func LD_B_E(this *cpu) int {
  (*this).bc.w_high((*this).de.r_low())
  return 4
}

func LD_B_H(this *cpu) int {
  (*this).bc.w_high((*this).hl.r_high())
  return 4
}

func LD_B_L(this *cpu) int {
  (*this).bc.w_high((*this).hl.r_low())
  return 4
}

func LD_B_HL(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  (*this).bc.w_high(val)
  return 8
}

func LD_C_B(this *cpu) int {
  (*this).bc.w_low((*this).bc.r_high())
  return 4
}

func LD_C_C(this *cpu) int {
  (*this).bc.w_low((*this).bc.r_low())
  return 4
}

func LD_C_D(this *cpu) int {
  (*this).bc.w_low((*this).de.r_high())
  return 4
}

func LD_D_C(this *cpu) int {
  (*this).de.w_high((*this).bc.r_low())
  return 4
}

func LD_D_D(this *cpu) int {
  (*this).de.w_high((*this).de.r_high())
  return 4
}

func LD_D_E(this *cpu) int {
  (*this).de.w_high((*this).de.r_low())
  return 4
}

func LD_D_H(this *cpu) int {
  (*this).de.w_high((*this).hl.r_high())
  return 4
}

func LD_C_E(this *cpu) int {
  (*this).bc.w_low((*this).de.r_low())
  return 4
}

func LD_C_H(this *cpu) int {
  (*this).bc.w_low((*this).hl.r_high())
  return 4
}

func JR_n(this *cpu) int {
  a := int8(this.fetch())
  (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
  return 8
}

func JR_nZ(this *cpu) int {
  a := int8(this.fetch())
  if(!this.get_f_zero()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
  }
  return 8
}

func JR_Z(this *cpu) int {
  a := int8(this.fetch())
  if(this.get_f_zero()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
  }
  return 8
}

func JR_nC(this *cpu) int {
  a := int8(this.fetch())
  if(!this.get_f_carry()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
  }
  return 8
}

func JR_C(this *cpu) int {
  a := int8(this.fetch())
  if(this.get_f_carry()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
  }
  return 8
}

func OR_A_A(this *cpu) int {
  val := (*this).af.r_high()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_B(this *cpu) int {
  val := (*this).bc.r_high()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_C(this *cpu) int {
  val := (*this).bc.r_low()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_D(this *cpu) int {
  val := (*this).de.r_high()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_E(this *cpu) int {
  val := (*this).de.r_low()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_H(this *cpu) int {
  val := (*this).hl.r_high()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_L(this *cpu) int {
  val := (*this).hl.r_low()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 4
}

func OR_A_HL(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 8
}

func OR_A_number(this *cpu) int {
  val := this.fetch()
  (*this).af.w_high(this.or((*this).af.r_high(), val))
  return 8
}

func JP(this *cpu) int {
  this.pc.value = this.fetch_16()
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

func CP_HL_memory(this *cpu) int {
  this.compare_8(this.af.r_high(), (*(*this).mu).Read_8((*this).hl.value))
  return 8
}

func AND_A_A(this *cpu) int {
  b := (*this).af.r_high()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_B(this *cpu) int {
  b := (*this).bc.r_high()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_C(this *cpu) int {
  b := (*this).bc.r_low()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_D(this *cpu) int {
  b := (*this).de.r_high()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_E(this *cpu) int {
  b := (*this).de.r_low()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_H(this *cpu) int {
  b := (*this).hl.r_high()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_L(this *cpu) int {
  b := (*this).hl.r_low()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 4
}

func AND_A_HL(this *cpu) int {
  b := (*(*this).mu).Read_8((*this).hl.value)
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 8
}

func AND_A_number(this *cpu) int {
  b := this.fetch()
  (*this).af.w_high((*this).and((*this).af.r_high(), b))
  return 8
}

func ADD_HL_SP(this *cpu) int {
  x := (*this).add_16((*this).hl.value, (*this).sp.value)
  (*this).hl.value = x
  return 8
}

func ADD_HL_HL(this *cpu) int {
  x := (*this).add_16((*this).hl.value, (*this).hl.value)
  (*this).hl.value = x
  return 8
}

func ADD_HL_DE(this *cpu) int {
  x := (*this).add_16((*this).hl.value, (*this).de.value)
  (*this).hl.value = x
  return 8
}

func ADD_HL_BC(this *cpu) int {
  x := (*this).add_16((*this).hl.value, (*this).bc.value)
  (*this).hl.value = x
  return 8
}

func ADD_A_HL_memory(this *cpu) int {
  val := (*(*this).mu).Read_8((*this).hl.value)
  a := this.add((*this).af.r_high(), val)
  (*this).af.w_high(a)
  return 8
}

func ADD_A_number(this *cpu) int {
  a := this.add((*this).af.r_high(), this.fetch())
  (*this).af.w_high(a)
  return 8
}

func ADD_A_L(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).hl.r_low())
  (*this).af.w_high(a)
  return 4
}

func ADD_A_H(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).hl.r_high())
  (*this).af.w_high(a)
  return 4
}

func ADD_A_E(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).de.r_low())
  (*this).af.w_high(a)
  return 4
}

func ADD_A_D(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).de.r_high())
  (*this).af.w_high(a)
  return 4
}

func ADD_A_C(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).bc.r_low())
  (*this).af.w_high(a)
  return 4
}

func ADD_A_B(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).bc.r_high())
  (*this).af.w_high(a)
  return 4
}

func ADD_A_A(this *cpu) int {
  a := this.add((*this).af.r_high(), (*this).af.r_high())
  (*this).af.w_high(a)
  return 4
}

func SUB_L(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).hl.r_low())
  (*this).af.w_high(a)
  return 4
}

func SUB_H(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).hl.r_high())
  (*this).af.w_high(a)
  return 4
}

func SUB_E(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).de.r_low())
  (*this).af.w_high(a)
  return 4
}

func SUB_D(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).de.r_high())
  (*this).af.w_high(a)
  return 4
}

func SUB_C(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).bc.r_low())
  (*this).af.w_high(a)
  return 4
}

func SUB_B(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).bc.r_high())
  (*this).af.w_high(a)
  return 4
}

func SUB_A(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*this).af.r_high())
  (*this).af.w_high(a)
  return 4
}

func SUB_HL_memory(this *cpu) int {
  a := this.subtract((*this).af.r_high(), (*(*this).mu).Read_8((*this).hl.value))
  (*this).af.w_high(a)
  return 8
}

func SUB_number(this *cpu) int {
  a := this.subtract((*this).af.r_high(), this.fetch())
  (*this).af.w_high(a)
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

func SBC_A_number(this *cpu) int {
  a := (*this).af.r_high()
  b := this.fetch()
  (*this).af.w_high(this.subtract_carry(a, b))
  return 8
}

func SBC_A_HL(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).mu.Read_8((*this).hl.value)
  (*this).af.w_high(this.subtract_carry(a, b))
  return 8
}

func RLA(this *cpu) int {
  res := this.rla((*this).af.r_high())
  (*this).af.w_high(res)
  return 4
}

func RST_38(this *cpu) int {
  this.restart(0x38)
  return 32
}

func RST_30(this *cpu) int {
  this.restart(0x30)
  return 32
}

func RST_28(this *cpu) int {
  this.restart(0x28)
  return 32
}

func RST_20(this *cpu) int {
  this.restart(0x20)
  return 32
}

func RST_18(this *cpu) int {
  this.restart(0x18)
  return 32
}

func RST_10(this *cpu) int {
  this.restart(0x10)
  return 32
}

func RST_08(this *cpu) int {
  this.restart(0x08)
  return 32
}

func RST_00(this *cpu) int {
  this.restart(0x00)
  return 32
}

func XOR_A(this *cpu) int {
  a := (*this).af.r_high()
  (*this).af.w_high(this.xor(a, a))
  return 4
}

func XOR_B(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).bc.r_high()
  (*this).af.w_high(this.xor(a, b))
  return 4
}

func XOR_C(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).bc.r_low()
  (*this).af.w_high(this.xor(a, b))
  return 4
}

func XOR_D(this *cpu) int {
  t := *this
  t.af.w_high(this.xor(t.af.r_high(), t.de.r_high()))
  return 4
}

func XOR_E(this *cpu) int {
  t := *this
  t.af.w_high(this.xor(t.af.r_high(), t.de.r_low()))
  return 4
}

func XOR_H(this *cpu) int {
  t := *this
  t.af.w_high(this.xor(t.af.r_high(), t.hl.r_high()))
  return 4
}

func XOR_L(this *cpu) int {
  t := *this
  t.af.w_high(this.xor(t.af.r_high(), t.hl.r_low()))
  return 4
}

func XOR_HL(this *cpu) int {
  t := *this
  hl := (*(t.mu)).Read_8(t.hl.value)
  t.af.w_high(this.xor(t.af.r_high(), hl))
  return 8
}

func XOR_number(this *cpu) int {
  t := *this
  n := this.fetch()
  t.af.w_high(this.xor(t.af.r_high(), n))
  return 8
}

func POP_BC(this *cpu) int {
  (*this).bc.value = this.popStack()
  return 12
}

func POP_AF(this *cpu) int {
  (*this).af.value = this.popStack()
  return 12
}

func POP_DE(this *cpu) int {
  (*this).de.value = this.popStack()
  return 12
}

func ADC_A_A(this *cpu) int {
  val := (*this).af.r_high()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_B(this *cpu) int {
  val := (*this).bc.r_high()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_C(this *cpu) int {
  val := (*this).bc.r_low()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_D(this *cpu) int {
  val := (*this).de.r_high()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_E(this *cpu) int {
  val := (*this).de.r_low()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_H(this *cpu) int {
  val := (*this).hl.r_high()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_L(this *cpu) int {
  val := (*this).hl.r_low()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 4
}

func ADC_A_HL(this *cpu) int {
  val := ((*this).mu).Read_8((*this).hl.value)
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 8
}

func ADC_A_number(this *cpu) int {
  val := this.fetch()
  (*this).af.w_high(this.adc((*this).af.r_high(), val))
  return 8
}

func POP_HL(this *cpu) int {
  (*this).hl.value = this.popStack()
  return 12
}

func CB_OP(this *cpu) int {
  return ((*this).cb_ops[this.fetch()])(this)
}

func CP_A_A(this *cpu) int {
  this.compare_8(this.af.r_high(), this.af.r_high())
  return 4
}

func CP_A_B(this *cpu) int {
  this.compare_8(this.af.r_high(), this.bc.r_high())
  return 4
}

func CP_A_C(this *cpu) int {
  this.compare_8(this.af.r_high(), this.bc.r_low())
  return 4
}

func CP_A_D(this *cpu) int {
  this.compare_8(this.af.r_high(), this.de.r_high())
  return 4
}

func CP_A_E(this *cpu) int {
  this.compare_8(this.af.r_high(), this.de.r_low())
  return 4
}

func CP_A_H(this *cpu) int {
  this.compare_8(this.af.r_high(), this.hl.r_high())
  return 4
}

func CP_A_L(this *cpu) int {
  this.compare_8(this.af.r_high(), this.hl.r_low())
  return 4
}

func RETI(this *cpu) int {
  (*this).pc.value = this.popStack()
  (*this).Interrupt = true
  return 8
}

func CCF(this *cpu) int {
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(!this.get_f_carry())
  return 4
}

func RRCA(this *cpu) int {
  val := (*this).af.r_high()
  res := uint8(val >> 1) | uint8((val & 1) << 7)
  (*this).af.w_high(res)
  this.set_f_zero(false)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(res > 0x7f)
  return 4
}

func ADD_SP_number(this *cpu) int {
  sp := (*this).sp.value
  num := int8(this.fetch())
  res := uint16(int32(sp) + int32(num))
  val := sp ^ uint16(num) ^ res
  this.set_f_zero(false)
  this.set_f_subtr(false)
  this.set_f_h_carry((val & 0x10) == 0x10)
  this.set_f_carry((val & 0x100) == 0x100)
  (*this).sp.value = res
  return 16
}

func JP_NZ_number(this *cpu) int {
  val := this.fetch_16()
  if(!this.get_f_zero()) {
    (*this).pc.value = val
  }
  return 12
}

func JP_Z_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_zero()) {
    (*this).pc.value = val
  }
  return 12
}

func JP_NC_number(this *cpu) int {
  val := this.fetch_16()
  if(!this.get_f_carry()) {
    (*this).pc.value = val
  }
  return 12
}

func JP_C_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_carry()) {
    (*this).pc.value = val
  }
  return 12
}

func JP_HL(this *cpu) int {
  (*this).pc.value = (*this).hl.value
  return 4
}

func STOP(this *cpu) int {
  //(*this).Halt = true
  println("STOP operation received")

  //nächster wert wird ignoriert
  (*this).fetch()
  return 4
}

func CALL_NZ_number(this *cpu) int {
  val := this.fetch_16()
  if(!this.get_f_zero()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
  }
  return 12
}

func CALL_Z_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_zero()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
  }
  return 12
}

func CALL_NC_number(this *cpu) int {
  val := this.fetch_16()
  if(!this.get_f_carry()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
  }
  return 12
}

func CALL_C_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_carry()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
  }
  return 12
}

func PUSH_AF(this *cpu) int {
  this.pushStack((*this).af.value)
  return 16
}

func PUSH_DE(this *cpu) int {
  this.pushStack((*this).de.value)
  return 16
}

func PUSH_HL(this *cpu) int {
  this.pushStack((*this).hl.value)
  return 16
}

func HALT(this *cpu) int {
  (*this).Halt = true
  return 4
}

func DEC_H(this *cpu) int {
  (*this).hl.w_high(this.decrement((*this).hl.r_high()))
  return 4
}

func DEC_L(this *cpu) int {
  (*this).hl.w_low(this.decrement((*this).hl.r_low()))
  return 4
}

func RET_NZ(this *cpu) int {
  if(!this.get_f_zero()) {
    this.pc.value = this.popStack()
  }
  return 8
}

func RET_NC(this *cpu) int {
  if(!this.get_f_carry()) {
    this.pc.value = this.popStack()
  }
  return 8
}

func RET_C(this *cpu) int {
  if(this.get_f_carry()) {
    this.pc.value = this.popStack()
  }
  return 8
}

func LDHL_SP_number(this *cpu) int {
  val := int8(this.fetch())
  sp := uint16((*this).sp.value)
  res := uint16(int32(sp) + int32(val))
  res_temp := sp ^ uint16(val) ^ res
  this.set_f_zero(false)
  this.set_f_subtr(false)
  this.set_f_h_carry((res_temp & 0x10) == 0x10)
  this.set_f_carry((res_temp & 0x100) == 0x100)
  return 12
}

func RLCA(this *cpu) int {
  a := (*this).af.r_high()
  res := uint8(a << 1) | (a >> 7)
  (*this).af.w_high(res)
  this.set_f_zero(false)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(res > 0x7f)
  return 4
}

func DAA(this *cpu) int {
  if(!this.get_f_subtr()) {
    if(this.get_f_carry() || (*this).af.r_high() > 0x99) {
      (*this).af.w_high((*this).af.r_high() + 0x60)
      this.set_f_carry(true)
    }
    if(this.get_f_h_carry() || ((*this).af.r_high() & 0xf) > 0x9) {
      (*this).af.w_high((*this).af.r_high() + 0x06)
      this.set_f_h_carry(false)
    }
  } else if(this.get_f_carry() && this.get_f_h_carry()) {
    (*this).af.w_high((*this).af.r_high() + 0x9a)
    (*this).set_f_h_carry(false)
  } else if(this.get_f_carry()) {
    (*this).af.w_high((*this).af.r_high() + 0xa0)
  } else if(this.get_f_h_carry()) {
    (*this).af.w_high((*this).af.r_high() + 0xfa)
    this.set_f_h_carry(false)
  }
  this.set_f_zero((*this).af.r_high() == 0)
  return 4
}

func SCF(this *cpu) int {
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(true)
  return 4
}

func CPL(this *cpu) int {
  (*this).af.w_high(^(*this).af.r_high())
  this.set_f_subtr(true)
  this.set_f_h_carry(true)
  return 4
}

func (this *cpu) init_ops() [0x100]func(*cpu) int {
  var ops [0x100]func(*cpu) int
  ops[0x00] = NOOP
  ops[0x01] = LD_BC_nn
  ops[0x02] = LD_BC_A
  ops[0x03] = INC_BC
  ops[0x04] = INC_B
  ops[0x05] = DEC_B
  ops[0x06] = LD_B_n
  ops[0x07] = RLCA
  ops[0x08] = LD_number_SP
  ops[0x09] = ADD_HL_BC
  ops[0x0a] = LD_A_BC
  ops[0x0c] = INC_C
  ops[0x0b] = DEC_BC
  ops[0x0d] = DEC_C
  ops[0x0e] = LD_C_n
  ops[0x0f] = RRCA
  ops[0x10] = STOP
  ops[0x11] = LD_DE_nn
  ops[0x12] = LD_DE_A
  ops[0x13] = INC_DE
  ops[0x14] = INC_D
  ops[0x15] = DEC_D
  ops[0x16] = LD_D_n
  ops[0x17] = RLA
  ops[0x18] = JR_n
  ops[0x19] = ADD_HL_DE
  ops[0x1a] = LD_A_DE
  ops[0x1b] = DEC_DE
  ops[0x1c] = INC_E
  ops[0x1d] = DEC_E
  ops[0x1e] = LD_E_n
  ops[0x20] = JR_nZ
  ops[0x21] = LD_HL_nn
  ops[0x22] = LDI_HL_A
  ops[0x23] = INC_HL
  ops[0x24] = INC_H
  ops[0x25] = DEC_H
  ops[0x26] = LD_H_n
  ops[0x27] = DAA
  ops[0x28] = JR_Z
  ops[0x29] = ADD_HL_HL
  ops[0x2a] = LD_A_HLI
  ops[0x2b] = DEC_HL
  ops[0x2c] = INC_L
  ops[0x2d] = DEC_L
  ops[0x2e] = LD_L_n
  ops[0x2f] = CPL
  ops[0x30] = JR_nC
  ops[0x31] = LD_SP_nn
  ops[0x32] = LDD_HL_A
  ops[0x33] = INC_SP
  ops[0x34] = INC_HL_mem
  ops[0x35] = DEC_HLadr
  ops[0x36] = LD_HL_number
  ops[0x37] = SCF
  ops[0x38] = JR_C
  ops[0x39] = ADD_HL_SP
  ops[0x3a] = LDD_A_HL
  ops[0x3b] = DEC_SP
  ops[0x3c] = INC_A
  ops[0x3d] = DEC_A
  ops[0x3e] = LD_A_number
  ops[0x3f] = CCF
  ops[0x40] = LD_B_B
  ops[0x41] = LD_B_C
  ops[0x42] = LD_B_D
  ops[0x43] = LD_B_E
  ops[0x44] = LD_B_H
  ops[0x45] = LD_B_L
  ops[0x46] = LD_B_HL
  ops[0x47] = LD_B_A
  ops[0x48] = LD_C_B
  ops[0x49] = LD_C_C
  ops[0x4a] = LD_C_D
  ops[0x4b] = LD_C_E
  ops[0x4c] = LD_C_H
  ops[0x4d] = LD_C_L
  ops[0x4e] = LD_C_HL
  ops[0x4f] = LD_C_A
  ops[0x50] = LD_D_B
  ops[0x51] = LD_D_C
  ops[0x52] = LD_D_D
  ops[0x53] = LD_D_E
  ops[0x54] = LD_D_H
  ops[0x55] = LD_D_L
  ops[0x56] = LD_D_HL
  ops[0x57] = LD_D_A
  ops[0x58] = LD_E_B
  ops[0x59] = LD_E_C
  ops[0x5a] = LD_E_D
  ops[0x5b] = LD_E_E
  ops[0x5c] = LD_E_H
  ops[0x5d] = LD_E_L
  ops[0x5e] = LD_E_HL
  ops[0x5f] = LD_E_A
  ops[0x60] = LD_H_B
  ops[0x61] = LD_H_C
  ops[0x62] = LD_H_D
  ops[0x63] = LD_H_E
  ops[0x64] = LD_H_H
  ops[0x65] = LD_H_L
  ops[0x66] = LD_H_HL
  ops[0x67] = LD_H_A
  ops[0x68] = LD_L_B
  ops[0x69] = LD_L_C
  ops[0x6a] = LD_L_D
  ops[0x6b] = LD_L_E
  ops[0x6c] = LD_L_H
  ops[0x6d] = LD_L_L
  ops[0x6e] = LD_L_HL
  ops[0x6f] = LD_L_A
  ops[0x70] = LD_HL_B
  ops[0x71] = LD_HL_C
  ops[0x72] = LD_HL_D
  ops[0x73] = LD_HL_E
  ops[0x74] = LD_HL_H
  ops[0x75] = LD_HL_L
  ops[0x76] = HALT
  ops[0x77] = LD_HL_A
  ops[0x78] = LD_A_B
  ops[0x79] = LD_A_C
  ops[0x7a] = LD_A_D
  ops[0x7b] = LD_A_E
  ops[0x7c] = LD_A_H
  ops[0x7d] = LD_A_L
  ops[0x7e] = LD_A_HL
  ops[0x7f] = LD_A_A
  ops[0x80] = ADD_A_B
  ops[0x81] = ADD_A_C
  ops[0x82] = ADD_A_D
  ops[0x83] = ADD_A_E
  ops[0x84] = ADD_A_H
  ops[0x85] = ADD_A_L
  ops[0x86] = ADD_A_HL_memory
  ops[0x87] = ADD_A_A
  ops[0x88] = ADC_A_B
  ops[0x89] = ADC_A_C
  ops[0x8a] = ADC_A_D
  ops[0x8b] = ADC_A_E
  ops[0x8c] = ADC_A_H
  ops[0x8d] = ADC_A_L
  ops[0x8e] = ADC_A_HL
  ops[0x8f] = ADC_A_A
  ops[0x90] = SUB_B
  ops[0x91] = SUB_C
  ops[0x92] = SUB_D
  ops[0x93] = SUB_E
  ops[0x94] = SUB_H
  ops[0x95] = SUB_L
  ops[0x96] = SUB_HL_memory
  ops[0x97] = SUB_A
  ops[0x98] = SBC_A_B
  ops[0x99] = SBC_A_C
  ops[0x9a] = SBC_A_D
  ops[0x9b] = SBC_A_E
  ops[0x9c] = SBC_A_H
  ops[0x9d] = SBC_A_L
  ops[0x9e] = SBC_A_HL
  ops[0x9f] = SBC_A_A
  ops[0xa0] = AND_A_B
  ops[0xa1] = AND_A_C
  ops[0xa2] = AND_A_D
  ops[0xa3] = AND_A_E
  ops[0xa4] = AND_A_H
  ops[0xa5] = AND_A_L
  ops[0xa6] = AND_A_HL
  ops[0xa7] = AND_A_A
  ops[0xa8] = XOR_B
  ops[0xa9] = XOR_C
  ops[0xaa] = XOR_D
  ops[0xab] = XOR_E
  ops[0xac] = XOR_H
  ops[0xad] = XOR_L
  ops[0xae] = XOR_HL
  ops[0xaf] = XOR_A
  ops[0xb0] = OR_A_B
  ops[0xb1] = OR_A_C
  ops[0xb2] = OR_A_D
  ops[0xb3] = OR_A_E
  ops[0xb4] = OR_A_H
  ops[0xb5] = OR_A_L
  ops[0xb6] = OR_A_HL
  ops[0xb7] = OR_A_A
  ops[0xb8] = CP_A_B
  ops[0xb9] = CP_A_C
  ops[0xba] = CP_A_D
  ops[0xbb] = CP_A_E
  ops[0xbc] = CP_A_H
  ops[0xbd] = CP_A_L
  ops[0xbe] = CP_HL_memory
  ops[0xbf] = CP_A_A
  ops[0xc0] = RET_NZ
  ops[0xc1] = POP_BC
  ops[0xc2] = JP_NZ_number
  ops[0xc3] = JP
  ops[0xc4] = CALL_NZ_number
  ops[0xc5] = PUSH_BC
  ops[0xc6] = ADD_A_number
  ops[0xc7] = RST_00
  ops[0xc8] = RET_Z
  ops[0xc9] = RET
  ops[0xca] = JP_Z_number
  ops[0xcb] = CB_OP
  ops[0xcc] = CALL_Z_number
  ops[0xcd] = CALL
  ops[0xce] = ADC_A_number
  ops[0xcf] = RST_08
  ops[0xd0] = RET_NC
  ops[0xd1] = POP_DE
  ops[0xd2] = JP_NC_number
  ops[0xd3] = NOOP
  ops[0xd4] = CALL_C_number
  ops[0xd5] = PUSH_DE
  ops[0xd6] = SUB_number
  ops[0xd7] = RST_10
  ops[0xd8] = RET_C
  ops[0xd9] = RETI
  ops[0xda] = JP_C_number
  ops[0xdb] = NOOP
  ops[0xdc] = CALL_C_number
  ops[0xdd] = NOOP
  ops[0xde] = SBC_A_number
  ops[0xdf] = RST_18
  ops[0xe0] = LD_n_A
  ops[0xe1] = POP_HL
  ops[0xe2] = LD_ff00_C_A
  ops[0xe3] = NOOP
  ops[0xe4] = NOOP
  ops[0xe5] = PUSH_HL
  ops[0xe6] = AND_A_number
  ops[0xe7] = RST_20
  ops[0xe8] = ADD_SP_number
  ops[0xe9] = JP_HL
  ops[0xea] = LD_nn_A
  ops[0xeb] = NOOP
  ops[0xec] = NOOP
  ops[0xed] = NOOP
  ops[0xee] = XOR_number
  ops[0xef] = RST_28
  ops[0xf0] = LD_A_n
  ops[0xf1] = POP_AF
  ops[0xf2] = LD_A_ff00_C
  ops[0xf3] = DI
  ops[0xf4] = NOOP
  ops[0xf5] = PUSH_AF
  ops[0xf6] = OR_A_number
  ops[0xf7] = RST_30
  ops[0xf8] = LDHL_SP_number
  ops[0xf9] = LD_SP_HL
  ops[0xfa] = LD_A_nn
  ops[0xfb] = EI
  ops[0xfc] = NOOP
  ops[0xfd] = NOOP
  ops[0xfe] = CP_n
  ops[0xff] = RST_38
  return ops
}

