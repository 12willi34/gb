package gb

//timer: https://hacktixme.ga/GBEDG/timers

import ()

const div = 0xff04 //divider register
const tima = 0xff05 //timer counter
const tma = 0xff06 //timer modulo
const tac = 0xff07 //timer control

type Timer struct {
  mu *memoryunit
}

func NewTimer(mu *memoryunit) *Timer {
  return Timer {
    mu: memoryunit,
  }
}

func (this *Timer) increment() {}
