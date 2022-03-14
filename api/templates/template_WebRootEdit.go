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
// WebRootEdit with struct types.WebRootEdits
func WebRootEdit(d types.WebRootEdits) string {
	return netbWebRootEdit(d)
}

// recovery function used to log a
// panic.
func templateFNWebRootEdit(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/webrootedit) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDWebRootEdit = "tmpl/ui/user/panel/webrootedit.tmpl"

// Render template with JSON string as
// data.
func netWebRootEdit(args ...interface{}) string {

	localid := templateIDWebRootEdit
	var d *types.WebRootEdits
	defer templateFNWebRootEdit(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.WebRootEdits{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdit")
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
func bWebRootEdit(d types.WebRootEdits) string {
	return netbWebRootEdit(d)
}

//

// template render function
func netbWebRootEdit(d types.WebRootEdits) string {
	localid := templateIDWebRootEdit
	defer templateFNWebRootEdit(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdit")
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
	d = types.WebRootEdits{}
	output.Reset()
	output = nil
	return outpescaped
}

// Unmarshal a json string to the template's struct
// type
func netcWebRootEdit(args ...interface{}) (d types.WebRootEdits) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.WebRootEdits{}
	}
	return
}

// Create a struct variable of template.
func cWebRootEdit(args ...interface{}) (d types.WebRootEdits) {
	if len(args) > 0 {
		d = netcWebRootEdit(args[0])
	} else {
		d = netcWebRootEdit()
	}
	return
}
