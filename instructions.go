package gb

//https://gbdev.io/gb-opcodes//optables/dark
//https://gb-archive.github.io/salvage/decoding_gbz80_opcodes/Decoding%20Gamboy%20Z80%20Opcodes.html

import (
	"os";
	//"fmt";
	"encoding/json";
	"strconv";
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
}

func initInstructionMap() map[int]Instruction {
	var instructions_unprefixed map[int]Instruction
	instructions_unprefixed = make(map[int]Instruction)
	var res map[string]interface{}

	data, err := os.ReadFile("data/Opcodes.json")
	if err == nil {
		json.Unmarshal([]byte(data), &res)
		unprefixed := map[string]interface{}(res["unprefixed"].(map[string]interface{}))
		for keyStr := range unprefixed {
			key, _ := strconv.ParseInt(keyStr, 16, 64)
			rawInstruction := map[string]interface{}(unprefixed[keyStr].(map[string]interface{}))
			instructions_unprefixed[int(key)] = Instruction {
				name: string(rawInstruction["mnemonic"].(string)),
				bytes: int(rawInstruction["bytes"].(float64)),
				//cycles: []int(rawInstruction["cycles"].(interface{})),
				immediate: bool(rawInstruction["immediate"].(bool)),
			}
		}
		/*
		prefixed := res["cbprefixed"]
		var instructions_prefixed map[int]Instruction
		for key := range map[string]interface{}(prefixed.(map[string]interface{})) {
			
		}
		*/
	}
	return instructions_unprefixed
}
