package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	methods "github.com/thestrukture/IDE/api/methods"
)

func POSTApiComposer(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	methods.GenerateComposeFile(r)
	response = "OK"

	callmet = true
	return
}