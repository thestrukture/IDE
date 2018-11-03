package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"

	gosweb "github.com/cheikhshift/gos/web"
	"github.com/fatih/color"
	"github.com/thestrukture/IDE/api/assets"
)

// Render HTML of template
// PluginList with struct gosweb.NoStruct
func PluginList(d gosweb.NoStruct) string {
	return NetbPluginList(d)
}

// recovery function used to log a
// panic.
func templateFNPluginList(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/pluginlist) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDPluginList = "tmpl/ui/pluginlist.tmpl"

func NetPluginList(args ...interface{}) string {

	localid := templateIDPluginList
	var d *gosweb.NoStruct
	defer templateFNPluginList(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &gosweb.NoStruct{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PluginList")
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
func bPluginList(d gosweb.NoStruct) string {
	return NetbPluginList(d)
}

//

func NetbPluginList(d gosweb.NoStruct) string {
	localid := templateIDPluginList
	defer templateFNPluginList(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PluginList")
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
	d = gosweb.NoStruct{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcPluginList(args ...interface{}) (d gosweb.NoStruct) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = gosweb.NoStruct{}
	}
	return
}

func cPluginList(args ...interface{}) (d gosweb.NoStruct) {
	if len(args) > 0 {
		d = NetcPluginList(args[0])
	} else {
		d = NetcPluginList()
	}
	return
}
