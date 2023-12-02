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
		bodyW = window.Get("innerWidth").Float()
		bodyH = window.Get("innerHeight").Float()
		canvasEl.Set("width", bodyW)
		canvasEl.Set("height", bodyH)
		return nil
	}))

	ctx := canvasEl.Call("getContext", "2d")

	image := window.Get("Image").New()
	image.Set("src", "images/gopher-front.png")
	image.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		imageWidth := image.Get("width").Float()
		imageHeight := image.Get("height").Float()
		ctx.Call("drawImage", image, bodyW/2-imageWidth/2, bodyH/2-imageHeight/2)
		return nil
	}))

	<-c
}
