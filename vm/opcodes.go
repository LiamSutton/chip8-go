package vm

import "fmt"

import "math/rand"

func (cpu *CPU) opcode0x0000() {
	switch cpu.opcode & 0x00FF {
	case 0x00E0: // Clear the screen
		cpu.opcode0x00E0()
	case 0x00EE:
		cpu.opcode0x00EE()
	}
}

func (cpu *CPU) opcode0x00E0() {
	fmt.Println("0x00E0: Clearing the screen")
	cpu.display = [64 * 32]uint8{}
	cpu.pc += 2
}

func (cpu *CPU) opcode0x00EE() {
	fmt.Printf("0x00EE: Returning from subroutine\n")
	cpu.pc = cpu.stack[cpu.sp] + 2
	cpu.sp--
}

func (cpu *CPU) opcode0x1000() {
	// Jump to address NNN
	nnn := cpu.opcode & 0x0FFF
	fmt.Printf("0x1000: Jumping to address NNN: 0x%X\n", nnn)
	cpu.pc = nnn
}

func (cpu *CPU) opcode0x2000() {
	// Call subroutine at NNN
	cpu.sp++
	cpu.stack[cpu.sp] = cpu.pc
	nnn := cpu.opcode & 0x0FFF
	cpu.pc = nnn
	fmt.Printf("0x2000: calling subroutine at 0x%X\n", nnn)
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

func (cpu *CPU) opcode0x8000() {
	switch cpu.opcode & 0x000F {
	case 0x0000:
		// Store Vy in Vx
		x := (cpu.opcode & 0x0F00) >> 8
		y := (cpu.opcode & 0x00F0) >> 4
		fmt.Printf("0x8000: Storing Vy: 0x%X in Vx: 0x%X\n", cpu.v[x], cpu.v[y])

		cpu.v[x] = cpu.v[y]

		cpu.pc += 2
	case 0x0001:
		// Set Vx = Vx OR Vy
		x := (cpu.opcode & 0x0F00) >> 8
		y := (cpu.opcode & 0x00F0) >> 4
		xy := cpu.v[x] | cpu.v[y]

		cpu.v[x] = xy
		fmt.Printf("0x8001: Bitwise OR on Vx: 0x%X and Vy: 0x%X, Result: Vx = 0x%X\n", cpu.v[x], cpu.v[y], xy)
		cpu.pc += 2
	case 0x0002:
		// Set Vx = Vx AND Vy
		x := (cpu.opcode & 0x0F00) >> 8
		y := (cpu.opcode & 0x00F0) >> 4
		xy := cpu.v[x] & cpu.v[y]
		cpu.v[x] = xy
		fmt.Printf("0x8002: Bitwise AND on Vx: 0x%X and Vy: 0x%X, Result: Vx = 0x%X\n", cpu.v[x], cpu.v[y], xy)
		cpu.pc += 2
	case 0x003:
		// Set Vx = Vx XOR Vy
		x := (cpu.opcode & 0x0F00) >> 8
		y := (cpu.opcode & 0x00F0) >> 4
		xy := cpu.v[x] | cpu.v[y]

		cpu.v[x] = xy
		fmt.Printf("0x8003: Bitwise XOR on Vx: 0x%X and Vy: 0x%X, Result: Vx = 0x%X\n", cpu.v[x], cpu.v[y], xy)
		cpu.pc += 2
	case 0x004:
		// Set Vx = Vx + Vy, set VF = carry.
		x := (cpu.opcode & 0x0F00) >> 8
		y := (cpu.opcode & 0x00F0) >> 4
		xy := cpu.v[x] + cpu.v[y]
		if xy > 0xFF {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}
		fmt.Printf("0x8004: Setting Vx: 0x%X += Vy: 0x%X, Vf = 0x%X\n", cpu.v[x], cpu.v[y], cpu.v[0xF])
		cpu.v[x] = xy
		cpu.pc += 2
	case 0x005:
		// Set Vx = Vx - Vy, set VF = NOT borrow.
		x := (cpu.opcode & 0x0F00) >> 8
		y := (cpu.opcode & 0x00F0) >> 4

		if cpu.v[x] > cpu.v[y] {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}
		fmt.Printf("0x8005: Setting Vx: 0x%X -= Vy: 0x%X, Vf = 0x%X\n", cpu.v[x], cpu.v[y], cpu.v[0xF])
		cpu.v[x] -= cpu.v[y]
		cpu.pc += 2
	case 0x006:
		x := (cpu.opcode & 0x0F00) >> 8
		cpu.v[0xF] = cpu.v[x] & 0x1
		fmt.Printf("0x800E: Shifting Vx: 0x%X Left 1 bit\n", cpu.v[x])
		cpu.v[x] >>= 1
		cpu.pc += 2
	case 0x00E:
		//Set Vx = Vx SHL 1.
		x := (cpu.opcode & 0x0F00) >> 8
		cpu.v[0xF] = cpu.v[x] >> 7
		fmt.Printf("0x800E: Shifting Vx: 0x%X Left 1 bit\n", cpu.v[x])
		cpu.v[x] <<= 1
		cpu.pc += 2

	default:
		fmt.Printf("Unimplemented opcode: 0x%X\n", cpu.opcode)
	}
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

func (cpu *CPU) opcode0xB000() {
	// Jump to location nnn + V0.
	nnn := cpu.opcode & 0x0FFF
	cpu.pc = nnn + uint16(cpu.v[0x0])
	fmt.Printf("0xB00: Jumping to address: 0x%X\n", cpu.pc)
}

func (cpu *CPU) opcode0xC000() {
	// Set Vx = random byte AND kk.
	x := uint16((cpu.opcode & 0x0F00)) >> 8
	kk := uint8(cpu.opcode & 0x00FF)
	r := uint8(rand.Intn(255))

	cpu.v[x] = r & kk
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

func (cpu *CPU) opcode0xE000() {
	switch cpu.opcode * 0x00FF {
	case 0x009E:
		// Skip next instruction if key with the value of Vx is pressed.
		x := (cpu.opcode & 0x0F00) >> 8
		if cpu.keyboard[cpu.v[x]] {
			fmt.Printf("0xE09E: Key: 0x%X is pressed, skipping next instruction\n", cpu.v[x])
			cpu.pc += 4
		} else {
			fmt.Printf("0xE09E: Key: 0x%X not pressed, not skipping next instruction\n", cpu.v[x])
			cpu.pc += 2
		}
	case 0x00A1:
		// Skip next instruction if key with the value of Vx is NOT pressed
		x := (cpu.opcode & 0x0F00) >> 8
		if !cpu.keyboard[cpu.v[x]] {
			fmt.Printf("0xE0A1: Key: 0x%X not pressed, skipping next instruction\n", cpu.v[x])
			cpu.pc += 4
		} else {
			fmt.Printf("0xE0A1: Key: 0x%X pressed, not skipping next instruction\n", cpu.v[x])
			cpu.pc += 2
		}
	}
}

func (cpu *CPU) opcode0xF000() {
	switch cpu.opcode & 0x00FF {
	case 0x0007:
		fmt.Printf("0xF007: Setting Vx = delay timer\n")
		// Set Vx -> delay timer
		x := (cpu.opcode & 0x0F00) >> 8
		cpu.v[x] = cpu.delayTimer
	case 0x000A:
		// Wait for a key press, store the value of the key in Vx.
		x := (cpu.opcode & 0x0F00) >> 8
		fmt.Printf("0x000A: Waiting for keypress: HALTING\n")
		for {
			for i := uint8(0); i < 16; i++ {
				if cpu.keyboard[i] == true {
					fmt.Printf("0x00A: Got keypress: 0x%v\n", i)
					cpu.v[x] = i
					break
				}
			}
		}
	case 0x0015:
		// Set delay timer = Vx
		fmt.Printf("0xF015: Setting delay timer = Vx\n")
		x := (cpu.opcode & 0x0F00) >> 8
		cpu.delayTimer = cpu.v[x]
	case 0x0018:
		// Set sound timer = Vx
		fmt.Printf("Setting sound timer = Vx\n")
		x := (cpu.opcode & 0x0F00) >> 8
		cpu.soundTimer = cpu.v[x]
	case 0x001E:
		// Adds Vx to I, if overflow occurs VF = 1
		x := (cpu.opcode & 0x0F00) >> 8
		result := cpu.i + uint16(cpu.v[x])
		if result >= 0xFFF {
			// Overflow
			cpu.v[0xF] = 1
		} else {
			// No overflow
			cpu.v[0xF] = 0
		}

		cpu.i = result
		cpu.pc += 2
	case 0x0029:
		// Set I = location of sprite for digit Vx.
		x := uint16(((cpu.opcode & 0x0F00) >> 8)) * 5
		location := uint16(cpu.v[x])
		cpu.i = location
		cpu.pc += 2
	case 0x0033:
		// Store BCD representation of Vx in memory locations I, I+1, and I+2.
		x := (cpu.opcode & 0x0F00) >> 8

		cpu.memory[cpu.i] = cpu.v[x] / 100
		cpu.memory[cpu.i+1] = (cpu.v[x] / 10) % 10
		cpu.memory[cpu.i+2] = (cpu.v[x] % 100) % 10
		cpu.pc += 2

		fmt.Printf("Storing BCD of Vx: 0x%X in memory locations I, I+1, I+2, I: 0x%X\n", cpu.v[x], cpu.i)
	case 0x0055:
		// Store registers V0 through Vx in memory starting at location I.
		for j := uint16(0); j < 16; j++ {
			cpu.memory[cpu.i+j] = cpu.v[j]
		}
		cpu.pc += 2
	case 0x0065:
		for j := uint16(0); j < 16; j++ {
			cpu.v[j] = cpu.memory[cpu.i+j]
		}
		cpu.pc += 2
	default:
		fmt.Printf("Unimplemented opcode: 0x%X\n", cpu.opcode)
	}
}
