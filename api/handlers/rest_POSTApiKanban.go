package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	methods "github.com/thestrukture/IDE/api/methods"
)

func POSTApiKanban(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	pkgName := r.FormValue("pkg")
	payload := r.FormValue("payload")

	methods.SaveKanBan(pkgName, payload)

	response = "OK"

	callmet = true
	return
}
