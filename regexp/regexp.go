//+build js,wasm

// Package regexp provides RegExp bindings.
package regexp

import (
	"syscall/js"
)

// references.
var regexp = js.Global().Get("RegExp")

// Flag enum.
type Flag int

// Flags available.
const (
	// Global finds all matches rather than stopping after the first match.
	Global Flag = 1 << iota

	// IgnoreCase ignore casing; if u flag is also enabled, use Unicode case folding.
	IgnoreCase

	// Multiline; treat beginning and end characters (^ and $) as working over multiple
	// lines (i.e., match the beginning or end of each line (delimited by \n or \r), not
	// only the very beginning or end of the whole input string)
	Multiline

	// Unicode; treat pattern as a sequence of Unicode code points.
	Unicode

	// Sticky; matches only from the index indicated by the lastIndex property of this regular
	// expression in the target string (and does not attempt to match from any later indexes).
	Sticky
)

// String returns the flag(s) specified in the single-letter format
// expected by the RegExp constructor.
func (f Flag) String() (s string) {
	if f&Global == Global {
		s += "g"
	}

	if f&IgnoreCase == IgnoreCase {
		s += "i"
	}

	if f&Multiline == Multiline {
		s += "m"
	}

	if f&Unicode == Unicode {
		s += "u"
	}

	if f&Sticky == Sticky {
		s += "y"
	}

	return
}

// Result is the match information.
type Result struct {
	v js.Value
}

// Input returns the original input string.
func (r *Result) Input() string {
	return r.v.Get("input").String()
}

// Index returns the index of the matched string.
func (r *Result) Index() int {
	return r.v.Get("index").Int()
}

// String returns the full matched string.
func (r *Result) String() string {
	return r.v.Index(0).String()
}

// Match returns the parenthesized substring match at index.
func (r *Result) Match(i int) string {
	return r.v.Index(i + 1).String()
}

// RegExp is a regular expression.
type RegExp struct {
	v js.Value
}

// New returns a regexp from the given pattern and flags.
func New(pattern string, flags Flag) *RegExp {
	return &RegExp{
		v: regexp.New(pattern, flags.String()),
	}
}

// Exec returns the result of a search execution match, or nil.
func (r *RegExp) Exec(s string) *Result {
	v := r.v.Call("exec", s)

	if v == js.Null() {
		return nil
	}

	return &Result{v}
}

// String implementation.
func (r *RegExp) String() string {
	return r.v.String()
}
