package vm

func (cpu *CPU) opcode0X0000() {
	switch cpu.opcode & 0x00FF {
	case 0x0000: // 0x00E0 : Clear the screen
		cpu.display = [64 * 32]uint8{}
		cpu.pc += 2
	case 0x00E0: // 0x00EE : Return from a subroutine
		cpu.pc = cpu.stack[cpu.sp]
		cpu.sp--
	}
}
