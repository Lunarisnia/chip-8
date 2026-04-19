package emulator

import (
	"testing"

	"github.com/Lunarisnia/chip-8/internal/font"
)

func TestFetch(t *testing.T) {
	c := New()
	c.Memory[START_ADDR] = 0xAB
	c.Memory[START_ADDR+1] = 0xCD
	c.Memory[START_ADDR+2] = 0x12
	c.Memory[START_ADDR+3] = 0x34

	if got := c.fetch(); got != 0xABCD {
		t.Errorf("fetch[0]: want 0xABCD, got 0x%04X", got)
	}
	if c.PC != START_ADDR+2 {
		t.Errorf("PC after fetch[0]: want 0x%03X, got 0x%03X", START_ADDR+2, c.PC)
	}

	if got := c.fetch(); got != 0x1234 {
		t.Errorf("fetch[1]: want 0x1234, got 0x%04X", got)
	}
	if c.PC != START_ADDR+4 {
		t.Errorf("PC after fetch[1]: want 0x%03X, got 0x%03X", START_ADDR+4, c.PC)
	}
}

func TestInit_FontsWrittenCorrectly(t *testing.T) {
	c := New()
	c.initFonts()

	addr := 0x050
	for i, f := range font.FontPacks {
		for j, b := range f {
			got := c.Memory[addr]
			if got != b {
				t.Errorf("font[%d][%d] at 0x%03X: want 0x%02X, got 0x%02X", i, j, addr, b, got)
			}
			addr++
		}
	}
}
