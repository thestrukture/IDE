package methods

import "strings"

//
func Fragmentize(args ...interface{}) (finall string) {
	inn := args[0]

	finall = strings.Replace(inn.(string), ".tmpl", "", -1)
	return

}
