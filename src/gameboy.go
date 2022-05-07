package gb

type GameBoy struct {
  Mu *memoryunit
  Cpu *cpu
  Timer *timer
  Interrupter *interrupter
  Gpu *Gpu
  first_rom_part []byte
  Display_callback func() bool
}

func NewGameBoy(boot [256]byte, game string) GameBoy {
  mu, first_rom_part := NewMemoryUnit(boot, game)
  mu_pointer := &mu
  cpu := NewCPU(mu_pointer)
  cpu_pointer := &cpu
  mu.Processor = cpu_pointer
  interrupter := NewInterrupter(mu_pointer, cpu_pointer)
  interrupter_pointer := &interrupter
  timer := NewTimer(mu_pointer, interrupter_pointer)
  gpu := NewGpu(mu_pointer, interrupter_pointer)
  return GameBoy {
    Mu: mu_pointer,
    Cpu: cpu_pointer,
    Timer: &timer,
    Interrupter: interrupter_pointer,
    Gpu: &gpu,
    first_rom_part: first_rom_part,
  }
}

func (this *GameBoy) Boot_loop() {
  for this.Display_callback() {
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    this.Interrupter.handle()
    if(this.Cpu.pc.value >= 0x100) {
      for i := 0; i < 0x100; i++ {
        this.Mu.addr[i] = this.first_rom_part[i]
      }
      return
    }
  }
}

func (this *GameBoy) Loop() {
  for this.Display_callback() {
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    this.Interrupter.handle()
  }
}
