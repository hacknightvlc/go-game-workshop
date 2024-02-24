package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth       = 800
	ScreenHeight      = 600
	SpriteScaleFactor = 4
	enemySpawnTime    = 500 * time.Millisecond
)

func NewGame() *Game {
	return &Game{
		player:          NewPlayer(),
		world:           NewWorld(),
		enemySpawnTimer: NewTimer(enemySpawnTime),
	}
}

type Game struct {
	player          *Player
	world           *World
	enemies         []*Enemy
	enemySpawnTimer *Timer
}

func (g *Game) Update() error {
	g.player.Update()
	g.enemySpawnTimer.Update()
	if g.enemySpawnTimer.IsReady() {
		g.enemySpawnTimer.Reset()

		m := NewEnemy()
		g.enemies = append(g.enemies, m)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
