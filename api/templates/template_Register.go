// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"

	"github.com/fatih/color"
	"github.com/thestrukture/IDE/api/assets"
	"github.com/thestrukture/IDE/types"
)

//
// Renders HTML of template
// Register with struct types.Dex
func Register(d types.Dex) string {
	return netbRegister(d)
}

// recovery function used to log a
// panic.
func templateFNRegister(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/register) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDRegister = "tmpl/ui/register.tmpl"

// Render template with JSON string as
// data.
func netRegister(args ...interface{}) string {

	localid := templateIDRegister
	var d *types.Dex
	defer templateFNRegister(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Register")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}

// alias of template render function.
func bRegister(d types.Dex) string {
	return netbRegister(d)
}

//

// template render function
func netbRegister(d types.Dex) string {
	localid := templateIDRegister
	defer templateFNRegister(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Register")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = types.Dex{}
	output.Reset()
	output = nil
	return outpescaped
}

// Unmarshal a json string to the template's struct
// type
func netcRegister(args ...interface{}) (d types.Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.Dex{}
	}
	return
}

// Create a struct variable of template.
func cRegister(args ...interface{}) (d types.Dex) {
	if len(args) > 0 {
		d = netcRegister(args[0])
	} else {
		d = netcRegister()
	}
	return
}
