package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	baseVelocity = 0.25
)

type Enemy struct {
	sprite   *ebiten.Image
	position Vector
	movement Vector
	count    int
}

func NewEnemy() *Enemy {
	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}
	angle := rand.Float64() * 2 * math.Pi
	r := ScreenWidth / 2.0

	pos := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	velocity := baseVelocity + rand.Float64()*5

	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}
	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}
	sprite := EnemySprite

	return &Enemy{
		position: pos,
		sprite:   sprite,
		movement: movement,
	}
}

func (e *Enemy) Update() {
	e.position.X += e.movement.X
	e.position.Y += e.movement.Y
	e.count++
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}

	transform := math.Sin(float64(e.count/5)) / 5
	opts.GeoM.Scale(SpriteScaleFactor, SpriteScaleFactor+transform)
	opts.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.sprite, opts)
}

func (e *Enemy) Collider() Rect {
	bounds := e.sprite.Bounds()

	return NewRect(
		e.position.X,
		e.position.Y,
		float64(bounds.Dx()*SpriteScaleFactor),
		float64(bounds.Dy()*SpriteScaleFactor),
	)
}
