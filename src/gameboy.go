package gb

import (
  "fmt"
  "time"
  "github.com/veandco/go-sdl2/sdl"
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
    last_vblank: 0,
  }
}

func (this *GameBoy) Init() {
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

  texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, width, height)
  if(err != nil) {
    fmt.Println("err creating texture")
    panic(err)
  }
  this.t = texture
  defer texture.Destroy()

  this.boot_loop()
  this.loop()
}

func (this *GameBoy) loop() {
  for this.sdl_loop() {
    this.Interrupter.handle()
    steps := this.Cpu.Step()
    if(steps == -1) { return }
    this.Gpu.Step(steps)
    this.Timer.Timing(steps)
  }
}

func (this *GameBoy) boot_loop() {
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

func (this *GameBoy) sdl_loop() bool {
  if(this.Debug_mode) { return true }
  if(this.Gpu.vblank) {
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
    pixels, _, _ := this.t.Lock(nil)
    i := 0
    for y := int32(0); y < height; y++ {
      for x := int32(0); x < width; x++ {
        a := this.Gpu.buffer[y][x]
        for j := 0; j < 4; j++ {
          pixels[i] = a
          i++
        }
      }
    }
    defer this.t.Unlock()
    this.r.Clear()
    this.r.Copy(this.t, nil, nil)
    this.r.Present()
    took := time.Now().UnixMilli() - this.last_vblank
    if took < vblank_duration {
      sdl.Delay(uint32(vblank_duration - took - 1))
    }
    this.last_vblank = time.Now().UnixMilli()
    this.Gpu.vblank = false
  }
  return true
}
