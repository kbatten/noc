package main

import (
	"github.com/kbatten/httpcanvas"
	"github.com/kbatten/noc"

//	"math"

//	"fmt"
)

func app(context *httpcanvas.Context) {
	ani := make([]noc.Animation, 100)
	l := len(ani) - 1
	for i := range ani {
		if i == l {
			ani[i] = newMouse(context)
		} else {
			ani[i] = newVehicle(i, context)
		}
	}
	noc.AnimateFrames(ani...)
}

func main() {
	httpcanvas.ListenAndServe(":8080", app)
}
