package vm

import "io/ioutil"

// ReadROM will load a given file into a variable
func ReadROM(filename string) []byte {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return data
}

// LoadROM will take a rom in the format of a byte array and load it into the CPU's memory
func (cpu *CPU) LoadROM(rom []byte) {
	for i := 0; i < len(rom); i++ {
		cpu.memory[0x200+i] = rom[i]
	}
}
