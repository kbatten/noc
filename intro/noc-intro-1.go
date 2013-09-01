package main

import (
	"github.com/kbatten/httpcanvas"
	"math"
	"math/rand"
)

type Animation interface {
	Display()
	Step()
}

type Walker struct {
	x       float64
	y       float64
	color string
	context *httpcanvas.Context
}

func (self *Walker) Display() {
	self.context.BeginPath()
	self.context.Arc(self.x, self.y, 1, 0, 2*math.Pi, false)
	self.context.FillStyle(self.color)
	self.context.Fill()
}

func (self *Walker) Step() {
	stepX := (2.0 * rand.Float64()) - 1.0
	stepY := (2.0 * rand.Float64()) - 1.0

	self.x += stepX
	self.y += stepY
}

type ForagingWalker struct {
	Walker
}
func (self *ForagingWalker) Step() {
	longX := 1.0
	longY := 1.0
	if rand.Intn(10) == 0 {
		// 10% chance of a long X step
		longX = 5.0
	}
	if rand.Intn(10) == 0 {
		longY = 5.0
	}
    stepX := (2.0 * rand.Float64()) - 1.0
    stepY := (2.0 * rand.Float64()) - 1.0

    self.x += longX * stepX
    self.y += longY * stepY
}

func Animate(al ...Animation) {
	for {
		for _, a := range al {
			a.Step()
			a.Display()
		}
	}
}
func app(context *httpcanvas.Context) {
	centerX := context.Width / 2.0
	centerY := context.Width / 2.0

	walker1 := &Walker{centerX, centerY, "#F00", context}
	walker2 := &Walker{centerX + 50, centerY, "#0F0", context}
	walker3 := &ForagingWalker{Walker{centerX - 50, centerY, "#00F", context}}

	Animate(walker1, walker2, walker3)
}

func main() {
	httpcanvas.ListenAndServe(":8080", app)
}
