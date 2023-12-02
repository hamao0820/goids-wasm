package main

import (
	"syscall/js"

	"github.com/hamao0820/goids-wasm/goids"
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

	e := goids.CreateEnv(bodyW, bodyH, 30, 4, 2, 100)

	drawImage := func(x, y float64, t goids.ImageType) {
		img := window.Get("Image").New()
		switch t {
		case goids.Pink:
			img.Set("src", "images/gopher-pink.png")
		case goids.Side:
			img.Set("src", "images/gopher-side.png")
		default:
			img.Set("src", "images/gopher-front.png")
		}
		imageWidth := img.Get("width").Float()
		imageHeight := img.Get("height").Float()
		resizeRatio := float64(goids.GopherSize) / imageHeight
		ctx.Call("drawImage", img, x-imageWidth*resizeRatio/2, y-imageHeight*resizeRatio/2, imageWidth*resizeRatio, imageHeight*resizeRatio)
		img.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return nil
		}))
	}

	var animation js.Func
	animation = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		clearCanvas()
		e.SetHeight(bodyH)
		e.SetWidth(bodyW)
		e.Update()
		for _, goid := range e.Goids() {
			drawImage(goid.Position().X, goid.Position().Y, goid.ImageType())
		}

		window.Call("requestAnimationFrame", animation)
		return nil
	})
	defer animation.Release()

	window.Call("requestAnimationFrame", animation)

	<-c
}
