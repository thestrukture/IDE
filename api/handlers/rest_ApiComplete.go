package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheikhshift/gos/core"
	"github.com/gorilla/sessions"
	"github.com/thestrukture/IDE/api/globals"
	"gopkg.in/mgo.v2/bson"
)

func ApiComplete(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	prefx := r.FormValue("pref")

	ret := []bson.M{}
	//return {name: ea.word, value: ea.insert, score: 0, meta: ea.meta}
	gxml := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml"

	if _, err := os.Stat(gxml); !os.IsNotExist(err) && r.FormValue("gocode") == "" {
		gos, _ := core.PLoadGos(gxml)
		score := 0
		for _, v := range gos.Variables {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name, "score": score, "meta": "Global variable | " + v.Type})
			}

		}

		for _, v := range gos.RootImports {

			if strings.Contains(v.Src, prefx) {
				score = score + 1
				paths := strings.Split(v.Src, "/")
				ret = append(ret, bson.M{"name": v.Src, "value": paths[len(paths)-1], "score": score, "meta": "package"})
			}

		}

		for _, v := range gos.Header.Structs {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name, "score": score, "meta": "Interface"})
			}

		}

		for _, v := range gos.Header.Objects {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name, "score": score, "meta": "{{Interface func group}}"})
			}

		}

		for _, v := range gos.Methods.Methods {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name + " ", "score": score, "meta": "{{Template pipeline}}"})
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": "Net" + v.Name + "(" + v.Variables + ")", "score": score, "meta": "Pipeline go function."})
			}

		}

		for _, v := range gos.Templates.Templates {

			if strings.Contains(v.Name, prefx) {
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": v.Name + " ", "score": score, "meta": "{{Template $struct}}"})
				score = score + 1
				ret = append(ret, bson.M{"name": v.Name, "value": "Netb" + v.Name + "(" + v.Struct + "{})", "score": score, "meta": "Template build go function."})
			}

		}

		response = mResponse(ret)
	} else {
		content := r.FormValue("content")
		id := r.FormValue("id")
		tempFile := filepath.Join(globals.AutocompletePath, id)

		ioutil.WriteFile(tempFile, []byte(content), 0700)
		cmd := fmt.Sprintf("gocode -f=json --in=%s autocomplete %s", tempFile, prefx)

		res, _ := core.RunCmdSmart(cmd)
		response = res
		os.Remove(tempFile)

	}

	callmet = true
	return
}
