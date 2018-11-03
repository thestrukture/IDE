package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thestrukture/IDE/api/globals"
	methods "github.com/thestrukture/IDE/api/methods"
)

func ApiSocket(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	c, err := globals.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	methods.AddConnection(c)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		if len(message) != 0 {
			methods.Broadcast(message)
		}

	}

	callmet = true
	return
}
