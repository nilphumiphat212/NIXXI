package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func main() {
	var jsChan chan struct{} = make(chan struct{}, 0)
	fmt.Println("Hello WASM")
	js.Global().Set("checkOs", js.FuncOf(checkOs))
	<-jsChan
}

func checkOs(this js.Value, args []js.Value) interface{} {
	fmt.Println(runtime.GOOS)
	return 0
}
