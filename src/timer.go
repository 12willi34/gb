package gb

import (
  //"fmt"
)

const div = 0xff04
const tima = 0xff05
const tma = 0xff06
const tac = 0xff07
var frequencies = []int{1024, 16, 64, 256}

type timer struct {
  bus memoryunit
  div_internal uint16
  tima_internal int
  inter interrupter
}

func NewTimer(mu memoryunit, inter interrupter) timer {
  return timer {
    bus: mu,
    div_internal: 0,
    tima_internal: 0,
    inter: inter,
  }
}

func (this *timer) tac_stopped() bool {
  return 1 == 1 & (2 >> (*this).bus.Read_8(tac))
}

func (this *timer) tac_freq() int {
  return frequencies[2 & (*this).bus.Read_8(tac)]
}

func (this *timer) timer_interrupt() {
  (*this).inter.Request(timer_flag)
}

func (this *timer) update_div(cycles int) {
  (*this).div_internal += uint16(cycles)
  for((*this).div_internal > 0xff) {
    (*this).div_internal -= 0xff
    (*this).bus.addr[div] = (*this).bus.Read_8(div) + 1
  }
}

func (this *timer) update_tima(cycles int) {
  (*this).tima_internal += cycles
  for((*this).tima_internal >= this.tac_freq()) {
    (*this).tima_internal -= this.tac_freq()
    t := (*this).bus.Read_8(tima)
    if(t == 0xff) {
      (*this).bus.addr[tima] = (*this).bus.Read_8(tma)
      (*this).timer_interrupt()
    } else {
      (*this).bus.addr[tima] = t + 1
    }
  }
}

func (this *timer) Timing(cycles int) {
  (*this).update_div(cycles)
  if(!this.tac_stopped()) {
    (*this).update_tima(cycles)
  }
}
