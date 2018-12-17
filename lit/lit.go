//+build js,wasm

// Package lit provides lit-html bindings.
package lit

import (
	"syscall/js"
)

// TODO: bundle lit-html?

// references.
var (
	lit      = js.Global().Get("lit")
	regexp   = js.Global().Get("RegExp")
	parseInt = js.Global().Get("parseInt")
)

// Template is a lit-html template result.
type Template = js.Value

// HTML returns a template.
func HTML(tmpl string, values ...interface{}) Template {
	return render("html", tmpl, values...)
}

// SVG returns a template.
func SVG(tmpl string, values ...interface{}) Template {
	return render("svg", tmpl, values...)
}

// Render a template.
func Render(t Template, container js.Value) {
	lit.Call("render", t, container)
}

// render template.
func render(kind, tmpl string, values ...interface{}) Template {
	var args []interface{}
	strings, values := parse(tmpl, values)
	args = append(args, strings)
	args = append(args, values...)
	res := lit.Call(kind, args...)
	return Template(res)
}

// parse template.
func parse(tmpl string, values []interface{}) (strings []interface{}, args []interface{}) {
	re := regexp.New(`\$(\d+)`, "g")
	var prevIndex int

	for {
		m := re.Call("exec", tmpl)

		if m == js.Null() {
			break
		}

		// matched string
		match := m.Index(0).String()

		// positional arg
		argIndex := parseInt.Invoke(m.Index(1)).Int()
		arg := values[argIndex]
		args = append(args, arg)

		// substring
		index := m.Get("index").Int()
		str := tmpl[prevIndex:index]
		strings = append(strings, str)

		prevIndex = index + len(match)
	}

	// substring
	str := tmpl[prevIndex:]
	strings = append(strings, str)

	return
}
