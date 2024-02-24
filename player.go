package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
	reverse  bool
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
	if p.reverse {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(float64(p.sprite.Bounds().Dy()*SpriteScaleFactor), 0)
	}
	opts.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, opts)
}

func (p *Player) Update() {
	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.position.Y < float64(ScreenHeight-p.sprite.Bounds().Dx()*SpriteScaleFactor) {
			p.position.Y += speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.position.Y > 0 {
			p.position.Y -= speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.position.X > 0 {
			p.position.X -= speed
		}
		p.reverse = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.position.X < float64(ScreenWidth-p.sprite.Bounds().Dy()*SpriteScaleFactor) {
			p.position.X += speed
		}
		p.reverse = false
	}
}
