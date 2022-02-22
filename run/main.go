package main

import (
	"gb";
	"os";
)

func main() {
  boot, _ := os.ReadFile("../boot/dmg_boot.bin")
	roms, _ := os.ReadDir("../rom/")
	rom, _ := os.ReadFile("../rom/" + roms[0].Name())
  gb.NewGameBoy(boot, rom).Init()
}
