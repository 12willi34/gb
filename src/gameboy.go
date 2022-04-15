package gb

import (
  "fmt"
  "time"
  "github.com/veandco/go-sdl2/sdl"
  //"image/color"
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
  r *sdl.Renderer
  t *sdl.Texture
  Debug_mode bool
}

func NewGameBoy(boot [256]byte, rom []byte) GameBoy {
  mu := NewMemoryUnit(boot, rom)
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
    first_rom_part: rom[:0x100],
    last_vblank: time.Now().Add(-1*time.Hour).UnixMilli(),
  }
}

func (this GameBoy) Init() {
  fmt.Println("starting gameboy")
  pos := int32(sdl.WINDOWPOS_CENTERED)
  window, err := sdl.CreateWindow(title, pos, pos, width, height, sdl.WINDOW_SHOWN)
  if(err != nil) {
    fmt.Println("err creating window")
    panic(err)
  }
  this.w = window
  defer window.Destroy()

  renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
  if(err != nil) {
    fmt.Println("err creating renderer")
    panic(err)
  }
  this.r = renderer
  defer renderer.Destroy()

  texture, err := renderer.CreateTexture(0, -1, width, height)
  if(err != nil) {
    fmt.Println("err creating texture")
    panic(err)
  }
  this.t = texture
  defer texture.Destroy()

  timeBeforeBoot := time.Now().UnixMilli()
  this.boot_loop()
  fmt.Printf("boot done (%dms)\n", time.Now().UnixMilli() - timeBeforeBoot)
  this.loop()
}

func (this GameBoy) loop() {
  for this.sdl_loop() {
    this.Interrupter.handle()
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
  }
}

func (this GameBoy) boot_loop() {
  for this.sdl_loop() {
    this.Interrupter.handle()
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
    if(this.Cpu.pc.value >= 0x100) {
      for i := 0; i < 0x100; i++ {
        this.Mu.addr[i] = this.first_rom_part[i]
      }
      return
    }
  }
}

func (this GameBoy) sdl_loop() bool {
  if(this.Debug_mode) { return true }
  for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
    switch t := event.(type) {
    case *sdl.KeyboardEvent:
      keyCode := t.Keysym.Sym
      if(t.State == sdl.PRESSED) {
        if(t.Repeat > 0) { break }
        if(this.Mu.SetBtn(int(keyCode), true)) {
          this.Interrupter.Request(4)
        }
      } else if(t.State == sdl.RELEASED) {
        if(this.Mu.SetBtn(int(keyCode), false)) {
          this.Interrupter.Request(4)
        }
      }
    case *sdl.QuitEvent:
      return false
    }
  }
  if(this.Gpu.vblank) {
    if(true) {
      for x := int32(0); x < width; x++ {
        for y := int32(0); y < height; y++ {
          a := this.Gpu.buffer[y][x]
          this.r.SetDrawColor(a, a, a, 255)
          this.r.DrawPoint(x, y)
        }
      }
      this.r.Present()
    } else {
      fmt.Println("")
    }
    this.w.UpdateSurface()
    this.last_vblank = time.Now().UnixMilli()
    this.Gpu.vblank = false
  }
  return true
}
