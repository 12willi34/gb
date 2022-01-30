package main

import (
	"fmt";
	"gb";
	"os";
)

func test_memory() {
	fmt.Println("testing memory")

	var key_8 uint16 = 0x1000
	var val_8 uint8 = 32
	var key_16 uint16 = 0x1000 + 2
	var val_16 uint16 = 512

	mmu := gb.NewMemoryUnit()
	mmu.Write_8(key_8, val_8)
	fmt.Println("read/write 8 bit addr:", mmu.Read_8(key_8) == val_8)

	mmu = gb.NewMemoryUnit()
	mmu.Write_16(key_16, val_16)
	fmt.Println("read/write 16 bit addr:", mmu.Read_16(key_16) == val_16)
}

func test_cpu() {
	fmt.Println("testing cpu")
	firstFile, _ := os.ReadDir("../rom/")
	rom, _ := os.ReadFile("../rom/" + firstFile[0].Name())
	processor := gb.NewCPU(rom)
	for i := 0; i < 10; i++ {
		processor.Step()
	}
}

func main() {
	test_memory()
	test_cpu()
}
