package gb

import (
  "fmt"
)

const blank_cycles = 69833

type GameBoy struct {
  Mu *memoryunit
  Processor *cpu
}

func NewGameBoy(rom []byte) *GameBoy {
  mu := NewMemoryUnit()
  processor := NewCPU(rom, &mu)
  gameboy := GameBoy {
    Mu: &mu,
    Processor: &processor,
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
    vblank -= steps
    if(vblank <= 0) {
      fmt.Println("first vblank done")
      return
    }
  }
}
