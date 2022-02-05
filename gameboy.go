package gb

//pandocs: http://bgb.bircd.org/pandocs.htm
//GBEDG: https://hacktixme.ga/GBEDG/
//golang implementation: https://github.com/Humpheh/goboy
//rust implementation: https://rylev.github.io/DMG-01/public/book/introduction.html
//timing: https://robertovaccari.com/blog/2020_09_26_gameboy/

import (
  "fmt"
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

func (this *GameBoy) Init() {
  fmt.Println("starting gameboy")
  this.loop()
}

func (this *GameBoy) loop() {
  for((*(*this).Processor).Step() != -1) {
    continue
  }
}
