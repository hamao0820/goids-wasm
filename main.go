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

	offScreenCanvasFront := document.Call("createElement", "canvas")
	offScreenCanvasSide := document.Call("createElement", "canvas")
	offScreenCanvasPink := document.Call("createElement", "canvas")
	imgFront := window.Get("Image").New()
	imgSide := window.Get("Image").New()
	imgPink := window.Get("Image").New()

	imgFront.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		imageWidth := imgFront.Get("width").Float()
		imageHeight := imgFront.Get("height").Float()
		resizeRatio := float64(goids.GopherSize) / imageHeight
		offScreenCanvasFront.Set("width", imageWidth*resizeRatio)
		offScreenCanvasFront.Set("height", imageHeight*resizeRatio)
		offScreenCtxFront := offScreenCanvasFront.Call("getContext", "2d")
		offScreenCtxFront.Call("drawImage", imgFront, 0, 0, imageWidth*resizeRatio, imageHeight*resizeRatio)
		return nil
	}))

	imgSide.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		imageWidth := imgSide.Get("width").Float()
		imageHeight := imgSide.Get("height").Float()
		resizeRatio := float64(goids.GopherSize) / imageHeight
		offScreenCanvasSide.Set("width", imageWidth*resizeRatio)
		offScreenCanvasSide.Set("height", imageHeight*resizeRatio)
		offScreenCtxSide := offScreenCanvasSide.Call("getContext", "2d")
		offScreenCtxSide.Call("drawImage", imgSide, 0, 0, imageWidth*resizeRatio, imageHeight*resizeRatio)
		return nil
	}))

	imgPink.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		imageWidth := imgPink.Get("width").Float()
		imageHeight := imgPink.Get("height").Float()
		resizeRatio := float64(goids.GopherSize) / imageHeight
		offScreenCanvasPink.Set("width", imageWidth*resizeRatio)
		offScreenCanvasPink.Set("height", imageHeight*resizeRatio)
		offScreenCtxPink := offScreenCanvasPink.Call("getContext", "2d")
		offScreenCtxPink.Call("drawImage", imgPink, 0, 0, imageWidth*resizeRatio, imageHeight*resizeRatio)
		return nil
	}))

	imgFront.Set("src", "images/gopher-front.png")
	imgSide.Set("src", "images/gopher-side.png")
	imgPink.Set("src", "images/gopher-pink.png")

	drawImage := func(x, y float64, t goids.ImageType) {
		switch t {
		case goids.Front:
			ctx.Call("drawImage", offScreenCanvasFront, x-offScreenCanvasFront.Get("width").Float()/2, y-offScreenCanvasFront.Get("height").Float()/2)
		case goids.Side:
			ctx.Call("drawImage", offScreenCanvasSide, x-offScreenCanvasSide.Get("width").Float()/2, y-offScreenCanvasSide.Get("height").Float()/2)
		case goids.Pink:
			ctx.Call("drawImage", offScreenCanvasPink, x-offScreenCanvasPink.Get("width").Float()/2, y-offScreenCanvasPink.Get("height").Float()/2)
		}
	}

	canvasEl.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		x := args[0].Get("clientX").Float()
		y := args[0].Get("clientY").Float()
		e.AddGoid(goids.NewGoid(goids.CreateVector(x, y), setting.maxSpeed, setting.maxForce, setting.sight))
		return nil
	}))

	var mouse goids.Vector
	canvasEl.Call("addEventListener", "pointermove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouse.X = args[0].Get("clientX").Float()
		mouse.Y = args[0].Get("clientY").Float()
		return nil
	}))

	canvasEl.Call("addEventListener", "touchmove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouse.X = args[0].Get("touches").Index(0).Get("clientX").Float()
		mouse.Y = args[0].Get("touches").Index(0).Get("clientY").Float()
		return nil
	}))

	canvasEl.Call("addEventListener", "pointerleave", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouse.X = -1
		mouse.Y = -1
		return nil
	}))

	canvasEl.Call("addEventListener", "pointerup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouse.X = -1
		mouse.Y = -1
		return nil
	}))

	canvasEl.Call("addEventListener", "touchend", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouse.X = -1
		mouse.Y = -1
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
