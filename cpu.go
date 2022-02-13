package gb

import (
	"fmt";
)

type cpu struct {
	memory *memoryunit
  Interrupt bool
	af Register
	bc Register
	de Register
	hl Register
	sp Register //stack pointer
	pc Register //program counter
  ops [0x100]func(*cpu) int
}

func NewCPU(rom []byte, mu *memoryunit) cpu {
	res := cpu {
		memory: mu,
    Interrupt: false,
		af: Register {value: 0x01B0,},
		bc: Register {value: 0x0013,},
		de: Register {value: 0x00D8,},
		hl: Register {value: 0x014D,},
		sp: Register {value: 0xFFFE,},
		pc: Register {value: 0x0100,},
	}
  res.ops = (&res).init_ops()
	for i := 0; i < len(rom); i++ {
		(*res.memory).Write_8(uint16(i), rom[i])
	}
	return res
}

func (this *cpu) fetch() uint8 {
	op := (*this.memory).Read_8(this.pc.value)
	(*this).pc.value += 1
	return op
}

func (this *cpu) fetch_16() uint16 {
  a := uint16(this.fetch())
  b := uint16(this.fetch())
  return ((b << 8) | a)
}

func (this *cpu) popStack() uint16 {
	x := (*this.memory).Read_16(this.sp.value)
	this.sp.value += 2
	return x
}

func (this *cpu) pushStack(a uint16) {
  (*this).memory.Write_8((*this).sp.value - 1, uint8(uint16(a & 0xFF) >> 8))
  (*this).memory.Write_8((*this).sp.value - 2, uint8(a & 0xFF))
  (*this).sp.value -= 2
}

func (this *cpu) compare_8(a uint8, b uint8) {
	x := a - b
	this.set_f_zero(x == 0)
	this.set_f_subtr(true)
	this.set_f_h_carry((b & 0x0f) > (a & 0x0f))
	this.set_f_carry(a < b)
}

func (this *cpu) decrement(x uint8) uint8 {
  x -= 1
	this.set_f_zero(x == 0)
	this.set_f_subtr(true)
	this.set_f_h_carry((x + 1) & 0x0f == 0)
  return x
}

func (this *cpu) Step() int {
	op := this.fetch()
	fmt.Printf("%02x\n", op)
  f := (*this).ops[op]
  if(f == nil) {
		fmt.Printf("opcode not implemented: %x\n", op)
    fmt.Printf("length of ops: %x\n", len(this.init_ops()))
		return -1
  }
  return f(this)
}
