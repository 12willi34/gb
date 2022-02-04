package main

import (
	"fmt";
	"gb";
	"os";
)

func test_gb() {
	fmt.Println("testing gameboy")
	firstFile, _ := os.ReadDir("../rom/")
	rom, _ := os.ReadFile("../rom/" + firstFile[0].Name())
    gameboy := gb.NewGameBoy(rom)
	for i := 0; i < 10; i++ {
		(*gameboy.Processor).Step()
	}
}

func main() {
	test_gb()
}
