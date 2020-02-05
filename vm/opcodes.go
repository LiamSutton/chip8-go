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
	nnn := cpu.opcode & 0x0FFF
	fmt.Printf("0x1000: Jumping to address NNN: 0x%X\n", nnn)
	cpu.pc = nnn
}

func (cpu *CPU) opcode0x3000() {
	// skip next instruction if Vx == NN
	x := (cpu.opcode & 0x0F00) >> 8
	nn := uint8(cpu.opcode & 0x00FF)
	if cpu.v[x] == nn {
		fmt.Printf("0x3000: Vx: 0x%X ==  NN: 0x%X, skipping next instruction\n", cpu.v[x], nn)
		cpu.pc += 4
	} else {
		fmt.Printf("0x3000: Vx: 0x%X != NN: 0x%X, not skipping next instruction\n", cpu.v[x], nn)
		cpu.pc += 2
	}
}

func (cpu *CPU) opcode0x4000() {
	// Skip next instruction if Vx != NN
	x := (cpu.opcode & 0x0F00) >> 8
	nn := uint8(cpu.opcode & 0x00FF)

	if cpu.v[x] != nn {
		fmt.Printf("0x4000: Vx: 0x%X != NN: 0x%X, skipping next instruction\n", cpu.v[x], nn)
		cpu.pc += 4
	} else {
		fmt.Printf("0x4000: Vx: 0x%X ==  NN: 0x%X, not skipping next instruction\n", cpu.v[x], nn)
		cpu.pc += 2
	}
}

func (cpu *CPU) opcode0x5000() {
	// Skip next instruction if Vx == Vy
	x := (cpu.opcode & 0x0F00) >> 8
	y := (cpu.opcode & 0x00F0) >> 4

	if cpu.v[x] == cpu.v[y] {
		fmt.Printf("0x5000: Vx: 0x%X == Vy: 0x%X, skipping next instruction", cpu.v[x], cpu.v[y])
		cpu.pc += 4
	} else {
		fmt.Printf("0x5000: Vx: 0x%X != Vy: 0x%X, not skipping next instruction", cpu.v[x], cpu.v[y])
		cpu.pc += 2
	}
}

func (cpu *CPU) opcode0x6000() {
	// Put the value KK into the register Vx
	x := (cpu.opcode & 0x0F00) >> 8
	kk := uint8(cpu.opcode & 0x00FF)
	fmt.Printf("0x6000: Putting the value KK: 0x%X  into register Vx: 0x%X\n", kk, x)
	cpu.v[x] = kk
	cpu.pc += 2
}

func (cpu *CPU) opcode0x7000() {
	// Adds NN to Vx (carry flag unchanged)
	x := (cpu.opcode & 0x0F00) >> 8
	nn := uint8(cpu.opcode & 0x00FF)
	fmt.Printf("0x7000: Adding NN: 0x%X to Vx: 0x%X (carry flag unchanged)\n", nn, x)
	cpu.v[x] += nn
	cpu.pc += 2
}

func (cpu *CPU) opcode0x9000() {
	// Skip next instruction if Vx != Vy
	x := (cpu.opcode & 0x0F00) >> 8
	y := (cpu.opcode & 0x00F0) >> 4

	if cpu.v[x] != cpu.v[y] {
		fmt.Printf("0x5000: Vx: 0x%X != Vy: 0x%X, skipping next instruction", cpu.v[x], cpu.v[y])
		cpu.pc += 4
	} else {
		fmt.Printf("0x5000: Vx: 0x%X == Vy: 0x%X, not skipping next instruction", cpu.v[x], cpu.v[y])
		cpu.pc += 2
	}
}

func (cpu *CPU) opcode0xA000() {
	// Set I register -> NNN
	nnn := cpu.opcode & 0x0FFF
	fmt.Printf("0xA000: Setting I = 0x%X\n", nnn)
	cpu.i = nnn
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
