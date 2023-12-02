package main

import (
	"math"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	window := js.Global()
	document := window.Get("document")
	canvasEl := document.Call("getElementById", "canvas")

	bodyW := window.Get("innerWidth").Float()
	bodyH := window.Get("innerHeight").Float()
	canvasEl.Set("width", bodyW)
	canvasEl.Set("height", bodyH)
	window.Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		bodyW = window.Get("innerWidth").Float()
		bodyH = window.Get("innerHeight").Float()
		canvasEl.Set("width", bodyW)
		canvasEl.Set("height", bodyH)
		return nil
	}))

	renderTriangle := func(x, y float64) {
		ctx := canvasEl.Call("getContext", "2d")
		ctx.Call("beginPath")
		ctx.Call("moveTo", x, y)
		ctx.Call("lineTo", x+50, y+25)
		ctx.Call("lineTo", x+50, y-25)
		ctx.Set("fillStyle", "#ff0000")
		ctx.Call("fill")
		ctx.Call("closePath")
	}

	clearCanvas := func() {
		ctx := canvasEl.Call("getContext", "2d")

		ctx.Call("clearRect", 0, 0, bodyW, bodyH)
	}

	t := 0.0
	var animation js.Func
	animation = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		t += 1
		t = math.Mod(t, 360)
		clearCanvas()
		v := math.Sqrt(bodyW*bodyW + 4*bodyH*bodyH) / 4
		x := bodyW/4*math.Sin(t*math.Pi/180)*500/v + bodyW/2
		y := bodyH/4*math.Sin(2*t*math.Pi/180)*500/v + bodyH/2
		renderTriangle(x, y)
		window.Call("requestAnimationFrame", animation)
		return nil
	})
	defer animation.Release()

	window.Call("requestAnimationFrame", animation)

	<-c
}
