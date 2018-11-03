package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adlane/exec"
	"github.com/gorilla/sessions"
	"github.com/thestrukture/IDE/api/globals"
	methods "github.com/thestrukture/IDE/api/methods"
)

func ApiTerminal_realtime(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	c, err := globals.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	ctx := exec.InteractiveExec("bash", "-i")
	reader := methods.Reader{Conn: c}
	go ctx.Receive(&reader, 5*time.Hour)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)

			if ctx != nil {
				ctx.Cancel()
				ctx.Stop()
			}

			break
		}

		if len(message) != 0 {
			msg := string(message)

			if msg == "killnow\n" {
				fmt.Println("Restarting")
				ctx.Cancel()
				ctx.Stop()
				ctx = exec.InteractiveExec("bash", "-i")
				reader = methods.Reader{Conn: c}
				go ctx.Receive(&reader, 5*time.Hour)
			} else {
				ctx.Send(msg)
			}

		}

	}

	callmet = true
	return
}