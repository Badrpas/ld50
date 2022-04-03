package common

type Vec2 struct {
	X, Y float64
}

func (v Vec2) XY() (float64, float64) {
	return v.X, v.Y
}

func (v Vec2) XYint() (int, int) {
	return int(v.X), int(v.Y)
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
