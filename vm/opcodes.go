package vm

import "fmt"

func (cpu *CPU) opcode0x0000() {
	switch cpu.opcode & 0x00FF {
	case 0x00E0: // Clear the screen
		fmt.Println("0x00E0: Clearing the screen")
		cpu.display = [64 * 32]uint8{}
		cpu.pc += 2

	case 0x00EE:
		fmt.Println("Return")
	}
}

func (cpu *CPU) opcode0xA000() {
	// Set I register -> NNNN
	fmt.Println("0xA000: Setting I -> NNN")
	cpu.i = cpu.opcode & 0x0FFF
	cpu.pc += 2
}

func (cpu *CPU) opcode0x6000() {
	// Put the value KK into the register Vx
	fmt.Println("0x6000: Putting the value KK into register Vx")
	cpu.v[(cpu.opcode&0x0F00)>>8] = uint8(cpu.opcode & 0x00FF)
	cpu.pc += 2
}

func (cpu *CPU) opcode0xD000() {
	// Draws sprite (dummy method for now)
	fmt.Println("0xD000: Drawing sprite to the screen")
	cpu.pc += 2
}
