package emulator

import (
	"testing"

	"github.com/Lunarisnia/chip-8/internal/font"
)

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
