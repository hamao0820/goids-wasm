package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"syscall/js"
)

//go:embed images/gopher-front.png
var gopherFront []byte

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

	// img := window.Get("Image").New()
	// img.Set("src", "images/gopher-front.png")
	// img.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	imageWidth := img.Get("width").Float()
	// 	imageHeight := img.Get("height").Float()
	// 	ctx.Call("drawImage", img, bodyW/2-imageWidth/2, bodyH/2-imageHeight/2)
	// 	return nil
	// }))

	src, _, err := image.Decode(bytes.NewReader(gopherFront))
	if err != nil {
		panic(err)
	}

	size := src.Bounds().Size()
	width, height := size.X, size.Y

	canvas := js.Global().Get("Uint8ClampedArray").New(width * height * 4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			canvas.SetIndex((y*width+x)*4+0, float64(r)/255.0)
			canvas.SetIndex((y*width+x)*4+1, float64(g)/255.0)
			canvas.SetIndex((y*width+x)*4+2, float64(b)/255.0)
			canvas.SetIndex((y*width+x)*4+3, float64(a)/255.0)
		}
	}

	imageData := js.Global().Get("ImageData").New(canvas, width, height)
	ctx.Call("putImageData", imageData, bodyW/2-float64(width)/2, bodyH/2-float64(height)/2)

	<-c
}
