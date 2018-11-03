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
// WebRootEdittwo with struct types.WebRootEdits
func WebRootEdittwo(d types.WebRootEdits) string {
	return NetbWebRootEdittwo(d)
}

// recovery function used to log a
// panic.
func templateFNWebRootEdittwo(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/user/panel/webtwo) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDWebRootEdittwo = "tmpl/ui/user/panel/webtwo.tmpl"

func NetWebRootEdittwo(args ...interface{}) string {

	localid := templateIDWebRootEdittwo
	var d *types.WebRootEdits
	defer templateFNWebRootEdittwo(localid, d)
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
		var localtemplate = template.New("WebRootEdittwo")
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
func bWebRootEdittwo(d types.WebRootEdits) string {
	return NetbWebRootEdittwo(d)
}

//

func NetbWebRootEdittwo(d types.WebRootEdits) string {
	localid := templateIDWebRootEdittwo
	defer templateFNWebRootEdittwo(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("WebRootEdittwo")
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
func NetcWebRootEdittwo(args ...interface{}) (d types.WebRootEdits) {
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

func cWebRootEdittwo(args ...interface{}) (d types.WebRootEdits) {
	if len(args) > 0 {
		d = NetcWebRootEdittwo(args[0])
	} else {
		d = NetcWebRootEdittwo()
	}
	return
}
