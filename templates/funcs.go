// Helper functions to add to templates via Template.Funcs()
//
// We can then use these in templates like (for the add function)
//     {{ add 1 2 }}

package templates

import (
	"html/template"
	"strings"
)

var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"split": strings.Split,
	"trimsuffix": strings.TrimSuffix,
}
