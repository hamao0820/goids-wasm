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

	clearCanvas := func() {
		ctx := canvasEl.Call("getContext", "2d")
		ctx.Call("clearRect", 0, 0, bodyW, bodyH)
	}

	ctx := canvasEl.Call("getContext", "2d")

	drawImage := func(x, y float64) {
		img := window.Get("Image").New()
		img.Set("src", "images/gopher-front.png")
		imageWidth := img.Get("width").Float()
		imageHeight := img.Get("height").Float()
		ctx.Call("drawImage", img, x-imageWidth/2, y-imageHeight/2)
		img.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return nil
		}))
	}

	t := 0.0
	var animation js.Func
	animation = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		t += 1
		t = math.Mod(t, 360)
		clearCanvas()
		x := bodyW/4*math.Sin(t*math.Pi/180) + bodyW/2
		y := bodyH/4*math.Sin(2*t*math.Pi/180) + bodyH/2
		drawImage(x, y)
		window.Call("requestAnimationFrame", animation)
		return nil
	})
	defer animation.Release()

	window.Call("requestAnimationFrame", animation)

	<-c
}
