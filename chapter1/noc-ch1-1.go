package main

import (
	"github.com/kbatten/httpcanvas"
	"github.com/kbatten/noc"
	"math"

//	"fmt"
)

type Mover struct {
	location     *noc.PVector
	velocity     *noc.PVector
	acceleration *noc.PVector
	color        string
	size         float64
	context      *httpcanvas.Context
}

func newMover(x, y float64, context *httpcanvas.Context) *Mover {
	return &Mover{noc.NewPVector(x, y),
		noc.NewPVector(0, 0),
		noc.NewPVector(0, 0),
		"rgba(0, 175, 0, 0.75)",
		20,
		context}
}

func (self *Mover) Display() {
	self.context.BeginPath()
	self.context.Arc(self.location.X, self.location.Y, self.size, 0, 2*math.Pi, false)
	self.context.FillStyle(self.color)
	self.context.StrokeStyle("black")
	self.context.LineWidth(2)
	self.context.Fill()
	self.context.Stroke()
}

func (self *Mover) Step() {
	self.acceleration = noc.RandomPVector(0.75, 1.25)
	self.velocity.Add(self.acceleration)
	self.velocity.Limit(4)
	self.location.Add(self.velocity)

	if self.location.X > self.context.Width || self.location.X < 0 {
		self.location.X = self.context.Width - self.location.X
		//		self.velocity.X = -self.velocity.X
	}
	if self.location.Y > self.context.Height || self.location.Y < 0 {
		self.location.Y = self.context.Height - self.location.Y
		//		self.velocity.Y = -self.velocity.Y
	}
}

func (self *Mover) NewFrame() {
	self.context.NewFrame()
}

func (self *Mover) ShowFrame() {
	self.context.ShowFrame()
}

func app(context *httpcanvas.Context) {
	ani := make([]noc.Animation, 100)
	for i := range ani {
		ani[i] = newMover(100, 100, context)
	}

	noc.AnimateFrames(ani...)
}

func main() {
	httpcanvas.ListenAndServe(":8080", app)
}
