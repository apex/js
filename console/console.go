//+build js,wasm

// Package console provides console logging API access.
package console

import "syscall/js"

// console reference.
var console = js.Global().Get("console")

// Log outputs a message to the browser console.
func Log(args ...interface{}) {
	console.Call("log", args...)
}

// Dir outputs an interactive list of the properties to the browser console.
func Dir(args ...interface{}) {
	console.Call("dir", args...)
}
