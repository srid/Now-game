package main

import (
	"net/http"
	"github.com/hypebeast/go-osc/osc"
)

func runWebServer() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "game.html")
	})
	http.ListenAndServe(":8000", nil)
}

func main() {
	addr := "0.0.0.0:5000"
	// osc.Server starts a UDP server
	server := &osc.Server{Addr: addr}

	server.Handle("/muse/elements/experimental/concentration", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	server.Handle("/muse/elements/experimental/mellow", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	go runWebServer()

	println("Listening.", addr)
	server.ListenAndServe()
}
