// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"log"

	"github.com/thestrukture/IDE/types"
)

// Template path
var templateIDButton = "tmpl/ui/button.tmpl"

//
// Renders HTML of template
// Button with struct types.Dex
func Button(d types.Dex) string {
	return netbButton(d)
}

// Render template with JSON string as
// data.
func netButton(args ...interface{}) string {

	// Get data from JSON
	var d = netcButton(args...)
	return netbButton(d)

}

// template render function
func netbButton(d types.Dex) string {
	localid := templateIDButton
	name := "Button"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcButton(args ...interface{}) (d types.Dex) {

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
