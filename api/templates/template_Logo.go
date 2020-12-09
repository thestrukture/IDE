// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"log"

	"github.com/thestrukture/IDE/types"
)

// Template path
var templateIDLogo = "tmpl/logo.tmpl"

//
// Renders HTML of template
// Logo with struct types.Dex
func Logo(d types.Dex) string {
	return netbLogo(d)
}

// Render template with JSON string as
// data.
func netLogo(args ...interface{}) string {

	// Get data from JSON
	var d = netcLogo(args...)
	return netbLogo(d)

}

// template render function
func netbLogo(d types.Dex) string {
	localid := templateIDLogo
	name := "Logo"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcLogo(args ...interface{}) (d types.Dex) {

	if len(args) > 0 {
		jsonData := args[0].(string)
		err := parseJSON(jsonData, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}

	return
}
