//+build js,wasm

// Package console provides console logging API access.
package console

import "syscall/js"

// Log outputs a message to the browser console.
func Log(args ...interface{}) {
	js.Global().Get("console").Call("log", args...)
}

// Dir outputs an interactive list of the properties to the browser console.
func Dir(args ...interface{}) {
	js.Global().Get("console").Call("dir", args...)
}
