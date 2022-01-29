package gb

import (
)

//instructions: https://gbdev.io/pandocs/CPU_Instruction_Set.html
type Instruction struct {
	name string
	op uint8
	cycles uint8
	f_affected bool
	params []int
}

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
	instructions map[uint8]Instruction
}

//ToDo
func NewCPU() cpu {
}
