package gb

import (
  "fmt"
  "time"
  "github.com/veandco/go-sdl2/sdl"
  "image/color"
  //"image"
)

const blank_cycles = 69833
const title = "GameBoy"
const vblank_duration = int64(1000/60)

type GameBoy struct {
  Mu *memoryunit
  Cpu *cpu
  Timer *timer
  Interrupter *interrupter
  Gpu *Gpu
  first_rom_part []byte
  last_vblank int64

  //sdl
  w *sdl.Window
}

func NewGameBoy(boot []byte, rom []byte) GameBoy {
  mu := NewMemoryUnit(boot, rom)
  cpu := NewCPU(mu)
  mu.Processor = cpu
  interrupter := NewInterrupter(mu, cpu)
  timer := NewTimer(mu, interrupter)
  gpu := NewGpu(mu, interrupter)
  return GameBoy {
    Mu: &mu,
    Cpu: &cpu,
    Timer: &timer,
    Interrupter: &interrupter,
    Gpu: &gpu,
    first_rom_part: rom[:0x100],
    last_vblank: time.Now().Add(-1*time.Hour).UnixMilli(),
    w: nil,
  }
}

func (this GameBoy) Init() {
  fmt.Println("starting gameboy")
  pos := int32(sdl.WINDOWPOS_CENTERED)
  window, err := sdl.CreateWindow(title, pos, pos, width, height, sdl.WINDOW_SHOWN)
  this.w = window
  if(err != nil) {
    fmt.Println("err creating window")
    panic(err)
  }
  defer window.Destroy()
  this.boot_loop()
  println("boot done")
  this.loop()
}

func (this GameBoy) loop() {
  for this.sdl_loop() {
    this.Interrupter.handle()
    steps := this.Cpu.Step_debug()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
  }
}

func (this GameBoy) boot_loop() {
  for this.sdl_loop() {
    steps := this.Cpu.Step()
    if(this.Cpu.pc.value >= 0x100) {
      for i := 0; i < 0x100; i++ {
        this.Mu.addr[i] = this.first_rom_part[i]
      }
      return
    }
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    this.Interrupter.handle()
  }
}

func (this GameBoy) sdl_loop() bool {
  if(this.Gpu.vblank) {
    surf, err := this.w.GetSurface()
    if(err == nil) {
      for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
          surf.Set(x, y, color.Alpha {A: this.Gpu.buffer[y][x]})
        }
      }
      for time.Now().UnixMilli() - this.last_vblank < vblank_duration {
        time.Sleep(10*time.Millisecond)
      }
    } else {
      fmt.Println(err)
    }
    this.w.UpdateSurface()
    this.last_vblank = time.Now().UnixMilli()
    this.Gpu.vblank = false
  }
  for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
    switch event.(type) {
    case *sdl.QuitEvent:
      return false
    }
  }
  return true
}
