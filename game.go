package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth       = 800
	ScreenHeight      = 600
	SpriteScaleFactor = 1
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
	score           int
}

func (g *Game) Update() error {
	g.player.Update()
	g.enemySpawnTimer.Update()
	if g.enemySpawnTimer.IsReady() {
		g.enemySpawnTimer.Reset()

		m := NewEnemy()
		g.enemies = append(g.enemies, m)
	}

	g.score += 3

	for _, e := range g.enemies {
		e.Update()
		if e.Collider().Intersects(g.player.Collider()) {
			g.score -= 10
		}
	}

	if g.score <= 0 {
		log.Fatal("Has palmao!")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen)
	// g.score = 100

	text.Draw(screen, fmt.Sprintf("%06d", g.score), ScoreFont, ScreenWidth/2-100, 50, color.White)

	for _, e := range g.enemies {
		e.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
