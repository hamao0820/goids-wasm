package main

import (
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")
	canvasEl := js.Global().Get("document").Call("getElementById", "canvas")

	bodyW := doc.Get("body").Get("clientWidth").Float()
	bodyH := doc.Get("body").Get("clientHeight").Float()
	canvasEl.Set("width", bodyW)
	canvasEl.Set("height", bodyH)
}
