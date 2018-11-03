package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	methods "github.com/thestrukture/IDE/api/methods"
)

func GETApiKanban(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	pkgName := r.FormValue("pkg")
	response = mResponse(methods.GetKanBan(pkgName))

	callmet = true
	return
}
