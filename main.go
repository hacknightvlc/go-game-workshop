package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	err := ebiten.RunGame(NewGame())
	if err != nil {
		panic(err)
	}
}
