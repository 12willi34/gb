package gb

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
	Instructions map[int]Instruction
}

func NewCPU() cpu {
	res := cpu {
		af: Register {value: 0,},
		bc: Register {value: 0,},
		de: Register {value: 0,},
		hl: Register {value: 0,},
		sp: Register {value: 0,},
		pc: Register {value: 0,},
		Instructions: initInstructionMap(),
	}
	return res
}

func (c cpu) Step() {
	c.state()
}

func (c cpu) state() {
	fmt.Println("---cpu state---")
	t := []Register{c.af, c.bc, c.de, c.hl, c.sp, c.pc}
	for _, x := range t {
		fmt.Printf("%x %x\n", x.high(), x.low())
	}
}
