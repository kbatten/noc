package main

import (
	"github.com/kbatten/httpcanvas"
	"github.com/kbatten/noc"
)

type Mouse struct {
	Mover
}

func newMouse(context *httpcanvas.Context) *Mouse {
	v := &Mouse{*newMover(10, 10, context)}
	v.color = "rgba(0, 175, 0, 0.80)"
	v.size = 10
	return v
}

func (self *Mouse) Step() {
	self.location = noc.NewPVector(self.context.MouseLocation())

	self.limitLocation()
}

func (self *Mouse) limitLocation() {
	if self.location.X > self.context.Width {
		self.location.X = self.context.Width
	}
	if self.location.X < 0 {
		self.location.X = 0
	}
	if self.location.Y > self.context.Height {
		self.location.Y = self.context.Height
	}
	if self.location.Y < 0 {
		self.location.Y = 0
	}
}
