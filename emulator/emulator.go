package emulator

import (
	"github.com/Lunarisnia/chip-8/file"
	"github.com/Lunarisnia/chip-8/internal/font"
	"github.com/Lunarisnia/chip-8/memory"
)

const START_ADDR = 0x200

type Chip8 struct {
	Memory        [4096]byte
	Stack         memory.Stack
	DisplayBuffer [2048]bool
	Registers     [16]byte
	DelayTimer    byte
	SoundTimer    byte
	Keypad        [16]byte
	PC            uint16 // Program Counter
	IR            uint16 // Index
}

func New() *Chip8 {
	return &Chip8{
		Memory:        [4096]byte{},
		Stack:         memory.NewStack(),
		DisplayBuffer: [2048]bool{},
		Registers:     [16]byte{},
		Keypad:        [16]byte{},
		PC:            START_ADDR,
		IR:            0,
	}
}

func (c *Chip8) Init() {
	c.initFonts()
}

func (c *Chip8) LoadROM(romPath string) error {
	rom, err := file.ReadFile(romPath)
	if err != nil {
		return err
	}
	addr := START_ADDR
	for _, b := range rom {
		c.Memory[addr] = b
		addr++
	}
	return nil
}

func (c *Chip8) Run() {
	for i := START_ADDR; i < 4096/2; i++ {
		c.fetch()
	}
}

func (c *Chip8) fetch() uint16 {
	rawInstructions := c.Memory[c.PC : c.PC+2]
	c.PC += 2

	var instruction uint16
	instruction = uint16(rawInstructions[0])<<8 | uint16(rawInstructions[1])
	return instruction
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
