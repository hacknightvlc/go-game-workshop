package main

import "github.com/hajimehoshi/ebiten/v2"

type World struct {
	sprite *ebiten.Image
}

func NewWorld() *World {
	return &World{
		sprite: WorldSprite,
	}
}

func (w *World) Draw(screen *ebiten.Image) {
	bounds := w.sprite.Bounds()

	var scaleFactor float64 = 5

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scaleFactor, scaleFactor)
			tileX := float64(x) * float64(bounds.Dx()) * scaleFactor
			tileY := float64(y) * float64(bounds.Dy()) * scaleFactor
			op.GeoM.Translate(tileX, tileY)
			screen.DrawImage(w.sprite, op)
		}
	}
}
