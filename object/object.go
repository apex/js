//+build js,wasm

// Package object provides functions for working with JavaScript objects.
package object

import "syscall/js"

// object reference.
var object = js.Global().Get("Object")

// Create returns a new object using the given object as its prototype.
func Create(v js.Value) js.Value {
	return object.Call("create", v)
}

// Keys returns the keys of an object.
func Keys(o js.Value) (keys []string) {
	k := object.Call("keys", o)
	l := k.Get("length").Int()
	for i := 0; i < l; i++ {
		keys = append(keys, k.Index(i).String())
	}
	return
}

// Entry is a single key/value pair.
type Entry struct {
	Key string
	js.Value
}

// Entries returns the object entries as key/value pairs.
func Entries(o js.Value) (entries []Entry) {
	e := object.Call("entries", o)
	l := e.Get("length").Int()
	for i := 0; i < l; i++ {
		v := e.Index(i)
		entries = append(entries, Entry{
			Key:   v.Index(0).String(),
			Value: v.Index(1),
		})
	}
	return
}
