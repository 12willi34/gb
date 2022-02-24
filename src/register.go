package gb

import (
  "math"
)

type Register struct {
  value uint16
}

func (this *Register) r_low() uint8 {
  return uint8((*this).value & 0xFF)
}

func (this *Register) r_high() uint8 {
  return uint8((*this).value >> 8)
}

func (this *Register) w_low(data uint8) {
  a := uint16((*this).r_high()) << 8
  b := uint16(data)
  (*this).value = a | b
}

func (this *Register) w_high(data uint8) {
  (*this).value = uint16(data) << 8 | uint16((*this).r_low())
}

func (this *cpu) _set_f(set bool, flag int) {
  if(set) {
    this.af.w_low(this.af.r_low() | (1 << flag))
  } else {
    this.af.w_low(this.af.r_low() & ^(1 << flag))
  }
}

func (this *cpu) _get_f(i int) bool {
  l := this.af.r_low() & uint8(math.Pow(2, float64(i)))
  r := uint8(math.Pow(2, float64(i)))
  return l == r
}

func (this *cpu) set_f_zero(x bool) { this._set_f(x, 7) }
func (this *cpu) set_f_subtr(x bool) { this._set_f(x, 6) }
func (this *cpu) set_f_h_carry(x bool) { this._set_f(x, 5) }
func (this *cpu) set_f_carry(x bool) { this._set_f(x, 4) }

func (this *cpu) get_f_zero() bool { return this._get_f(7) }
func (this *cpu) get_f_subtr() bool { return this._get_f(6) }
func (this *cpu) get_f_h_carry() bool { return this._get_f(5) }
func (this *cpu) get_f_carry() bool { return this._get_f(4) }
