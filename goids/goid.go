package goids

import (
	"math"
	"math/rand"
)

type ImageType int

const (
	Front ImageType = iota
	Side
	Pink
)

const GopherSize = 32

type Goid struct {
	position     Vector
	velocity     Vector
	acceleration Vector
	maxSpeed     float64
	maxForce     float64
	sight        float64
	imageType    ImageType
}

func NewGoid(width, height float64, n int, maxSpeed, maxForce float64, sight float64) Goid {
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

	return Goid{position: position, velocity: velocity, maxSpeed: float64(maxSpeed), maxForce: float64(maxForce), sight: sight, imageType: t}
}

func (g *Goid) Seek(t Vector) {
	tv := Sub(t, g.position)
	tv.Limit(g.maxSpeed)
	force := Sub(tv, g.velocity)
	g.acceleration.Add(force)
}

func (g *Goid) Flee(t Vector) {
	tv := Sub(t, g.position)
	tv.Limit(g.maxSpeed)
	force := Sub(tv, g.velocity)
	g.acceleration.Sub(force)
}

func (g Goid) IsInsight(g2 Goid) bool {
	d := Sub(g.position, g2.position).Len()
	return d < g.sight
}

func (g *Goid) Align(goids []Goid) {
	var avgVel Vector
	n := 0
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		avgVel.Add(other.velocity)
		n++
	}
	if n > 0 {
		avgVel.ScalarMul(1 / float64(n))
		avgVel.Limit(g.maxSpeed)
		g.acceleration.Add(Sub(avgVel, g.velocity))
	}
}

func (g *Goid) Separate(goids []Goid) {
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		d := Sub(g.position, other.position).Len()
		if d < 50 {
			g.Flee(other.position)
		}
	}
}

func (g *Goid) Cohesive(goids []Goid) {
	var avgPos Vector
	n := 0
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		avgPos.Add(other.position)
		n++
	}
	if n > 0 {
		avgPos.ScalarMul(1 / float64(n))
		g.Seek(avgPos)
	}
}

func (g *Goid) Flock(goids []Goid) {
	g.Align(goids)
	g.Separate(goids)
	g.Cohesive(goids)
}

func (g *Goid) AdjustEdge(width, height float64) {
	if g.position.X < float64(GopherSize)/2 {
		g.position.X = float64(GopherSize) / 2
		g.velocity.X = math.Abs(g.velocity.X)
	} else if g.position.X >= width-float64(GopherSize)/2 {
		g.position.X = width - float64(GopherSize)/2 - 1
		g.velocity.X = -math.Abs(g.velocity.X)
	}

	if g.position.Y < float64(GopherSize)/2 {
		g.position.Y = float64(GopherSize) / 2
		g.velocity.Y = math.Abs(g.velocity.Y)
	} else if g.position.Y >= height-float64(GopherSize)/2 {
		g.position.Y = height - float64(GopherSize)/2 - 1
		g.velocity.Y = -math.Abs(g.velocity.Y)
	}
}

func (g *Goid) Update(width, height float64) {
	g.acceleration.Limit(g.maxForce)
	g.velocity.Add(g.acceleration)
	g.velocity.Limit(g.maxSpeed)
	g.position.Add(g.velocity)
	g.acceleration.ScalarMul(0)

	g.AdjustEdge(width, height)
}

func (g Goid) Position() Vector {
	return g.position
}

func (g Goid) ImageType() ImageType {
	return g.imageType
}
