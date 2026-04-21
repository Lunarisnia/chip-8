package main

import (
	"log"
	"os"

	"github.com/Lunarisnia/chip-8/internal/emulator"
	"github.com/Lunarisnia/chip-8/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: Test with another more complete test roms

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("invalid argument count, need 2")
	}
	romPath := os.Args[1]

	chip8 := emulator.New()
	err := chip8.LoadROM(romPath)
	if err != nil {
		log.Fatalln(err)
	}

	gameEngine, err := engine.NewEngine(chip8)
	if err != nil {
		log.Fatalln(err)
	}
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("CHIP-8")
	if err := ebiten.RunGame(gameEngine); err != nil {
		log.Fatal(err)
	}
}
