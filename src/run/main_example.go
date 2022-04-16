package main

import (
  "gb"
  "os"
  "fmt"
)

var boot [256]uint8 = [256]uint8 {
  //...
}

func main() {
  rom, err := os.ReadFile(os.Args[1])
  if(err != nil) {
    fmt.Println("could read rom:", os.Args[1])
    panic(err)
  }
  gameboy := gb.NewGameBoy(boot, rom)
  (&gameboy).Init()
}
