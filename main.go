package main

import (
	"log"

	"github.com/Lunarisnia/chip-8/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gameEngine, err := engine.NewEngine()
	if err != nil {
		log.Fatalln(err)
	}
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("CHIP-8")
	if err := ebiten.RunGame(gameEngine); err != nil {
		log.Fatal(err)
	}
}
