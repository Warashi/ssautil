package ssautil

import (
	"go/ast"
	"regexp"
)

var GeneratedRegexp = regexp.MustCompile(`^Code generated .* DO NOT EDIT\.$`)

func IsGenerated(f *ast.File) bool {
	for _, c := range f.Comments {
		if GeneratedRegexp.MatchString(c.Text()) {
			return true
		}
	}
	return false
}
