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
  paused bool
}

func NewGameBoy(boot []byte, rom []byte) GameBoy {
  mu := NewMemoryUnit()
  cpu := NewCPU(boot, rom, mu)
  interrupter := NewInterrupter(mu, cpu)
  timer := NewTimer(mu, interrupter)
  gpu := NewGpu(mu, interrupter)
  return GameBoy {
    Mu: &mu,
    Cpu: &cpu,
    Timer: &timer,
    Interrupter: &interrupter,
    Gpu: &gpu,
    paused: false,
  }
}

func (this GameBoy) Init() {
  fmt.Println("starting gameboy")
  this.loop()
}

func (this GameBoy) loop() {
  showWindow()
  for(!this.paused) {
    steps := this.Cpu.Step()
    if(this.Cpu.pc.value == 0x100) {
      fmt.Println("boot finished")
      return
    }
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    this.Interrupter.handle()
  }
}
