package gb

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
    return 12
  }
  return 8
}

func JR_Z(this *cpu) int {
  a := int8(this.fetch())
  if(this.get_f_zero()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
    return 12
  }
  return 8
}

func JR_nC(this *cpu) int {
  a := int8(this.fetch())
  if(!this.get_f_carry()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
    return 12
  }
  return 8
}

func JR_C(this *cpu) int {
  a := int8(this.fetch())
  if(this.get_f_carry()) {
    (*this).pc.value = uint16(int32((*this).pc.value) + int32(a))
    return 12
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
    return 20
  }
  return 8
}

func RET(this *cpu) int {
  (*this).pc.value = this.popStack()
  return 16
}

func CALL(this *cpu) int {
  (*this).pushStack((*this).pc.value)
  (*this).pc.value = this.fetch_16()
  return 24
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
  return 16
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
    return 16
  }
  return 12
}

func JP_Z_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_zero()) {
    (*this).pc.value = val
    return 16
  }
  return 12
}

func JP_NC_number(this *cpu) int {
  val := this.fetch_16()
  if(!this.get_f_carry()) {
    (*this).pc.value = val
    return 16
  }
  return 12
}

func JP_C_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_carry()) {
    (*this).pc.value = val
    return 16
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
    return 24
  }
  return 12
}

func CALL_Z_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_zero()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
    return 24
  }
  return 12
}

func CALL_NC_number(this *cpu) int {
  val := this.fetch_16()
  if(!this.get_f_carry()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
    return 24
  }
  return 12
}

func CALL_C_number(this *cpu) int {
  val := this.fetch_16()
  if(this.get_f_carry()) {
    (*this).pushStack((*this).pc.value)
    (*this).pc.value = val
    return 24
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
    return 20
  }
  return 8
}

func RET_NC(this *cpu) int {
  if(!this.get_f_carry()) {
    this.pc.value = this.popStack()
    return 20
  }
  return 8
}

