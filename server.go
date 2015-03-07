package main

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

type MindState struct {
	Quality string
	Value   float32 // Value ranges from 0.0 to 1.0
}

func (s MindState) GetPercent() int {
	// Why add 0.5? To get the 'nearest integer' instead of just the floor.
	return int(s.Value*100 + 0.5)
}

func dummyStreamHandler(ws *websocket.Conn) {
	for {
		fmt.Fprintf(ws, "Time now is: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

func runWebServer() {
	http.Handle("/dummy", websocket.Handler(dummyStreamHandler))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset("build/elm/Game.html")
		if err != nil {
			fmt.Fprintf(w, "ERROR: Game.html missing in executable")
		} else {
			fmt.Fprintf(w, "%s", string(data))
		}
	})
	fmt.Printf("Web server running at: http://localhost:8000/\n")
	http.ListenAndServe(":8000", nil)
}

func main() {
	addr := "0.0.0.0:5000"
	// osc.Server starts a UDP server
	server := &osc.Server{Addr: addr}

	stream := make(chan MindState)

	server.Handle("/muse/elements/experimental/concentration", func(msg *osc.Message) {
		value := msg.Arguments[0].(float32)
		stream <- MindState{"concentration", value}
	})

	server.Handle("/muse/elements/experimental/mellow", func(msg *osc.Message) {
		value := msg.Arguments[0].(float32)
		stream <- MindState{"mellow", value}
	})

	go runWebServer()

	go func() {
		for state := range stream {
			fmt.Printf("%15v => %3v%% (%v)\n", state.Quality, state.GetPercent(), state.Value)
		}
	}()

	fmt.Printf("Listening at Muse OSC url: osc.udp://%s\n", addr)
	server.ListenAndServe()
}
