package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/sessions"
	"github.com/thestrukture/IDE/api/assets"
	"github.com/thestrukture/IDE/api/handlers"
	sessionStore "github.com/thestrukture/IDE/api/sessions"
	"github.com/thestrukture/IDE/types"
)

func init() {

	gob.Register(&types.SoftUser{})

}
func dummy_timer() {
	dg := time.Second * 5
	log.Println(dg)
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
		errgos := h.ListenAndServe()
		if errgos != nil {
			log.Fatal(errgos)
		}
	}()

	<-stop

	log.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	h.Shutdown(ctx)

	Shutdown()

	log.Println("Server gracefully stopped")

}

//+++extendgxmlroot+++
