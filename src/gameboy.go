package gb

import (
  "fmt"
  //"time"
)

const blank_cycles = 69833

type GameBoy struct {
  Mu *memoryunit
  Processor *cpu
  timer *timer
  interrupter *interrupter
  gpu *Gpu
}

func NewGameBoy(boot []byte, rom []byte) *GameBoy {
  mu := NewMemoryUnit()
  processor := NewCPU(boot, rom, &mu)
  interrupter := Interrupter(mu, processor)
  timer := Timer(mu, interrupter)
  gpu := NewGpu(&mu, interrupter)
  gameboy := GameBoy {
    Mu: &mu,
    Processor: &processor,
    timer: &timer,
    interrupter: &interrupter,
    gpu: gpu,
  }
  return &gameboy
}

func (this *GameBoy) Init() {
  fmt.Println("starting gameboy")
  this.loop()
}

func (this *GameBoy) loop() {
  for true {
    steps := (*(*this).Processor).Step()
    if(steps == -1) { return }
    (*(*this).gpu).Step(steps)
    (*(*this).timer).Timing(steps)
    (*(*this).interrupter).handle()
  }
}
