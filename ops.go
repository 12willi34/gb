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

func LDD_HL_A(this *cpu) int {
  hl := (*this).hl.value
  (*(*this).memory).Write_8(hl, (*this).af.r_high())
  (*this).hl.value = hl - 1
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
  hl := (*(t.memory)).Read_8(t.hl.value)
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
  return this.init_cb_ops()[this.fetch()](this)
}

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

func (this *cpu) init_ops() [0x100]func(*cpu) int {
  var ops [0x100]func(*cpu) int
  ops[0x00] = NOOP
  ops[0x01] = LD_BC_nn
  ops[0x05] = DEC_B
  ops[0x06] = LD_B_n
  ops[0x0a] = LD_A_BC
  ops[0x11] = LD_DE_nn
  ops[0x1a] = LD_A_DE
  ops[0x20] = JR_nZ
  ops[0x21] = LD_HL_nn
  ops[0x31] = LD_SP_nn
  ops[0x32] = LDD_HL_A
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
  ops[0xa8] = XOR_B
  ops[0xa9] = XOR_C
  ops[0xaa] = XOR_D
  ops[0xab] = XOR_E
  ops[0xac] = XOR_H
  ops[0xad] = XOR_L
  ops[0xae] = XOR_HL
  ops[0xaf] = XOR_A
  ops[0xc1] = POP_BC
  ops[0xc3] = JP
  ops[0xc5] = PUSH_BC
  ops[0xc7] = RST_00
  ops[0xc8] = RET_Z
  ops[0xc9] = RET
  ops[0xcb] = CB_OP
  ops[0xcd] = CALL
  ops[0xcf] = RST_08
  ops[0xdf] = RST_18
  ops[0xd7] = RST_10
  ops[0xe0] = LD_n_A
  ops[0xe7] = RST_20
  ops[0xee] = XOR_number
  ops[0xef] = RST_28
  ops[0xf0] = LD_A_n
  ops[0xf3] = DI
  ops[0xf7] = RST_30
  ops[0xfa] = LD_A_nn
  ops[0xfb] = EI
  ops[0xfe] = CP_n
  ops[0xff] = RST_38
  return ops
}

func (this *cpu) init_cb_ops() [0x100]func(*cpu) int {
  var cb_ops [0x100]func(*cpu) int
  cb_ops[0x30] = SWAP_B
  cb_ops[0x31] = SWAP_C
  cb_ops[0x37] = SWAP_A
  cb_ops[0x32] = SWAP_D
  cb_ops[0x33] = SWAP_E
  cb_ops[0x34] = SWAP_H
  cb_ops[0x35] = SWAP_L
  cb_ops[0x36] = SWAP_HL
  return cb_ops
}

