package engine

import (
	"image/color"

	"github.com/Lunarisnia/chip-8/internal/emulator"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Engine struct {
	chip8 *emulator.Chip8

	pixelCount int
}

func NewEngine() (*Engine, error) {
	chip8 := emulator.New()
	chip8.Init()
	err := chip8.LoadROM("./testroms/IBM-logo.ch8")
	if err != nil {
		return nil, err
	}
	return &Engine{
		chip8:      chip8,
		pixelCount: 240,
	}, nil
}

func (e *Engine) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		e.chip8.DisplayBuffer[e.pixelCount] = false
		e.pixelCount--
		e.chip8.DisplayBuffer[e.pixelCount] = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		e.chip8.DisplayBuffer[e.pixelCount] = false
		e.pixelCount++
		e.chip8.DisplayBuffer[e.pixelCount] = true
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	w, _ := e.chip8.DisplayResolution()
	for i, on := range e.chip8.DisplayBuffer {
		if !on {
			continue
		}
		x := i % w
		y := i / w

		vector.FillRect(screen, float32(10*x), float32(10*y), 10, 10, color.White, false)
	}
}

func (e *Engine) Layout(w int, h int) (int, int) {
	return 640, 320
}
