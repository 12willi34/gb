package gb

//https://gbdev.io/gb-opcodes//optables/dark
//https://gb-archive.github.io/salvage/decoding_gbz80_opcodes/Decoding%20Gamboy%20Z80%20Opcodes.html

import (
	"os";
	"fmt";
	"encoding/json";
)

type Operand struct {
	name string
	bytes int
	value int
	immediate bool
}

type Instruction struct {
	name string
	immediate bool
	operands []Operand
	cycles []int
	bytes int
	commend string
}

func initInstructionMap() map[int]Instruction {
	data, err := os.ReadFile("data/Opcodes.json")
	if err == nil {
		var res map[string]interface{}
		json.Unmarshal([]byte(data), &res)

		//temp
		x := res["unprefixed"]
		_, ok := x.(map[string]interface{})
		if ok {
			y := map[string]interface{}(x.(map[string]interface{}))["0x00"]
			fmt.Println(y)
		}
	}
	return map[int]Instruction {}
}
