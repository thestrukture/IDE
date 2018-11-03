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
// Debugger with struct types.DebugObj
func Debugger(d types.DebugObj) string {
	return NetbDebugger(d)
}

// recovery function used to log a
// panic.
func templateFNDebugger(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/debugger) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDDebugger = "tmpl/ui/debugger.tmpl"

func NetDebugger(args ...interface{}) string {

	localid := templateIDDebugger
	var d *types.DebugObj
	defer templateFNDebugger(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.DebugObj{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Debugger")
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
func bDebugger(d types.DebugObj) string {
	return NetbDebugger(d)
}

//

func NetbDebugger(d types.DebugObj) string {
	localid := templateIDDebugger
	defer templateFNDebugger(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("Debugger")
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
	d = types.DebugObj{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcDebugger(args ...interface{}) (d types.DebugObj) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.DebugObj{}
	}
	return
}

func cDebugger(args ...interface{}) (d types.DebugObj) {
	if len(args) > 0 {
		d = NetcDebugger(args[0])
	} else {
		d = NetcDebugger()
	}
	return
}