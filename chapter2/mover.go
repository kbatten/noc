package main

import (
	"github.com/kbatten/httpcanvas"
	"github.com/kbatten/noc"
	"math"
)

type Mover struct {
	location     *noc.PVector
	velocity     *noc.PVector
	acceleration *noc.PVector
	color        string
	size         float64
	maxSpeed     float64
	maxForce     float64
	context      *httpcanvas.Context
}

func newMover(x, y float64, context *httpcanvas.Context) *Mover {
	return &Mover{noc.NewPVector(x, y),
		noc.NewPVector(0, 0),
		noc.NewPVector(0, 0),
		"rgba(0, 175, 0, 0.75)",
		20,
		0, 0,
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
	self.limitAcceleration()
	self.velocity.Add(self.acceleration)
	self.limitVelocity()
	self.location.Add(self.velocity)
	self.limitLocation()
}

func (self *Mover) limitAcceleration() {
	if self.maxForce > 0 {
		self.acceleration.Limit(self.maxForce)
	}
}

func (self *Mover) limitVelocity() {
	if self.maxSpeed > 0 {
		self.velocity.Limit(self.maxSpeed)
	}
}

func (self *Mover) limitLocation() {
	if self.location.X > self.context.Width || self.location.X < 0 {
		self.location.X = self.context.Width - self.location.X
	}
	if self.location.Y > self.context.Height || self.location.Y < 0 {
		self.location.Y = self.context.Height - self.location.Y
	}
}

func (self *Mover) ClearFrame() {
	self.context.ClearFrame()
}

func (self *Mover) ShowFrame() {
	self.context.ShowFrame()
}

func (self *Mover) ApplyForce(v *noc.PVector) {
	self.acceleration.Add(v)
}

func (self *Mover) Seek(target *noc.PVector) {
	if noc.PVectorCmp(target, self.location) {
		return
	}
	desired := noc.PVectorSub(target, self.location)
	if self.maxSpeed > 0 {
		desired.Normalize()
		desired.Mult(self.maxSpeed)
	}

	steer := noc.PVectorSub(desired, self.velocity)
	if self.maxForce > 0 {
		steer.Limit(self.maxForce)
	}
	self.ApplyForce(steer)
}
