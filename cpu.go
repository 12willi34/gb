package gb

//high level guides
//https://raphaelstaebler.medium.com/building-a-gameboy-from-scratch-part-2-the-cpu-d6986a5c6c74

import (
	"fmt";
	"math";
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
	(*this).value = uint16((*this).r_high()<<8)|uint16(data)
}

func (this *Register) w_high(data uint8) {
	(*this).value = (uint16(data)<<8)|uint16((*this).r_low())
}

type cpu struct {
	af Register
	bc Register
	de Register
	hl Register
	sp Register //stack pointer
	pc Register //program counter
	memory memoryunit
}

//initial values from: https://mstojcevich.github.io/post/d-gb-emu-registers/
func NewCPU(rom []byte) cpu {
	res := cpu {
		af: Register {value: 0x01B0,},
		bc: Register {value: 0x0013,},
		de: Register {value: 0x00D8,},
		hl: Register {value: 0x014D,},
		sp: Register {value: 0xFFFE,},
		pc: Register {value: 0x0100,},
		memory: NewMemoryUnit(),
	}
	for i := 0; i < len(rom); i++ {
		res.memory.Write_8(uint16(i), rom[i])
	}
	return res
}

func (this *cpu) set_f(x bool, i int) {
	if x {
		this.af.w_low(this.af.r_low() | uint8(math.Pow(2, float64(i))))
	} else {
		this.af.w_low(this.af.r_low() & ^uint8(math.Pow(2, float64(i))))
	}
}

func (this *cpu) set_f_zero(x bool) { this.set_f(x, 7) }
func (this *cpu) set_f_subtr(x bool) { this.set_f(x, 6) }
func (this *cpu) set_f_h_carry(x bool) { this.set_f(x, 5) }
func (this *cpu) set_f_carry(x bool) { this.set_f(x, 4) }

func (this *cpu) get_f(i int) bool {
	return (this.af.r_low() & uint8(math.Pow(2, float64(i)))) == uint8(math.Pow(2, float64(i)))
}

func (this *cpu) get_f_zero() bool { return this.get_f(7) }
func (this *cpu) get_f_subtr() bool { return this.get_f(6) }
func (this *cpu) get_f_h_carry() bool { return this.get_f(5) }
func (this *cpu) get_f_carry() bool { return this.get_f(4) }

func (this *cpu) state() {
	c := *this
	t := []Register{c.af, c.bc, c.de, c.hl, c.sp, c.pc}
	names := []string{"AF", "BC", "DE", "HL", "SP", "PC"}
	fmt.Println("---cpu state---")
	for i, x := range t {
		if(names[i] == "PC") {
			fmt.Printf("%s: %02x%02x\n", names[i], x.r_high(), x.r_low())
		} else {
			fmt.Printf("%s: %02x\t%02x\n", names[i], x.r_high(), x.r_low())
		}
	}
}

func (this *cpu) popStack() uint16 {
	x := this.memory.Read_16(this.sp.value)
	this.sp.value += 2
	return x
}

func (this *cpu) fetch() uint8 {
	op := (*this).memory.Read_8(this.pc.value)
	(*this).pc.value += 1
	return op
}

func (this *cpu) compare_8(a uint8, b uint8) {
	x := a - b
	this.set_f_zero(x == 0)
	this.set_f_subtr(true)
	this.set_f_h_carry((b & 0x0f) > (a & 0x0f))
	this.set_f_carry(a < b)
}

func (this *cpu) Step() {
	op := this.fetch()
	fmt.Printf("%02x\n", op)
	switch op {
	case 0x0:
		break
	case 0xc3:
		a := this.fetch()
		b := this.fetch()
		this.pc.value = uint16(a | b<<8)
		break
	case 0xc8:
		if(this.get_f_zero()) {
			this.pc.value = this.popStack()
		}
		break
	case 0xe0:
		i := 0xff00 + uint16(this.fetch())
		this.memory.Write_8(i, this.af.r_high())
		break
	case 0xf0:
		a := this.memory.Read_8(0xFF00 + uint16(this.fetch()))
		this.af.w_high(a)
		break
	case 0xfe:
		this.compare_8(this.af.r_high(), this.fetch())
		break
	default:
		fmt.Printf("opcode not implemented: %x\n", op)
	}
	//this.state()
}
