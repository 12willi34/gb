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
  Debug_mode bool
}

func NewGameBoy(boot [0x100]byte, rom []byte) GameBoy {
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
    switch t := event.(type) {
    case *sdl.KeyboardEvent:
      keyCode := t.Keysym.Sym
      if(t.State == sdl.PRESSED) {
        if(t.Repeat > 0) { break }
        this.Mu.SetBtn(int(keyCode), true)
      } else if(t.State == sdl.RELEASED) {
        this.Mu.SetBtn(int(keyCode), false)
      }
    case *sdl.QuitEvent:
      return false
    }
  }
  return true
}
