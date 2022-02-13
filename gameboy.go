package gb

import (
  "fmt"
)

const blank_cycles = 69833

type GameBoy struct {
  Mu *memoryunit
  Processor *cpu
  timer *timer
}

func NewGameBoy(rom []byte) *GameBoy {
  mu := NewMemoryUnit()
  processor := NewCPU(rom, &mu)
  timer := Timer(mu)
  gameboy := GameBoy {
    Mu: &mu,
    Processor: &processor,
    timer: timer,
  }
  return &gameboy
}

func (this *GameBoy) Init() {
  fmt.Println("starting gameboy")
  this.loop()
}

func (this *GameBoy) loop() {
  vblank := blank_cycles
  for true {
    steps := (*(*this).Processor).Step()
    if(steps == -1) {
      return
    }
    this.timer.Timing(steps)
    vblank -= steps
    if(vblank <= 0) {
      fmt.Println("first vblank done")
      return
    }
  }
}
