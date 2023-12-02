package goids

type Environment struct {
	width    float64
	height   float64
	goidsNum int
	goids    []Goid
}

func CreateEnv(width, height float64, n int, maxSpeed, maxForce float64, sight float64) Environment {
	goids := make([]Goid, n)
	for i := range goids {
		goids[i] = NewGoid(width, height, i, maxSpeed, maxForce, sight)
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
