package gb

import ()

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
    cb_ops[0x7c] = BIT_7_H
    return cb_ops
}

func BIT_6_B(this *cpu) int {
  this.bit(6, (*this).bc.r_high())
  return 8
}

func BIT_7_H(this *cpu) int {
  this.bit(7, (*this).hl.r_high())
  return 8
}
