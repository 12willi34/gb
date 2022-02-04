package main

import (
	"fmt";
	"gb";
	"os";
)

func test_gb() {
	fmt.Println("testing gameboy")
	roms, _ := os.ReadDir("../rom/")
	rom, _ := os.ReadFile("../rom/" + roms[0].Name())
    gameboy := gb.NewGameBoy(rom)
	for ((*gameboy.Processor).Step() > -1) {
		continue
	}
}

func main() {
	test_gb()
}
