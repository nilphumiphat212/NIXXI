package wasm

import "testing"

func TestIsWasm(t *testing.T) {
	result := isWasm()
	if result {
		t.Error("Can not detect WASM mode")
	}
}
