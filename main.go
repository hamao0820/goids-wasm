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

	setting := NewSetting()
	e := goids.CreateEnv(bodyW, bodyH, setting.goidsNum, setting.maxSpeed, setting.maxForce, setting.sight)

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

	canvasEl.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		x := args[0].Get("clientX").Float()
		y := args[0].Get("clientY").Float()
		e.AddGoid(goids.NewGoid(goids.CreateVector(x, y), setting.maxSpeed, setting.maxForce, setting.sight))
		return nil
	}))

	var mouse goids.Vector
	canvasEl.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		x := args[0].Get("clientX").Float()
		y := args[0].Get("clientY").Float()
		mouse = goids.CreateVector(x, y)
		return nil
	}))

	canvasEl.Call("addEventListener", "mouseout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouse = goids.CreateVector(-1, -1)
		return nil
	}))

	var animation js.Func
	animation = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		clearCanvas()
		e.SetHeight(bodyH)
		e.SetWidth(bodyW)
		e.Update(mouse)
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
