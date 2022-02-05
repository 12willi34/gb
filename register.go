package gb

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
  a := uint16(data) << 8
  b := uint16((*this).r_low())
  (*this).value = a | b
}
