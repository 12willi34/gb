package gb

//pandocs: http://bgb.bircd.org/pandocs.htm
//golang implementation: https://github.com/Humpheh/goboy
//rust implementation: https://rylev.github.io/DMG-01/public/book/introduction.html#why-rust

import (
  "fmt";
)

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

func (gameboy *GameBoy) Init() {
  fmt.Println("starting gameboy")
  for((*(*gameboy).Processor).Step() != -1) {
    continue
  }
}
