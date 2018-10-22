//+build js,wasm

// Package weakmap provides a map of weakly reference key/value pairs
// which are released upon garbage collection.
package weakmap

import "syscall/js"

// WeakMap is a collection of
type WeakMap struct {
	v js.Value
}

// New returns a new WeakMap.
func New() *WeakMap {
	return &WeakMap{
		v: js.Global().Get("WeakMap").New(),
	}
}

// Set a key's value.
func (m *WeakMap) Set(key, value interface{}) {
	m.v.Call("set", key, value)
}

// Get a key's value. Nil is returned if the key does not exist.
func (m *WeakMap) Get(key interface{}) *js.Value {
	v := m.v.Call("get", key)
	if v == js.Undefined() {
		return nil
	}
	return &v
}

// Delete a key's value.
func (m *WeakMap) Delete(key interface{}) {
	m.v.Call("delete", key)
}

// Has returns true if the key is present.
func (m *WeakMap) Has(key interface{}) bool {
	return m.v.Call("has", key).Bool()
}
