package main

import (
	"fmt";
	"gb"
)

func main() {
	test_memory()
}

func test_memory() {
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
