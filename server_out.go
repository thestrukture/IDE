// File generated by Gopher Sauce
// DO NOT EDIT!!
package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/sessions"
	"github.com/thestrukture/IDE/api/assets"
	"github.com/thestrukture/IDE/api/handlers"
	sessionStore "github.com/thestrukture/IDE/api/sessions"
	"github.com/thestrukture/IDE/types"
)

func init() {

	gob.Register(&types.SoftUser{})

}
func main() {
	fmt.Fprintf(os.Stdout, "%v\n", os.Getpid())

	LaunchServer()
	sessionStore.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
		Domain:   "",
	}

	port := ":8884"
	if envport := os.ExpandEnv("$PORT"); envport != "" {
		port = fmt.Sprintf(":%s", envport)
	}
	log.Printf("Listenning on Port %v\n", port)

	//+++extendgxmlmain+++

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)
	http.Handle("/dist/", http.FileServer(&assetfs.AssetFS{Asset: assets.Asset, AssetDir: assets.AssetDir, Prefix: "web"}))
	http.HandleFunc("/", handlers.MakeHandler(handlers.Handler))

	h := &http.Server{Addr: port}

	go func() {
		<-stop
		log.Println("\nShutting down the server...")
		err := h.Close()

		if err != nil {
			panic(err)
		}

		Shutdown()
		log.Println("Server gracefully stopped")
	}()

	errgos := h.ListenAndServe()
	if errgos != nil {
		log.Fatal(errgos)
	}

}

//+++extendgxmlroot+++
