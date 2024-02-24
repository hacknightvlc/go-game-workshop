package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth       = 800
	ScreenHeight      = 600
	SpriteScaleFactor = 4
)

func NewGame() *Game {
	return &Game{
		player: NewPlayer(),
		world:  NewWorld(),
	}
}

type Game struct {
	player *Player
	world  *World
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
