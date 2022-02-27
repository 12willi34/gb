package gb

import (
  //"fmt"
)

func NOOP(this *cpu) int {
  return 4
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

func LD_DE_A(this *cpu) int {
  a := (*this).af.r_high()
  b := (*this).de.value
  (*(*this).mu).Write_8(b, a)
  return 8
}

func LD_nn_A(this *cpu) int {
  a := (*this).af.r_high()
  b := this.fetch_16()
  (*(*this).mu).Write_8(b, a)
  return 16
}

func LDD_HL_A(this *cpu) int {
  hl := (*this).hl.value
  (*(*this).mu).Write_8(hl, (*this).af.r_high())
  (*this).hl.value = hl - 1
  return 8
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

func CP_HL_memory(this *cpu) int {
  this.compare_8(this.af.r_high(), (*(*this).mu).Read_8((*this).hl.value))
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

func CB_OP(this *cpu) int {
  return ((*this).cb_ops[this.fetch()])(this)
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
  ops[0x0a] = LD_A_BC
  ops[0x0c] = INC_C
  ops[0x0d] = DEC_C
  ops[0x0e] = LD_C_n
  ops[0x11] = LD_DE_nn
  ops[0x12] = LD_DE_A
  ops[0x13] = INC_DE
  ops[0x14] = INC_D
  ops[0x15] = DEC_D
  ops[0x16] = LD_D_n
  ops[0x17] = RLA
  ops[0x18] = JR_n
  ops[0x1a] = LD_A_DE
  ops[0x1c] = INC_E
  ops[0x1d] = DEC_E
  ops[0x1e] = LD_E_n
  ops[0x20] = JR_nZ
  ops[0x21] = LD_HL_nn
  ops[0x22] = LDI_HL_A
  ops[0x23] = INC_HL
  ops[0x24] = INC_H
  ops[0x26] = LD_H_n
  ops[0x28] = JR_Z
  ops[0x2c] = INC_L
  ops[0x2e] = LD_L_n
  ops[0x30] = JR_nC
  ops[0x31] = LD_SP_nn
  ops[0x32] = LDD_HL_A
  ops[0x33] = INC_SP
  ops[0x34] = INC_HL_mem
  ops[0x38] = JR_C
  ops[0x3c] = INC_A
  ops[0x3d] = DEC_A
  ops[0x3e] = LD_A_number
  ops[0x40] = LD_B_B
  ops[0x47] = LD_B_A
  ops[0x4f] = LD_C_A
  ops[0x57] = LD_D_A
  ops[0x5f] = LD_E_A
  ops[0x67] = LD_H_A
  ops[0x6f] = LD_L_A
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
  ops[0xa8] = XOR_B
  ops[0xa9] = XOR_C
  ops[0xaa] = XOR_D
  ops[0xab] = XOR_E
  ops[0xac] = XOR_H
  ops[0xad] = XOR_L
  ops[0xae] = XOR_HL
  ops[0xaf] = XOR_A
  ops[0xbe] = CP_HL_memory
  ops[0xc1] = POP_BC
  ops[0xc3] = JP
  ops[0xc5] = PUSH_BC
  ops[0xc6] = ADD_A_number
  ops[0xc7] = RST_00
  ops[0xc8] = RET_Z
  ops[0xc9] = RET
  ops[0xcb] = CB_OP
  ops[0xcd] = CALL
  ops[0xcf] = RST_08
  ops[0xdf] = RST_18
  ops[0xd6] = SUB_number
  ops[0xd7] = RST_10
  ops[0xe0] = LD_n_A
  ops[0xe2] = LD_ff00_C_A
  ops[0xe7] = RST_20
  ops[0xea] = LD_nn_A
  ops[0xee] = XOR_number
  ops[0xef] = RST_28
  ops[0xf0] = LD_A_n
  ops[0xf2] = LD_A_ff00_C
  ops[0xf3] = DI
  ops[0xf7] = RST_30
  ops[0xf9] = LD_SP_HL
  ops[0xfa] = LD_A_nn
  ops[0xfb] = EI
  ops[0xfe] = CP_n
  ops[0xff] = RST_38
  return ops
}
