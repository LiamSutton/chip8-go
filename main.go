package main

import (
	"github.com/LiamSutton/chip8-go/vm"
)

var roms = []string{"roms/Chip8 Picture.ch8"}

func main() {
	cpu := vm.NewCPU()
	cpu.ResetCPU()

	rom := vm.ReadROM(roms[0])

	cpu.LoadROM(rom)

	for i := 0; i < 10; i++ {
		cpu.EmulateCycle()
	}
}
