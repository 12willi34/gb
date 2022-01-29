package gb

import (
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
	instructions map[int]Instruction
}

//ToDo
func NewCPU() cpu {
	res := cpu {
		//ToDo
		af: Register {value: 0,},
		bc: Register {value: 0,},
		de: Register {value: 0,},
		hl: Register {value: 0,},
		sp: Register {value: 0,},
		pc: Register {value: 0,},
		instructions: initInstructionMap(),
	}
	return res
}
