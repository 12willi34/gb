package gb

import (
  "fmt"
  //"time"
)

const blank_cycles = 69833

type GameBoy struct {
  Mu *memoryunit
  Cpu *cpu
  Timer *timer
  Interrupter *interrupter
  Gpu *Gpu
}

func NewGameBoy(boot []byte, rom []byte) GameBoy {
  mu := NewMemoryUnit()
  cpu := NewCPU(boot, rom, mu)
  interrupter := NewInterrupter(mu, cpu)
  timer := NewTimer(mu, interrupter)
  gpu := NewGpu(mu, interrupter)
  gameboy := GameBoy {
    Mu: &mu,
    Cpu: &cpu,
    Timer: &timer,
    Interrupter: &interrupter,
    Gpu: &gpu,
  }
  return gameboy
}

func (this GameBoy) Init() {
  fmt.Println("starting gameboy")
  this.loop()
}

func (this GameBoy) loop() {
  for true {
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    this.Interrupter.handle()
  }
}
