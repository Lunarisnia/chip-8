package emulator

import (
	"github.com/Lunarisnia/chip-8/internal/font"
	"github.com/Lunarisnia/chip-8/memory"
)

type Chip8 struct {
	Memory        [4096]byte
	PC            uint16 // Program Counter
	IR            uint16 // Index
	DisplayBuffer [2048]bool
	Stack         memory.Stack
}

func New() *Chip8 {
	return &Chip8{
		Memory: [4096]byte{},
		PC:     0,
		IR:     0,
		Stack:  memory.NewStack(),
	}
}

func (c *Chip8) Init() {
	c.initFonts()
}

func (c *Chip8) initFonts() {
	addr := 0x050
	for _, pack := range font.FontPacks {
		for _, f := range pack {
			c.Memory[addr] = f
			addr++
		}
	}
}
