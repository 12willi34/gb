package gb

import ()

type GameBoy struct {
    Mu *memoryunit
    Processor *cpu

    //ToDo - clock cycles
    //ToDo - all states of the gameboy
}

func NewGameBoy(rom []byte) GameBoy {
    mu := NewMemoryUnit()
    processor := NewCPU(rom, &mu)
    gameboy := GameBoy {
        Mu: &mu,
        Processor: &processor,
    }
    return gameboy
}
