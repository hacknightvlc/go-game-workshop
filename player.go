package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := Vector{
		X: ScreenWidth/2 - halfW*SpriteScaleFactor,
		Y: ScreenHeight/2 - halfH*SpriteScaleFactor,
	}
	return &Player{
		position: position,
		sprite:   sprite,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Scale(SpriteScaleFactor, SpriteScaleFactor)
	opts.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, opts)
}
