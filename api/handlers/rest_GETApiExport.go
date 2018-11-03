package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cheikhshift/gos/core"
	"github.com/gorilla/sessions"
	"github.com/thestrukture/IDE/api/globals"
	methods "github.com/thestrukture/IDE/api/methods"
)

func GETApiExport(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	os.Chdir(os.ExpandEnv("$GOPATH") + "/src/")
	os.Remove(strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".zip ")

	if _, err := os.Stat("./gos.gxml"); os.IsNotExist(err) {
		core.RunCmdB("go build")
	} else {
		core.RunCmdB("gos --export")
	}

	pkgpath := r.FormValue("pkg")
	zipname := strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".zip"
	if globals.Windows {
		pkgpath = strings.Replace(pkgpath, "/", "\\", -1)
		methods.Zipit(pkgpath, zipname)
	} else {
		core.RunCmdB("zip -r " + zipname + " " + pkgpath + "/")
	}
	time.Sleep(500 * time.Millisecond)

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipname))
	http.ServeFile(w, r, strings.Replace(r.FormValue("pkg"), "/", ".", -1)+".zip")

	callmet = true
	return
}
