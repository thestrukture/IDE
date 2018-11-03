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
// ObjectEditor with struct types.VHuf
func ObjectEditor(d types.VHuf) string {
	return NetbObjectEditor(d)
}

// recovery function used to log a
// panic.
func templateFNObjectEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/objects) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDObjectEditor = "tmpl/editor/objects.tmpl"

func NetObjectEditor(args ...interface{}) string {

	localid := templateIDObjectEditor
	var d *types.VHuf
	defer templateFNObjectEditor(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.VHuf{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("ObjectEditor")
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
func bObjectEditor(d types.VHuf) string {
	return NetbObjectEditor(d)
}

//

func NetbObjectEditor(d types.VHuf) string {
	localid := templateIDObjectEditor
	defer templateFNObjectEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("ObjectEditor")
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
	d = types.VHuf{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcObjectEditor(args ...interface{}) (d types.VHuf) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.VHuf{}
	}
	return
}

func cObjectEditor(args ...interface{}) (d types.VHuf) {
	if len(args) > 0 {
		d = NetcObjectEditor(args[0])
	} else {
		d = NetcObjectEditor()
	}
	return
}
