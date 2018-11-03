package methods

import "github.com/thestrukture/IDE/types"

//
func BindID(args ...interface{}) types.Dex {
	id := args[0]
	nav := args[1]

	Nav := nav.(types.Dex)
	Nav.Misc = id.(string)
	return Nav

}
