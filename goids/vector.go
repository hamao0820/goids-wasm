package goids

import "math"

type Vector struct {
	X float64
	Y float64
}

func CreateVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) ScalarMul(c float64) {
	v.X *= c
	v.Y *= c
}

func (v *Vector) Scale(l float64) {
	if v.Len() == 0 {
		return
	}
	v.ScalarMul(l / v.Len())
}

func (v *Vector) Limit(l float64) {
	if v.Len() > l {
		v.Scale(l)
	}
}

func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector) Sub(v2 Vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func Sub(v2, v1 Vector) Vector {
	return Vector{X: v2.X - v1.X, Y: v2.Y - v1.Y}
}
