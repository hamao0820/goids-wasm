package goids

import "math/rand"

type Environment struct {
	width  float64
	height float64
	goids  []Goid
}

func CreateEnv(width, height float64, n int, maxSpeed, maxForce float64, sight float64) Environment {
	goids := make([]Goid, n)
	for i := range goids {
		position := CreateVector(rand.Float64()*width, rand.Float64()*height)
		goids[i] = NewGoid(position, i, maxSpeed, maxForce, sight)
	}

	return Environment{width: width, height: height, goids: goids}
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
	return len(e.goids)
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

func (e *Environment) AddGoid(g Goid) {
	e.goids = append(e.goids, g)
}
