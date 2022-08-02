// File generated by Gopher Sauce
// DO NOT EDIT!!
package handlers

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gorilla/sessions"
)

//
func POSTApiGovet(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	pkg := r.FormValue("pkg")
	file := r.FormValue("path")
	path := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg, file)

	cmd := exec.Command("go", "vet", path)
	stOut, _ := cmd.CombinedOutput()

	response = string(stOut)

	callmet = true
	return
}
