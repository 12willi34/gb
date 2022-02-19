package gb

import (
  "fmt"
  "time"
)

const blank_cycles = 69833

type GameBoy struct {
  Mu *memoryunit
  Processor *cpu
  timer *timer
  interrupter *interrupter
  gpu *Gpu
}

func NewGameBoy(rom []byte) *GameBoy {
  mu := NewMemoryUnit()
  processor := NewCPU(rom, &mu)
  interrupter := Interrupter(mu, processor)
  timer := Timer(mu, interrupter)
  gpu := NewGpu(&mu)
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
    vblank := blank_cycles
    end := time.Now().UnixMilli() + (1/60)*int64(time.Millisecond)
    for(vblank > 0) {
      steps := (*(*this).Processor).Step()
      if(steps == -1) {
        return
      }
      (*(*this).gpu).Step(steps)
      (*(*this).timer).Timing(steps)
      (*(*this).interrupter).handle()
      vblank -= steps
    }
    for(time.Now().UnixMilli() < end) {
      time.Sleep(1*time.Millisecond)
    }
  }
}
