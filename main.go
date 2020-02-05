package main

import (
	"github.com/LiamSutton/chip8-go/vm"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var roms = []string{"roms/Chip8 Picture.ch8", "roms/test_opcode.ch8"}

var win *pixelgl.Window

const (
	sizeX, sizeY              = 64, 32
	screenWidth, screenHeight = float64(1024), float64(768)
)

func main() {
	cpu := vm.NewCPU()
	cpu.ResetCPU()

	// rom := vm.ReadROM(roms[0])

	// cpu.LoadROM(rom)

	// for i := 0; i < 100; i++ {
	// 	cpu.EmulateCycle()
	// }
	pixelgl.Run(run)
}

func run() {
	cpu := vm.NewCPU()
	cpu.ResetCPU()
	rom := vm.ReadROM(roms[1])
	cpu.LoadROM(rom)

	programSetup()

	for !win.Closed() {
		cpu.EmulateCycle()
		if cpu.ShouldDraw() {
			draw(cpu.GetDisplay())
		} else {
			win.UpdateInput()
		}
	}
}

func programSetup() {
	cfg := pixelgl.WindowConfig{
		Title:  "Chip-8 Go",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
}

func draw(display [64 * 32]uint8) {
	win.Clear(colornames.Black)                            // Clear the window
	imd := imdraw.New(nil)                                 // Create a new imdraw object
	imd.Color = pixel.RGB(1, 1, 1)                         // Set the color to white
	screenWidth := win.Bounds().W()                        // The width of the screen
	width, height := screenWidth/sizeX, screenHeight/sizeY //
	for x := 0; x < 64; x++ {                              // The screen is 64 pixels across
		for y := 0; y < 32; y++ { // The screen is 32 pixels high
			if display[(31-y)*64+x] == 1 {
				imd.Push(pixel.V(width*float64(x), height*float64(y)))
				imd.Push(pixel.V(width*float64(x)+width, height*float64(y)+height))
				imd.Rectangle(0)
			}
		}
	}
	imd.Draw(win)
	win.Update()
}
