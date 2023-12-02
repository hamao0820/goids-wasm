package main

import (
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
		bodyW := window.Get("innerWidth").Float()
		bodyH := window.Get("innerHeight").Float()
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

	canvasEl.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		clearCanvas()
		x := args[0].Get("clientX").Float()
		y := args[0].Get("clientY").Float()
		renderTriangle(x, y)
		return nil
	}))

	<-c
}
