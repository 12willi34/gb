package main

import "github.com/veandco/go-sdl2/sdl"
import "gb"
import "time"
import "os"
import "fmt"

type SdlApplication struct {
  gameboy *gb.GameBoy
  w *sdl.Window
  r *sdl.Renderer
  t *sdl.Texture
  last_vblank int64
}

const vblank_duration = int64(1000/60)
const height = 144
const width = 160

func main() {
  gameboy := gb.NewGameBoy(gb.Boot, os.Args[1])
  pos := int32(sdl.WINDOWPOS_CENTERED)

  window, err := sdl.CreateWindow("SimpleGB", pos, pos, width, height, sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)
  if(err != nil) {
    fmt.Println("err creating window")
    panic(err)
  }
  defer window.Destroy()

  renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
  if(err != nil) {
    fmt.Println("err creating renderer")
    panic(err)
  }
  renderer.RenderWindowToLogical(width, height)
  defer renderer.Destroy()

  texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, width, height)
  if(err != nil) {
    fmt.Println("err creating texture")
    panic(err)
  }
  defer texture.Destroy()

  sdlApp := SdlApplication {
    gameboy: &gameboy,
    w: window,
    r: renderer,
    t: texture,
    last_vblank: 0,
  }
  x := &sdlApp
  x.gameboy.Display_callback = x.loop
  x.gameboy.Boot_loop()
  x.gameboy.Loop()
}

func (this *SdlApplication) loop() bool {
  if(this.gameboy.Gpu.Vblank) {
    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
      switch t := event.(type) {
      case *sdl.KeyboardEvent:
        keyCode := t.Keysym.Sym
        if(t.State == sdl.PRESSED) {
          if(t.Repeat > 0) { break }
          if(this.gameboy.Mu.SetBtn(int(keyCode), true)) {
            this.gameboy.Interrupter.Request(4)
          }
        } else if(t.State == sdl.RELEASED) {
          if(this.gameboy.Mu.SetBtn(int(keyCode), false)) {
            this.gameboy.Interrupter.Request(4)
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
        a := this.gameboy.Gpu.Buffer[y][x]
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
    this.gameboy.Gpu.Vblank = false
  }
  return true
}
