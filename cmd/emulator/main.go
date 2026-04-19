package main

import (
	"log"

	"github.com/Lunarisnia/chip-8/internal/emulator"
)

func main() {
	chip8 := emulator.New()
	chip8.Init()
	err := chip8.LoadROM("./testroms/IBM-logo.ch8")
	if err != nil {
		log.Fatalln(err)
	}

	chip8.Run()
}
