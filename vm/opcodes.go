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

func (cpu *CPU) opcode0x1000() {
	// Jump to address NNN
	fmt.Printf("0x1000: Jumping to address NNN: 0x%X\n", cpu.opcode&0x0FFF)
	cpu.pc = cpu.opcode & 0x0FFF
}

func (cpu *CPU) opcode0x3000() {
	// skip next instruction if Vx == NN
	if cpu.v[(cpu.opcode&0x0F00)>>8] == uint8(cpu.opcode&0x00FF) {
		fmt.Printf("0x3000: 0x%X = 0x%X, skipping next instruction\n", (cpu.opcode&0x0F00)>>8, uint8(cpu.opcode&0x00FF))
		cpu.pc += 4
	} else {
		fmt.Printf("0x3000: 0x%X != 0x%X, not skipping next instruction\n", (cpu.opcode&0x0F00)>>8, uint8(cpu.opcode&0x00FF))
		cpu.pc += 2
	}
}

func (cpu *CPU) opcode0x6000() {
	// Put the value KK into the register Vx
	fmt.Printf("0x6000: Putting the value KK: 0x%X  into register Vx: 0x%X\n", uint8(cpu.opcode&0x00FF), (cpu.opcode&0x0F00)>>8)
	cpu.v[(cpu.opcode&0x0F00)>>8] = uint8(cpu.opcode & 0x00FF)
	cpu.pc += 2
}

func (cpu *CPU) opcode0x7000() {
	// Adds NN to Vx (carry flag unchanged)
	fmt.Printf("0x7000: Adding NN: 0x%X to Vx: 0x%X (carry flag unchanged)\n", cpu.opcode&0x00FF, cpu.opcode&0x0F00>>8)
	cpu.v[(cpu.opcode&0x0F00)>>8] += uint8(cpu.opcode & 0x00FF)
	cpu.pc += 2
}

func (cpu *CPU) opcode0xA000() {
	// Set I register -> NNNN
	fmt.Printf("0xA000: Setting I = 0x%X\n", cpu.opcode&0x0FFF)
	cpu.i = cpu.opcode & 0x0FFF
	cpu.pc += 2
}

func (cpu *CPU) opcode0xD000() {
	x := uint16(cpu.v[(cpu.opcode&0x0F00)>>8])
	y := uint16(cpu.v[(cpu.opcode&0x00F0)>>4])
	height := cpu.opcode & 0x000F
	fmt.Printf("0xD000: Drawing a sprite at X: 0x%X, Y: 0x%X with a height of 0x%X\n", x, y, height)
	cpu.v[0xF] = 0

	for yline := uint16(0); yline < height; yline++ {
		pixel := uint16(cpu.memory[cpu.i+yline])
		for xline := uint16(0); xline < 8; xline++ {
			index := (x + xline + ((y + yline) * 64))
			if index > uint16(len(cpu.display)) {
				continue
			}
			if pixel&(0x80>>xline) != 0 {
				if cpu.display[index] == 1 {
					cpu.v[0xF] = 1
				}
				cpu.display[index] ^= 1
			}
		}
	}
	cpu.drawFlag = true
	cpu.pc += 2
}
