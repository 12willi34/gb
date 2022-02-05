package main

import (
	"gb";
	"os";
)

func main() {
	roms, _ := os.ReadDir("../rom/")
	rom, _ := os.ReadFile("../rom/" + roms[0].Name())
  gb.NewGameBoy(rom).Init()
}
