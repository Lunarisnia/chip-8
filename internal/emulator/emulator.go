package emulator

import (
	"github.com/Lunarisnia/chip-8/internal/file"
	"github.com/Lunarisnia/chip-8/internal/font"
	"github.com/Lunarisnia/chip-8/internal/memory"
)

const START_ADDR = 0x200
const FONT_START_ADDR = 0x050

type Chip8 struct {
	Memory        [4096]byte
	Stack         memory.Stack
	DisplayBuffer [32][64]bool
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
		DisplayBuffer: [32][64]bool{},
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

func (c *Chip8) Step() {
	opcode := c.fetch()
	c.decode(opcode)
}

func (c *Chip8) fetch() uint16 {
	rawInstructions := c.Memory[c.PC : c.PC+2]
	c.PC += 2

	var instruction uint16
	instruction = uint16(rawInstructions[0])<<8 | uint16(rawInstructions[1])
	return instruction
}

func (c *Chip8) decode(opcode uint16) {
	firstNibble := opcode & 0xF000

	switch firstNibble >> 12 {
	case 0x0:
		code := opcode & 0x0FFF
		if code == 0x0E0 {
			c.DisplayBuffer = [32][64]bool{}
		}
	case 0x1: // Jump
		jumpAddr := opcode & 0x0FFF
		c.PC = jumpAddr
	case 0x6: // Set Register VX
		xAddr := opcode & 0x0F00 >> 8
		xValue := byte(opcode & 0x00FF)
		c.Registers[xAddr] = xValue
	case 0x7: // Add Register VX
		xAddr := opcode & 0x0F00 >> 8
		xValue := byte(opcode & 0x00FF)
		c.Registers[xAddr] += xValue
	case 0xA: // Point Index to addr
		indexAddr := opcode & 0x0FFF
		c.IR = indexAddr
	case 0xD: // Draw
		xAddr := opcode & 0x0F00 >> 8
		yAddr := opcode & 0x00F0 >> 4

		w, h := c.DisplayResolution()
		xCoord := c.Registers[xAddr] % byte(w)
		yCoord := c.Registers[yAddr] % byte(h)
		c.Registers[0xF] = 0
		value := uint16(opcode & 0x000F)

		for row := range value {
			sprite := c.Memory[c.IR+row]
			for i := range 8 {
				mask := byte(0b1000_0000) >> byte(i)
				signal := (sprite & mask) >> byte(7-i)
				currentY := yCoord + byte(row)
				if currentY >= byte(h) || xCoord >= byte(w) {
					break
				}
				if on := c.DisplayBuffer[currentY][xCoord+byte(i)]; on && signal == 1 {
					c.DisplayBuffer[currentY][xCoord+byte(i)] = false
					c.Registers[0xF] = 1
				} else if signal == 1 {
					c.DisplayBuffer[currentY][xCoord+byte(i)] = true
				}
			}
		}
	}
}

func (c *Chip8) initFonts() {
	addr := FONT_START_ADDR
	for _, pack := range font.FontPacks {
		for _, f := range pack {
			c.Memory[addr] = f
			addr++
		}
	}
}

func (c *Chip8) DisplayResolution() (int, int) {
	return 64, 32
}
