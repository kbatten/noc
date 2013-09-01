package main

import (
	"github.com/kbatten/httpcanvas"
	"math/rand"
)

type Animation interface {
	Display()
	Step()
}

type RandDist struct {
	context *httpcanvas.Context
}

func (self *RandDist) Display() {
		index := rand.Intn(len(randomCounts))
		randomCounts[index]++

		self.context.FillStyle("gray")
		self.context.StrokeStyle("black")

		w := self.context.Width / float64(len(randomCounts))
		for x := range randomCounts {
			x1 := float64(x) * w
			y1 := self.context.Height - float64(randomCounts[x])
			h1 := w - 1
			w1 := float64(randomCounts[x])
			self.context.FillRect(x1, y1, h1, w1)
			self.context.StrokeRect(x1, y1, h1, w1)
		}
}

func (self *RandDist) Step() {
}

var randomCounts = [20]int{}

func Animate(al ...Animation) {
	for {
		for _, a := range al {
			a.Step()
			a.Display()
		}
	}
}

func app(context *httpcanvas.Context) {
	ani := []*RandDist{
		&RandDist{context},
	}
	Animate(ani[0])
}

func main() {
	httpcanvas.ListenAndServe(":8080", app)
}
