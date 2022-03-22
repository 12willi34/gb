package gb

import (
  "fmt"
  "bufio"
  "os"
)

type cpu struct {
  mu *memoryunit
  Interrupt bool
  Halt bool
  af Register
  bc Register
  de Register
  hl Register
  sp Register //stack pointer
  pc Register //program counter
}

func NewCPU(mu memoryunit) cpu {
  res := cpu {
    mu: &mu,
    Interrupt: false,
    Halt: false,
    af: Register {value: 0x0000,},
    bc: Register {value: 0x0000,},
    de: Register {value: 0x0000,},
    hl: Register {value: 0x0000,},
    sp: Register {value: 0x0000,},
    pc: Register {value: 0x0000,},
  }
  return res
}

func (this *cpu) fetch() uint8 {
  op := (*this.mu).Read_8(this.pc.value)
  (*this).pc.value += 1
  return op
}

func (this *cpu) fetch_16() uint16 {
  a := uint16(this.fetch())
  b := uint16(this.fetch())
  return ((b << 8) | a)
}

func (this *cpu) popStack() uint16 {
  x := (*this.mu).Read_16(this.sp.value)
  this.sp.value += 2
  return x
}

func (this *cpu) pushStack(a uint16) {
  (*this).mu.Write_8((*this).sp.value - 1, uint8(uint16(a & 0xFF00) >> 8))
  (*this).mu.Write_8((*this).sp.value - 2, uint8(a & 0xFF))
  (*this).sp.value -= 2
}

func (this *cpu) compare_8(a uint8, b uint8) {
  x := a - b
  this.set_f_zero(x == 0)
  this.set_f_subtr(true)
  this.set_f_h_carry((b & 0x0f) > (a & 0x0f))
  this.set_f_carry(a < b)
}

func (this *cpu) increment(x uint8) uint8 {
  res := x + 1
  this.set_f_zero(res == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry((x & 0xf) + (1 & 0xf) > 0xf)
  return res
}

//DEC
func (this *cpu) decrement(x uint8) uint8 {
  res := x - 1
  this.set_f_zero(res == 0)
  this.set_f_subtr(true)
  this.set_f_h_carry(x & 0x0f == 0)
  return res
}

func (this *cpu) xor(a uint8, b uint8) uint8 {
  x := a ^ b
  this.set_f_zero(x == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(false)
  return x
}

func (this *cpu) swap(x uint8) uint8 {
  x = uint8(x << 4) | uint8(x >> 4)
  this.set_f_zero(x == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(false)
  return x
}

func (this *cpu) bit(i uint8, val uint8) {
  this.set_f_zero(val & (1 << i) == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(true)
}

//=ADC
func (this *cpu) adc(a uint8, b uint8) uint8 {
  carr := int16(0)
  if(this.get_f_carry()) { carr = 1 }
  res_temp := int16(a) + int16(b) + carr
  res := uint8(res_temp)
  this.set_f_zero(res == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry((b & 0xf) + (a & 0xf) + uint8(carr) > 0xf)
  this.set_f_carry(res_temp > 0xff)
  return res
}

//=ADD
func (this *cpu) add(a uint8, b uint8) uint8 {
  res_temp := int16(a) + int16(b)
  res := uint8(res_temp)
  this.set_f_zero(res == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry((b & 0xf) + (a & 0xf) > 0xf)
  this.set_f_carry(res_temp > 0xff)
  return res
}

//=ADD 16
func (this *cpu) add_16(a uint16 , b uint16) uint16 {
  res_temp := int32(a) + int32(b)
  this.set_f_subtr(false)
  this.set_f_h_carry(int32(a & 0xfff) > (res_temp & 0xfff))
  this.set_f_carry(res_temp > 0xffff)
  return uint16(res_temp)
}

//=RST
func (this *cpu) restart(next uint16) {
  this.pushStack((*this).pc.value)
  (*this).pc.value = next
}

//=SUB
func (this *cpu) subtract(a uint8, b uint8) uint8 {
  res_temp := int16(a) - int16(b)
  res := uint8(res_temp)
  this.set_f_zero(res == 0)
  this.set_f_subtr(true)
  this.set_f_h_carry((int16(a & 0x0f) - int16(b & 0xf)) < 0)
  this.set_f_carry(res_temp < 0)
  return res
}

//=SBC
func (this *cpu) subtract_carry(a uint8, b uint8) uint8 {
  carry := uint16(0)
  if((*this).get_f_carry()) {carry = uint16(1)}
  res := uint16(a) - (uint16(b) + carry)
  this.set_f_zero(res == 0)
  this.set_f_subtr(true)
  hc := int16(a & 0x0f) - (int16(b & 0xf) + int16(carry))
  this.set_f_h_carry(hc < 0)
  this.set_f_carry(a < b)
  return uint8(res)
}

//=RL
func (this *cpu) rotate_left(a uint8) uint8 {
  old_cy := uint8(0)
  if(this.get_f_carry()) { old_cy = uint8(1) }
  new_cy := a >> 7
  res := uint8(((a << 1) & 0xff) | old_cy)
  this.set_f_zero(res == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(new_cy == 1)
  return res
}

//=RLA
func (this *cpu) rla(a uint8) uint8 {
  old_cy := uint8(0)
  if(this.get_f_carry()) { old_cy = uint8(1) }
  new_cy := a >> 7
  res := uint8(((a << 1) & 0xff) | old_cy)
  this.set_f_zero(false)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(new_cy == 1)
  return res
}

//=OR
func (this *cpu) or(a uint8, b uint8) uint8 {
  a |= b
  this.set_f_zero(a == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(false)
  return a
}

//=AND
func (this *cpu) and(a uint8, b uint8) uint8 {
  res := a & b
  this.set_f_zero(res == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(true)
  this.set_f_carry(false)
  return res
}

//=SRL
func (this *cpu) srl(a uint8) uint8 {
  carr := a & uint8(1)
  res := a >> 1
  this.set_f_zero(res == 0)
  this.set_f_subtr(false)
  this.set_f_h_carry(false)
  this.set_f_carry(carr == 1)
  return res
}

func (this *cpu) Step() int {
  if((*this).Halt) { return 4 }
  var cycles = -1
  op := this.fetch()
  if(op == 0xcb) {
    cb_op := this.fetch()
    cycles = this.do_cb_op(cb_op)
  } else {
    cycles = this.do_op(op)
  }
  if(cycles == -1) {
    fmt.Printf("opcode not implemented: %x\n", op)
    return -1
  }
  return cycles
}

var reader = bufio.NewReader(os.Stdin)

func (this *cpu) Step_debug() int {
  if((*this).Halt) { return 4 }
  var cycles = -1
  op := this.fetch()
  if(op == 0xcb) {
    cb_op := this.fetch()
    cycles = this.do_cb_op(cb_op)
    fmt.Printf("opcode: cb %02x\n", cb_op)
  } else {
    cycles = this.do_op(op)
    fmt.Printf("opcode: %02x\n", op)
  }
  if(cycles == -1) {
    fmt.Printf("opcode not implemented")
    return -1
  }
  fmt.Printf("af: %04x\n", (*this).af.value)
  fmt.Printf("bc: %04x\n", (*this).bc.value)
  fmt.Printf("de: %04x\n", (*this).de.value)
  fmt.Printf("hl: %04x\n", (*this).hl.value)
  fmt.Printf("sp: %04x\n", (*this).sp.value)
  fmt.Printf("pc: %04x\n\n", (*this).pc.value)
  reader.ReadString('\n')
  return cycles
}
