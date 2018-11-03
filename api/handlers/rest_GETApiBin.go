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

func GETApiBin(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	os.Chdir(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg"))
	os.Remove(strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".binary.zip ")
	bPath := strings.Split(r.FormValue("pkg"), "/")
	gp := os.ExpandEnv("$GOPATH")
	//coreTemplate,_ := core.LoadGos(gp + "/src/" +  r.FormValue("pkg") + "/gos.gxml")

	//core.RunCmdB("gos ru " +  r.FormValue("pkg") + " gos.gxml web tmpl")
	os.Remove(gp + "/src/" + r.FormValue("pkg") + "/server_out.go")

	//core.Process(coreTemplate,gp + "/src/" +  r.FormValue("pkg"), "web","tmpl")
	if _, err := os.Stat("./gos.gxml"); os.IsNotExist(err) {
		core.RunCmdB("go build")
	} else {
		core.RunCmdB("gos --export")
	}

	zipname := strings.Replace(r.FormValue("pkg"), "/", ".", -1) + ".binary.zip"
	if globals.Windows {
		bPath[len(bPath)-1] += ".exe"
		methods.Zipit(bPath[len(bPath)-1], zipname)
	} else {
		core.RunCmdB("zip -r " + zipname + " " + bPath[len(bPath)-1])
	}
	time.Sleep(500 * time.Millisecond)

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipname))
	http.ServeFile(w, r, zipname)

	callmet = true
	return
}
