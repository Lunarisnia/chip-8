package engine

import (
	"image/color"

	"github.com/Lunarisnia/chip-8/internal/emulator"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Engine struct {
	chip8 *emulator.Chip8
}

func NewEngine(emulator *emulator.Chip8) (*Engine, error) {
	return &Engine{
		chip8: emulator,
	}, nil
}

func (e *Engine) Update() error {
	e.chip8.Step()
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	for y, xRows := range e.chip8.DisplayBuffer {
		for x, on := range xRows {
			if !on {
				continue
			}
			vector.FillRect(screen, float32(10*x), float32(10*y), 10, 10, color.White, false)
		}
	}
}

func (e *Engine) Layout(w int, h int) (int, int) {
	return 640, 320
}
