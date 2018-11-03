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

// Render HTML of template
// EndpointEditor with struct types.TEditor
func EndpointEditor(d types.TEditor) string {
	return NetbEndpointEditor(d)
}

// recovery function used to log a
// panic.
func templateFNEndpointEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/endpoints) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDEndpointEditor = "tmpl/editor/endpoints.tmpl"

func NetEndpointEditor(args ...interface{}) string {

	localid := templateIDEndpointEditor
	var d *types.TEditor
	defer templateFNEndpointEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.TEditor{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointEditor")
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
func bEndpointEditor(d types.TEditor) string {
	return NetbEndpointEditor(d)
}

//

func NetbEndpointEditor(d types.TEditor) string {
	localid := templateIDEndpointEditor
	defer templateFNEndpointEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointEditor")
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
	d = types.TEditor{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcEndpointEditor(args ...interface{}) (d types.TEditor) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.TEditor{}
	}
	return
}

func cEndpointEditor(args ...interface{}) (d types.TEditor) {
	if len(args) > 0 {
		d = NetcEndpointEditor(args[0])
	} else {
		d = NetcEndpointEditor()
	}
	return
}
