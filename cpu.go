package gb

import ()

type registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	g uint8
	h uint8
	flags uint8
	sp uint16
	pc uint16
}

type cpu struct {}

func NewCPU() cpu {
	return cpu {}
}
