// File generated by Gopher Sauce
// DO NOT EDIT!!
package methods

import (
	"reflect"

	"github.com/thestrukture/IDE/types"
)

//
func UpdateApp(args ...interface{}) []types.App {
	apps := args[0]
	name := args[1]
	app := args[2]

	s := reflect.ValueOf(apps)
	n := make([]types.App, s.Len())
	slice := make([]types.App, s.Len())
	for i, _ := range slice {
		v := s.Index(i).Interface().(types.App)

		if v.Name == name.(string) {
			n = append(n, app.(types.App))
		} else if v.Name != "" {
			n = append(n, v)
		}
	}
	return n

}
