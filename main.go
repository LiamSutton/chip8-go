package main

import (
	"fmt"
	"github.com/LiamSutton/chip8-go/vm"
)

func main() {
	fmt.Println("Hello World!")
	cpu := vm.NewCPU()
	cpu.ResetCPU()
	cpu.PrintStatus()
}
