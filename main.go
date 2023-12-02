package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	window := js.Global()
	document := window.Get("document")
	canvasEl := document.Call("getElementById", "canvas")

	bodyW := window.Get("innerWidth").Float()
	bodyH := window.Get("innerHeight").Float()
	fmt.Println(bodyW, bodyH)
	canvasEl.Set("width", bodyW)
	canvasEl.Set("height", bodyH)
	window.Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		bodyW := window.Get("innerWidth").Float()
		bodyH := window.Get("innerHeight").Float()
		fmt.Println(bodyW, bodyH)
		canvasEl.Set("width", bodyW)
		canvasEl.Set("height", bodyH)
		return nil
	}))

	<-c
}
