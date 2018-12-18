//+build js,wasm

// Package lit provides lit-html bindings.
package lit

import (
	"syscall/js"

	"github.com/apex/js/regexp"
)

// TODO: bundle lit-html?
// TODO: implement caching in a less hacky way haha...

// references.
var (
	lit      = js.Global().Get("lit")
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

// cache entry.
type cache struct {
	strings    []interface{}
	argIndexes []int
}

// caches is a set of cache entries.
var caches = map[string]cache{}

// parse returns the sub-strings and arguments for use
// in the calls to lit-html's html() and svg() functions,
// letting us use $N to reference positional arguments.
func parse(tmpl string, values []interface{}) (strings []interface{}, args []interface{}) {
primed:
	c, ok := caches[tmpl]
	if ok {
		strings = c.strings
		for _, i := range c.argIndexes {
			args = append(args, values[i])
		}
		return
	}

	re := regexp.New(`\$(\d+)`, regexp.Global)
	var prevIndex int

	for {
		m := re.Exec(tmpl)
		if m == nil {
			break
		}

		// matched string
		match := m.String()

		// positional arg
		argIndex := parseInt.Invoke(m.Match(0)).Int()
		c.argIndexes = append(c.argIndexes, argIndex)

		// substring
		index := m.Index()
		str := tmpl[prevIndex:index]
		c.strings = append(c.strings, str)

		prevIndex = index + len(match)
	}

	// substring
	str := tmpl[prevIndex:]
	c.strings = append(c.strings, str)

	// cache
	caches[tmpl] = c
	goto primed
}
