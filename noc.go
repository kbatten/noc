package noc

import (
	"math"
	"math/rand"
	"time"
)

type PVector struct {
	X float64
	Y float64
}

func NewPVector(x, y float64) *PVector {
	return &PVector{x, y}
}

func RandomPVector(min, max float64) *PVector {
	v := PVector{2.0*rand.Float64() - 1.0,
		2.0*rand.Float64() - 1.0}
	v.Extend(min)
	v.Limit(max)
	return &v
}

func CopyPVector(p *PVector) *PVector {
	return NewPVector(p.X, p.Y)
}

func PVectorCmp(a, b *PVector) bool {
	return a.X == b.X && a.Y == b.Y
}

func (self *PVector) Add(other *PVector) {
	self.X += other.X
	self.Y += other.Y
}

func PVectorAdd(a, b *PVector) *PVector {
    return NewPVector(a.X + b.X, a.Y + b.Y)
}

func PVectorSub(a, b *PVector) *PVector {
	return NewPVector(a.X - b.X, a.Y - b.Y)
}

func (self *PVector) Extend(min float64) {
	length := self.Length()
	if length >= min {
		return
	}
	scale := min / length
	self.Mult(scale)
}

func (self *PVector) Limit(max float64) {
	length := self.Length()
	if length <= max {
		return
	}
	scale := max / length
	self.Mult(scale)
}

func (self *PVector) Length() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y)
}

// normalize the vector to a unit length of 1
func (self *PVector) Normalize() {
	length := self.Length()
	scale := 1.0 / length
	self.Mult(scale)
}

// scale the vector with multiplication
func (self *PVector) Mult(f float64) {
	self.X *= f
	self.Y *= f
}

type Animation interface {
	Display()
	Step()
	ClearFrame()
	ShowFrame()
}

func AnimateFrames(al ...Animation) {
	for _ = range time.Tick(30 * time.Millisecond) {
        al[0].ClearFrame()
		for _, a := range al {
			a.Step()
			a.Display()
		}
		al[0].ShowFrame()
	}
}
