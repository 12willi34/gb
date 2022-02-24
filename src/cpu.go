package gb

import (
	"fmt"
  "bufio"
  "os"
)

var reader = bufio.NewReader(os.Stdin)

type cpu struct {
	mu *memoryunit
  Interrupt bool
	af Register
	bc Register
	de Register
	hl Register
	sp Register //stack pointer
	pc Register //program counter
  ops [0x100]func(*cpu) int
  cb_ops [0x100]func(*cpu) int
}

func NewCPU(boot []byte, rom []byte, mu memoryunit) cpu {
	res := cpu {
		mu: &mu,
    Interrupt: false,
		af: Register {value: 0x0000,},
		bc: Register {value: 0x0000,},
		de: Register {value: 0x0000,},
		hl: Register {value: 0x0000,},
		sp: Register {value: 0x0000,},
		pc: Register {value: 0x0000,},
	}
  res.ops = (&res).init_ops()
  res.cb_ops = (&res).init_cb_ops()
	for i := 0; i < len(rom); i++ {
		(*res.mu).Write_8(uint16(i), rom[i])
	}
	for i := 0; i < len(boot); i++ {
		(*res.mu).Write_8(uint16(i), boot[i])
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

func (this *cpu) Step() int {
	op := this.fetch()
  f := (*this).ops[op]
  if(f == nil) {
		fmt.Printf("opcode not implemented: %x\n", op)
		return -1
  }
  return f(this)
}

func (this *cpu) Step_debug() int {
	op := this.fetch()
  f := (*this).ops[op]
  if(f == nil) {
		fmt.Printf("opcode not implemented: %x\n", op)
		return -1
  }
  if(op == 0xcb) {
    fmt.Printf("opcode: cb %02x\n", (*this).mu.Read_8((*this).pc.value))
  } else {
	  fmt.Printf("opcode: %02x\n", op)
  }
  steps := f(this)
	fmt.Printf("master interrupt flag: %t\n", (*this).Interrupt)
	fmt.Printf("af: %04x\n", (*this).af.value)
	fmt.Printf("bc: %04x\n", (*this).bc.value)
	fmt.Printf("de: %04x\n", (*this).de.value)
	fmt.Printf("hl: %04x\n", (*this).hl.value)
	fmt.Printf("sp: %04x\n", (*this).sp.value)
	fmt.Printf("pc: %04x\n\n", (*this).pc.value)
  return steps
}