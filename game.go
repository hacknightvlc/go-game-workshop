package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth       = 800
	ScreenHeight      = 600
	SpriteScaleFactor = 4
	enemySpawnTime    = 500 * time.Millisecond
)

func NewGame() *Game {
	audio := NewAudio()
	audio.PlayMusic()

	return &Game{
		player:          NewPlayer(),
		world:           NewWorld(),
		enemySpawnTimer: NewTimer(enemySpawnTime),
		audio:           audio,
	}
}

type Game struct {
	player          *Player
	world           *World
	enemies         []*Enemy
	enemySpawnTimer *Timer
	score           int
	audio           *Audio
}

func (g *Game) Update() error {
	g.player.Update()
	g.enemySpawnTimer.Update()
	if g.enemySpawnTimer.IsReady() {
		g.enemySpawnTimer.Reset()

		m := NewEnemy()
		g.enemies = append(g.enemies, m)
	}

	for _, e := range g.enemies {
		e.Update()
		if e.Collider().Intersects(g.player.Collider()) {
			e.Stomp()
			g.audio.PlaySound()
			g.score++
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen)

	for _, e := range g.enemies {
		e.Draw(screen)
	}
	text.Draw(screen, fmt.Sprintf("%v", g.score), ScoreFont, ScreenWidth/2-100, 50, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
