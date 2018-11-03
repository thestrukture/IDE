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
// TimerEditor with struct types.TEditor
func TimerEditor(d types.TEditor) string {
	return NetbTimerEditor(d)
}

// recovery function used to log a
// panic.
func templateFNTimerEditor(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (editor/timers) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDTimerEditor = "tmpl/editor/timers.tmpl"

func NetTimerEditor(args ...interface{}) string {

	localid := templateIDTimerEditor
	var d *types.TEditor
	defer templateFNTimerEditor(localid, d)
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
		var localtemplate = template.New("TimerEditor")
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
func bTimerEditor(d types.TEditor) string {
	return NetbTimerEditor(d)
}

//

func NetbTimerEditor(d types.TEditor) string {
	localid := templateIDTimerEditor
	defer templateFNTimerEditor(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TimerEditor")
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
func NetcTimerEditor(args ...interface{}) (d types.TEditor) {
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

func cTimerEditor(args ...interface{}) (d types.TEditor) {
	if len(args) > 0 {
		d = NetcTimerEditor(args[0])
	} else {
		d = NetcTimerEditor()
	}
	return
}
