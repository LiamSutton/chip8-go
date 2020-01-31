package vm

// CPU - the chip-8
type CPU struct {
	memory     [4096]uint8 // 4k of RAM
	v          [16]uint8   // 16 8-bit registers
	i          uint16      // The index register
	pc         uint16      // The program counter
	opcode     uint16      // The currrent opcode
	stack      [16]uint16  // Used to store return addresses when subroutines are called
	sp         uint16      // The stack pointer
	keyboard   [16]bool    // Keyboard keys are mapped for values 0x0 -> 0xF
	drawFlag   bool        // Whether the screen should draw
	delayTimer uint8       // Used for timing events
	soundTimer uint8       // Used for sound effects, when value != 0, a beeping sound is made
}
