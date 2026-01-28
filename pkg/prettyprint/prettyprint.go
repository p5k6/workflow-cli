// Package prettyprint contains tools for formatting text.
// This is a minimal version inlined from github.com/teamhephy/pkg/prettyprint
package prettyprint

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

// Colors contains a map of the standard ANSI color codes.
var Colors = map[string]string{
	"Default": "\033[0m",
	"Black":   "\033[0;30m",
	"Red":     "\033[0;31m",
	"Green":   "\033[0;32m",
	"Yellow":  "\033[0;33m",
	"Blue":    "\033[0;34m",
	"Purple":  "\033[0;35m",
	"Cyan":    "\033[0;36m",
	"White":   "\033[0;37m",
}

// ColorizeVars provides template rendering with color support.
//
// The template is given a datum with two objects: `.V` and `.C`. `.V` contains
// the `vars` passed into the function. `.C` contains the color map.
func ColorizeVars(msg string, vars interface{}) string {
	var t = struct {
		V interface{}
		C map[string]string
	}{
		V: vars,
		C: Colors,
	}

	tpl, err := template.New(msg).Parse(msg)
	if err != nil {
		return msg
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, t); err != nil {
		return msg
	}

	return buf.String()
}

// PrettyTabs formats a map with aligned keys and values.
func PrettyTabs(msg map[string]string, spaces int) string {
	// find the longest key so we know how much padding to use
	max := 0
	for key := range msg {
		if len(key) > max {
			max = len(key)
		}
	}
	max += spaces

	// sort the map keys so we can print them alphabetically
	var keys []string
	for k := range msg {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var output string
	for _, k := range keys {
		output += fmt.Sprintf("%s%s%s\n", k, strings.Repeat(" ", max-len(k)), msg[k])
	}
	return output
}
