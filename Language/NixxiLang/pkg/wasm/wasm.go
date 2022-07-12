package wasm

import "runtime"

func isWasm() bool {
	return runtime.GOOS == "js" && runtime.GOARCH == "wasm"
}
