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
// TemplateEditTwo with struct types.TemplateEdits
func TemplateEditTwo(d types.TemplateEdits) string {
	return NetbTemplateEditTwo(d)
}

// recovery function used to log a
// panic.
func templateFNTemplateEditTwo(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/tpetwo) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDTemplateEditTwo = "tmpl/ui/user/panel/tpetwo.tmpl"

func NetTemplateEditTwo(args ...interface{}) string {

	localid := templateIDTemplateEditTwo
	var d *types.TemplateEdits
	defer templateFNTemplateEditTwo(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.TemplateEdits{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TemplateEditTwo")
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
func bTemplateEditTwo(d types.TemplateEdits) string {
	return NetbTemplateEditTwo(d)
}

//

func NetbTemplateEditTwo(d types.TemplateEdits) string {
	localid := templateIDTemplateEditTwo
	defer templateFNTemplateEditTwo(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("TemplateEditTwo")
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
	d = types.TemplateEdits{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcTemplateEditTwo(args ...interface{}) (d types.TemplateEdits) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.TemplateEdits{}
	}
	return
}

func cTemplateEditTwo(args ...interface{}) (d types.TemplateEdits) {
	if len(args) > 0 {
		d = NetcTemplateEditTwo(args[0])
	} else {
		d = NetcTemplateEditTwo()
	}
	return
}
