package main

import (
	"github.com/kbatten/httpcanvas"
	"github.com/kbatten/noc"
	"math"
	"math/rand"
)

type Vehicle struct {
	Mover
	id int
}

func newVehicle(id int, context *httpcanvas.Context) *Vehicle {
	x := rand.Float64() * context.Width
	y := rand.Float64() * context.Height

	v := &Vehicle{*newMover(x, y, context),
		id}
	v.acceleration = noc.RandomPVector(0.75, 1.25)
	v.maxSpeed = 4
	v.maxForce = 0.25
	v.color = "rgba(175, 175, 175, 0.75)"
	return v
}

func (self *Vehicle) Step() {
	target := noc.NewPVector(self.context.MouseLocation())
	self.Seek(target)

	self.limitAcceleration()
	self.velocity.Add(self.acceleration)
	self.limitVelocity()
	self.location.Add(self.velocity)
	self.limitLocation()
}

func (self *Vehicle) Display() {
	self.context.BeginPath()
	self.context.Arc(self.location.X, self.location.Y, self.size, 0, 2*math.Pi, false)

	direction := noc.CopyPVector(self.velocity)
	direction.Normalize()
	direction.Mult(self.size)
	front := noc.PVectorAdd(self.location, direction)
	self.context.MoveTo(self.location.X, self.location.Y)
	self.context.LineTo(front.X, front.Y)

	self.context.FillStyle(self.color)
	self.context.StrokeStyle("black")
	self.context.LineWidth(2)
	self.context.Fill()
	self.context.Stroke()
}
