package goids

import (
	"math/rand"
)

type Environment struct {
	width    float64
	height   float64
	goidsNum int
	goids    []Goid
}

func CreateEnv(width, height float64, n int, maxSpeed, maxForce float64, sight float64) Environment {
	goids := make([]Goid, n)
	for i := range goids {
		position := CreateVector(rand.Float64()*width, rand.Float64()*height)
		velocity := CreateVector(rand.Float64()*2-1, rand.Float64()*2-1)
		velocity.Scale(rand.Float64()*4 - rand.Float64()*2)

		var t ImageType

		r := rand.Float64()

		if r < 0.001 { // 0.1%
			t = Pink
		} else if r < 0.011 { // 1%
			t = Side
		} else {
			t = Front
		}

		goids[i] = Goid{position: position, velocity: velocity, maxSpeed: float64(maxSpeed), maxForce: float64(maxForce), sight: sight, imageType: t}
	}

	return Environment{width: width, height: height, goidsNum: n, goids: goids}
}

func (e *Environment) Update() {
	for i := 0; i < len(e.goids); i++ {
		goid := &e.goids[i]
		goid.Flock(e.goids)
		goid.Update(e.width, e.height)
	}
}

func (e Environment) Goids() []Goid {
	return e.goids
}

func (e Environment) GoidsNum() int {
	return e.goidsNum
}

func (e Environment) Width() float64 {
	return e.width
}

func (e Environment) Height() float64 {
	return e.height
}

func (e *Environment) SetWidth(width float64) {
	e.width = width
}

func (e *Environment) SetHeight(height float64) {
	e.height = height
}
