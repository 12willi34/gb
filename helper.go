package gb

import (
  "fmt";
  "math";
)

func (this *cpu) _set_f(x bool, i int) {
  a := this.af.r_low()
  b := uint8(math.Pow(2, float64(i)))
  if(x) {
    this.af.w_low(a | b)
  } else {
    this.af.w_low(a & ^b)
  }
}

func (this *cpu) _get_f(i int) bool {
  l := this.af.r_low() & uint8(math.Pow(2, float64(i)))
  r := uint8(math.Pow(2, float64(i)))
  return l == r
}

func (this *cpu) state() {
  c := *this
  names := []string{"AF", "BC", "DE", "HL", "SP", "PC"}
  fmt.Println("---cpu state---")
  for i, x := range []Register{c.af, c.bc, c.de, c.hl, c.sp, c.pc} {
    if(names[i] == "PC") {
      fmt.Printf("%s: %02x%02x\n", names[i], x.r_high(), x.r_low())
    } else {
      fmt.Printf("%s: %02x\t%02x\n", names[i], x.r_high(), x.r_low())
    }
  }
}
