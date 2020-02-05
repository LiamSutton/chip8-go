package vm

import "fmt"

// Fontset : the characters 0-F represented in hex
var fontset = []uint8{
	0xF0, 0x90, 0x90, 0x90, 0xf0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x20, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

// CPU - the chip-8
type CPU struct {
	memory     [4096]uint8    // 4k of RAM
	display    [64 * 32]uint8 // The display of the Chip-8 is a 64 * 32 monochrome grid
	v          [16]uint8      // 16 8-bit registers
	i          uint16         // The index register
	pc         uint16         // The program counter
	opcode     uint16         // The currrent opcode
	stack      [16]uint16     // Used to store return addresses when subroutines are called
	sp         uint16         // The stack pointer
	keyboard   [16]bool       // Keyboard keys are mapped for values 0x0 -> 0xF
	drawFlag   bool           // Whether the screen should draw
	delayTimer uint8          // Used for timing events
	soundTimer uint8          // Used for sound effects, when value != 0, a beeping sound is made
}

// NewCPU creates and returns a new CPU
func NewCPU() *CPU {
	cpu := CPU{}

	return &cpu
}

// PrintStatus displays the current status of the CPU
func (cpu *CPU) PrintStatus() {
	fmt.Println("\n-- PRINTING STATUS OF CPU --")
	fmt.Printf("\nV Registers: \n%v \n", cpu.v)
	fmt.Printf("\nStack: \n%v \n", cpu.stack)
	fmt.Printf("\nStack Pointer: \nHEX: 0x%X \nDEC: %d \n", cpu.sp, cpu.sp)
	fmt.Printf("\nIndex Register: \nHEX: 0x%X \nDEC: %d \n", cpu.i, cpu.i)
	fmt.Printf("\nProgram Counter: \nHEX: 0x%X \nDEC: %d \n", cpu.pc, cpu.pc)
	fmt.Printf("\nCurrent Opcode: \nHEX: 0x%X \nDEC: %d \n", cpu.opcode, cpu.opcode)
	fmt.Printf("\nDraw Flag: \n%v \n", cpu.drawFlag)
	fmt.Printf("\nDelay Timer: \n%v \n", cpu.delayTimer)
	fmt.Printf("\nSound Timer: \n%v \n", cpu.soundTimer)
	fmt.Printf("\nKeyboard: \n%v \n", cpu.keyboard)
	fmt.Println("\n-- FINISHED PRINTING STATUS OF CPU --")
}

// ResetCPU sets the CPU to an initial state ready to run
func (cpu *CPU) ResetCPU() {
	cpu.pc = 0x200                 // Program counter starts ot 0x200
	cpu.opcode = 0x0               // Reset current opcode
	cpu.i = 0x0                    // Reset index register
	cpu.sp = 0x0                   // Reset stack pointer
	cpu.display = [64 * 32]uint8{} // Clear the display
	cpu.stack = [16]uint16{}       // Clear the stack
	cpu.v = [16]uint8{}            // Clear registers V0-VF
	cpu.memory = [4096]uint8{}     // Clear the memory

	for i := 0x0; i < 0x50; i++ { // Load the Fontset into memory
		cpu.memory[i] = fontset[i]
	}

	cpu.delayTimer = 0x0 // Reset the delay timer
	cpu.soundTimer = 0x0 // Reset the sound timer
}

// GetDisplay will return the current display for the CPU
func (cpu *CPU) GetDisplay() [64 * 32]uint8 {
	return cpu.display
}

// ShouldDraw returns the current state of the drawFlag
func (cpu *CPU) ShouldDraw() bool {
	return cpu.drawFlag
}

// EmulateCycle will fetch, decode and execute a single opcode
func (cpu *CPU) EmulateCycle() {
	// Fetch Opcode
	cpu.opcode = uint16(cpu.memory[cpu.pc])<<8 | uint16(cpu.memory[cpu.pc+1])

	switch cpu.opcode & 0xF000 {
	case 0x0000:
		cpu.opcode0x0000()
	case 0x1000:
		cpu.opcode0x1000()
	case 0x3000:
		cpu.opcode0x3000()
	case 0x4000:
		cpu.opcode0x4000()
	case 0x5000:
		cpu.opcode0x5000()
	case 0x6000:
		cpu.opcode0x6000()
	case 0x7000:
		cpu.opcode0x7000()
	case 0xA000:
		cpu.opcode0xA000()
	case 0xD000:
		cpu.opcode0xD000()
	default:
		fmt.Printf("Unimplemented Opcode: [0x%X]\n", cpu.opcode)
	}
}
