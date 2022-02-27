package gb

import (
  "fmt"
  "github.com/veandco/go-sdl2/sdl"
)

func showWindow() {
  window, err := sdl.CreateWindow("GameBoy", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 160, 144, sdl.WINDOW_SHOWN)

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

  rect := sdl.Rect{0, 0, 100, 100}
  surface.FillRect(&rect, 0xffff0000)
  window.UpdateSurface()

  /*
  running := true
  for(running) {
    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
      switch event.(type) {
      case *sdl.QuitEvent:
        println("Quit")
        running = false
        break
      }
    }
  }
  */
}