func RET_C(this *cpu) int {
  if(this.get_f_carry()) {
    this.pc.value = this.popStack()
    return 20
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

func (this *cpu) do_op(op uint8) int {
  switch(op) {
  case 0x00:
    return NOOP(this)
  case 0x01:
    return LD_BC_nn(this)
  case 0x02:
    return LD_BC_A(this)
  case 0x03:
    return INC_BC(this)
  case 0x04:
    return INC_B(this)
  case 0x05:
    return DEC_B(this)
  case 0x06:
    return LD_B_n(this)
  case 0x07:
    return RLCA(this)
  case 0x08:
    return LD_number_SP(this)
  case 0x09:
    return ADD_HL_BC(this)
  case 0x0a:
    return LD_A_BC(this)
  case 0x0c:
    return INC_C(this)
  case 0x0b:
    return DEC_BC(this)
  case 0x0d:
    return DEC_C(this)
  case 0x0e:
    return LD_C_n(this)
  case 0x0f:
    return RRCA(this)
  case 0x10:
    return STOP(this)
  case 0x11:
    return LD_DE_nn(this)
  case 0x12:
    return LD_DE_A(this)
  case 0x13:
    return INC_DE(this)
  case 0x14:
    return INC_D(this)
  case 0x15:
    return DEC_D(this)
  case 0x16:
    return LD_D_n(this)
  case 0x17:
    return RLA(this)
  case 0x18:
    return JR_n(this)
  case 0x19:
    return ADD_HL_DE(this)
  case 0x1a:
    return LD_A_DE(this)
  case 0x1b:
    return DEC_DE(this)
  case 0x1c:
    return INC_E(this)
  case 0x1d:
    return DEC_E(this)
  case 0x1e:
    return LD_E_n(this)
  case 0x20:
    return JR_nZ(this)
  case 0x21:
    return LD_HL_nn(this)
  case 0x22:
    return LDI_HL_A(this)
  case 0x23:
    return INC_HL(this)
  case 0x24:
    return INC_H(this)
  case 0x25:
    return DEC_H(this)
  case 0x26:
    return LD_H_n(this)
  case 0x27:
    return DAA(this)
  case 0x28:
    return JR_Z(this)
  case 0x29:
    return ADD_HL_HL(this)
  case 0x2a:
    return LD_A_HLI(this)
  case 0x2b:
    return DEC_HL(this)
  case 0x2c:
    return INC_L(this)
  case 0x2d:
    return DEC_L(this)
  case 0x2e:
    return LD_L_n(this)
  case 0x2f:
    return CPL(this)
  case 0x30:
    return JR_nC(this)
  case 0x31:
    return LD_SP_nn(this)
  case 0x32:
    return LDD_HL_A(this)
  case 0x33:
    return INC_SP(this)
  case 0x34:
    return INC_HL_mem(this)
  case 0x35:
    return DEC_HLadr(this)
  case 0x36:
    return LD_HL_number(this)
  case 0x37:
    return SCF(this)
  case 0x38:
    return JR_C(this)
  case 0x39:
    return ADD_HL_SP(this)
  case 0x3a:
    return LDD_A_HL(this)
  case 0x3b:
    return DEC_SP(this)
  case 0x3c:
    return INC_A(this)
  case 0x3d:
    return DEC_A(this)
  case 0x3e:
    return LD_A_number(this)
  case 0x3f:
    return CCF(this)
  case 0x40:
    return LD_B_B(this)
  case 0x41:
    return LD_B_C(this)
  case 0x42:
    return LD_B_D(this)
  case 0x43:
    return LD_B_E(this)
  case 0x44:
    return LD_B_H(this)
  case 0x45:
    return LD_B_L(this)
  case 0x46:
    return LD_B_HL(this)
  case 0x47:
    return LD_B_A(this)
  case 0x48:
    return LD_C_B(this)
  case 0x49:
    return LD_C_C(this)
  case 0x4a:
    return LD_C_D(this)
  case 0x4b:
    return LD_C_E(this)
  case 0x4c:
    return LD_C_H(this)
  case 0x4d:
    return LD_C_L(this)
  case 0x4e:
    return LD_C_HL(this)
  case 0x4f:
    return LD_C_A(this)
  case 0x50:
    return LD_D_B(this)
  case 0x51:
    return LD_D_C(this)
  case 0x52:
    return LD_D_D(this)
  case 0x53:
    return LD_D_E(this)
  case 0x54:
    return LD_D_H(this)
  case 0x55:
    return LD_D_L(this)
  case 0x56:
    return LD_D_HL(this)
  case 0x57:
    return LD_D_A(this)
  case 0x58:
    return LD_E_B(this)
  case 0x59:
    return LD_E_C(this)
  case 0x5a:
    return LD_E_D(this)
  case 0x5b:
    return LD_E_E(this)
  case 0x5c:
    return LD_E_H(this)
  case 0x5d:
    return LD_E_L(this)
  case 0x5e:
    return LD_E_HL(this)
  case 0x5f:
    return LD_E_A(this)
  case 0x60:
    return LD_H_B(this)
  case 0x61:
    return LD_H_C(this)
  case 0x62:
    return LD_H_D(this)
  case 0x63:
    return LD_H_E(this)
  case 0x64:
    return LD_H_H(this)
  case 0x65:
    return LD_H_L(this)
  case 0x66:
    return LD_H_HL(this)
  case 0x67:
    return LD_H_A(this)
  case 0x68:
    return LD_L_B(this)
  case 0x69:
    return LD_L_C(this)
  case 0x6a:
    return LD_L_D(this)
  case 0x6b:
    return LD_L_E(this)
  case 0x6c:
    return LD_L_H(this)
  case 0x6d:
    return LD_L_L(this)
  case 0x6e:
    return LD_L_HL(this)
  case 0x6f:
    return LD_L_A(this)
  case 0x70:
    return LD_HL_B(this)
  case 0x71:
    return LD_HL_C(this)
  case 0x72:
    return LD_HL_D(this)
  case 0x73:
    return LD_HL_E(this)
  case 0x74:
    return LD_HL_H(this)
  case 0x75:
    return LD_HL_L(this)
  case 0x76:
    return HALT(this)
  case 0x77:
    return LD_HL_A(this)
  case 0x78:
    return LD_A_B(this)
  case 0x79:
    return LD_A_C(this)
  case 0x7a:
    return LD_A_D(this)
  case 0x7b:
    return LD_A_E(this)
  case 0x7c:
    return LD_A_H(this)
  case 0x7d:
    return LD_A_L(this)
  case 0x7e:
    return LD_A_HL(this)
  case 0x7f:
    return LD_A_A(this)
  case 0x80:
    return ADD_A_B(this)
  case 0x81:
    return ADD_A_C(this)
  case 0x82:
    return ADD_A_D(this)
  case 0x83:
    return ADD_A_E(this)
  case 0x84:
    return ADD_A_H(this)
  case 0x85:
    return ADD_A_L(this)
  case 0x86:
    return ADD_A_HL_memory(this)
  case 0x87:
    return ADD_A_A(this)
  case 0x88:
    return ADC_A_B(this)
  case 0x89:
    return ADC_A_C(this)
  case 0x8a:
    return ADC_A_D(this)
  case 0x8b:
    return ADC_A_E(this)
  case 0x8c:
    return ADC_A_H(this)
  case 0x8d:
    return ADC_A_L(this)
  case 0x8e:
    return ADC_A_HL(this)
  case 0x8f:
    return ADC_A_A(this)
  case 0x90:
    return SUB_B(this)
  case 0x91:
    return SUB_C(this)
  case 0x92:
    return SUB_D(this)
  case 0x93:
    return SUB_E(this)
  case 0x94:
    return SUB_H(this)
  case 0x95:
    return SUB_L(this)
  case 0x96:
    return SUB_HL_memory(this)
  case 0x97:
    return SUB_A(this)
  case 0x98:
    return SBC_A_B(this)
  case 0x99:
    return SBC_A_C(this)
  case 0x9a:
    return SBC_A_D(this)
  case 0x9b:
    return SBC_A_E(this)
  case 0x9c:
    return SBC_A_H(this)
  case 0x9d:
    return SBC_A_L(this)
  case 0x9e:
    return SBC_A_HL(this)
  case 0x9f:
    return SBC_A_A(this)
  case 0xa0:
    return AND_A_B(this)
  case 0xa1:
    return AND_A_C(this)
  case 0xa2:
    return AND_A_D(this)
  case 0xa3:
    return AND_A_E(this)
  case 0xa4:
    return AND_A_H(this)
  case 0xa5:
    return AND_A_L(this)
  case 0xa6:
    return AND_A_HL(this)
  case 0xa7:
    return AND_A_A(this)
  case 0xa8:
    return XOR_B(this)
  case 0xa9:
    return XOR_C(this)
  case 0xaa:
    return XOR_D(this)
  case 0xab:
    return XOR_E(this)
  case 0xac:
    return XOR_H(this)
  case 0xad:
    return XOR_L(this)
  case 0xae:
    return XOR_HL(this)
  case 0xaf:
    return XOR_A(this)
  case 0xb0:
    return OR_A_B(this)
  case 0xb1:
    return OR_A_C(this)
  case 0xb2:
    return OR_A_D(this)
  case 0xb3:
    return OR_A_E(this)
  case 0xb4:
    return OR_A_H(this)
  case 0xb5:
    return OR_A_L(this)
  case 0xb6:
    return OR_A_HL(this)
  case 0xb7:
    return OR_A_A(this)
  case 0xb8:
    return CP_A_B(this)
  case 0xb9:
    return CP_A_C(this)
  case 0xba:
    return CP_A_D(this)
  case 0xbb:
    return CP_A_E(this)
  case 0xbc:
    return CP_A_H(this)
  case 0xbd:
    return CP_A_L(this)
  case 0xbe:
    return CP_HL_memory(this)
  case 0xbf:
    return CP_A_A(this)
  case 0xc0:
    return RET_NZ(this)
  case 0xc1:
    return POP_BC(this)
  case 0xc2:
    return JP_NZ_number(this)
  case 0xc3:
    return JP(this)
  case 0xc4:
    return CALL_NZ_number(this)
  case 0xc5:
    return PUSH_BC(this)
  case 0xc6:
    return ADD_A_number(this)
  case 0xc7:
    return RST_00(this)
  case 0xc8:
    return RET_Z(this)
  case 0xc9:
    return RET(this)
  case 0xca:
    return JP_Z_number(this)
  case 0xcb:
    return NOOP(this)
  case 0xcc:
    return CALL_Z_number(this)
  case 0xcd:
    return CALL(this)
  case 0xce:
    return ADC_A_number(this)
  case 0xcf:
    return RST_08(this)
  case 0xd0:
    return RET_NC(this)
  case 0xd1:
    return POP_DE(this)
  case 0xd2:
    return JP_NC_number(this)
  case 0xd3:
    return NOOP(this)
  case 0xd4:
    return CALL_C_number(this)
  case 0xd5:
    return PUSH_DE(this)
  case 0xd6:
    return SUB_number(this)
  case 0xd7:
    return RST_10(this)
  case 0xd8:
    return RET_C(this)
  case 0xd9:
    return RETI(this)
  case 0xda:
    return JP_C_number(this)
  case 0xdb:
    return NOOP(this)
  case 0xdc:
    return CALL_C_number(this)
  case 0xdd:
    return NOOP(this)
  case 0xde:
    return SBC_A_number(this)
  case 0xdf:
    return RST_18(this)
  case 0xe0:
    return LD_n_A(this)
  case 0xe1:
    return POP_HL(this)
  case 0xe2:
    return LD_ff00_C_A(this)
  case 0xe3:
    return NOOP(this)
  case 0xe4:
    return NOOP(this)
  case 0xe5:
    return PUSH_HL(this)
  case 0xe6:
    return AND_A_number(this)
  case 0xe7:
    return RST_20(this)
  case 0xe8:
    return ADD_SP_number(this)
  case 0xe9:
    return JP_HL(this)
  case 0xea:
    return LD_nn_A(this)
  case 0xeb:
    return NOOP(this)
  case 0xec:
    return NOOP(this)
  case 0xed:
    return NOOP(this)
  case 0xee:
    return XOR_number(this)
  case 0xef:
    return RST_28(this)
  case 0xf0:
    return LD_A_n(this)
  case 0xf1:
    return POP_AF(this)
  case 0xf2:
    return LD_A_ff00_C(this)
  case 0xf3:
    return DI(this)
  case 0xf4:
    return NOOP(this)
  case 0xf5:
    return PUSH_AF(this)
  case 0xf6:
    return OR_A_number(this)
  case 0xf7:
    return RST_30(this)
  case 0xf8:
    return LDHL_SP_number(this)
  case 0xf9:
    return LD_SP_HL(this)
  case 0xfa:
    return LD_A_nn(this)
  case 0xfb:
    return EI(this)
  case 0xfc:
    return NOOP(this)
  case 0xfd:
    return NOOP(this)
  case 0xfe:
    return CP_n(this)
  case 0xff:
    return RST_38(this)
  default:
    return -1
  }
}

