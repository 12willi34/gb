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
  paused bool

  //sdl
  w *sdl.Window
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
  surface, err := window.GetSurface()

  if(err != nil) {
    fmt.Println("err getting surface")
    panic(err)
  }

  rect := sdl.Rect{0, 0, width, height}
  surface.FillRect(&rect, 0xffffffff)
  window.UpdateSurface()
  this.loop()
}

func (this GameBoy) loop() {
  for(!this.paused && this.sdl_loop()) {
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

func (this GameBoy) sdl_loop() bool {
  if(this.Gpu.vblank) {
    surf, err := this.w.GetSurface()
    if(err == nil) {
      for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
          surf.Set(x, y, color.Alpha {A: this.Gpu.buffer[y][x]*(0xff/4)})
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
