package main

import "github.com/Lunarisnia/chip-8/emulator"

func main() {
	chip8 := emulator.New()
	chip8.Init()
}
