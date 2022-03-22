package gb

import (
  "fmt"
  //"time"
  "github.com/veandco/go-sdl2/sdl"
  "image/color"
  //"image"
)

const blank_cycles = 69833
const title = "GameBoy"

type GameBoy struct {
  Mu *memoryunit
  Cpu *cpu
  Timer *timer
  Interrupter *interrupter
  Gpu *Gpu
  rom []byte

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
    rom: rom,
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
        this.Mu.addr[i] = this.rom[i]
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
    } else {
      fmt.Println(err)
    }
    this.w.UpdateSurface()
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
