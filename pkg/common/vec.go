package common

import (
	"github.com/badrpas/ld50/pkg/config"
	"math"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) XY() (float64, float64) {
	return v.X, v.Y
}

func (v Vec2) XYint() (int, int) {
	return int(v.X), int(v.Y)
}

func (v Vec2) Normalize() Vec2 {
	return v.Scale(1.0 / (v.Length() + math.SmallestNonzeroFloat64))
}

func (v Vec2) Scale(mod float64) Vec2 {
	return Vec2{
		X: v.X * mod,
		Y: v.Y * mod,
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return v.Add(other.Invert())
}

func (v Vec2) LengthSqr() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.LengthSqr())
}

func (v Vec2) Invert() Vec2 {
	return Vec2{
		-v.X,
		-v.Y,
	}
}

func (v Vec2) GridXY() (int, int) {
	x, y := v.Scale(1. / config.CELL_SIZE).XYint()
	return x, y
}
