package main

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"net/http"
)

type MindState struct {
	Quality string
	Value   float32 // Value ranges from 0.0 to 1.0
}

func (s MindState) GetPercent() int {
	// Why add 0.5? To get the 'nearest integer' instead of just the floor.
	return int(s.Value*100 + 0.5)
}

type Stream chan MindState

func runWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "game.html")
	})
	http.ListenAndServe(":8000", nil)
}

func main() {
	addr := "0.0.0.0:5000"
	// osc.Server starts a UDP server
	server := &osc.Server{Addr: addr}

	stream := make(Stream)

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

	println("Listening.", addr)
	server.ListenAndServe()
}
