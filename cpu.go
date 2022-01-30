package gb

//high level guides
//https://raphaelstaebler.medium.com/building-a-gameboy-from-scratch-part-2-the-cpu-d6986a5c6c74

import (
	"fmt"
)

type Register struct {
	value uint16
}

func (reg Register) low() uint8 {
	return uint8(reg.value & 0xFF)
}

func (reg Register) high() uint8 {
	return uint8(reg.value >> 8)
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
func NewCPU() cpu {
	res := cpu {
		af: Register {value: 0x01B0,},
		bc: Register {value: 0x0013,},
		de: Register {value: 0x00D8,},
		hl: Register {value: 0x014D,},
		sp: Register {value: 0xFFFE,},
		pc: Register {value: 0x0100,},
		memory: NewMemoryUnit(),
	}
	return res
}

func (c cpu) state() {
	t := []Register{c.af, c.bc, c.de, c.hl, c.sp, c.pc}
	names := []string{"AF", "BC", "DE", "HL", "SP", "PC"}
	fmt.Println("---cpu state---")
	for i, x := range t {
		if(names[i] == "PC") {
			fmt.Printf("%s: %02x%02x\n", names[i], x.high(), x.low())
		} else {
			fmt.Printf("%s: %02x\t%02x\n", names[i], x.high(), x.low())
		}
	}
}

func (this cpu) fetch() uint8 {
	op := this.memory.Read_8(this.pc.value)
	this.pc.value += 1
	return op
}

func (this cpu) Step() {
	op := this.fetch()
	switch op {
	default:
		fmt.Printf("opcode not implemented: %d\n", op)
	}
	this.state()
}
