package main

import "github.com/hypebeast/go-osc/osc"

func main() {
	addr := "0.0.0.0:5000"
	server := &osc.Server{Addr: addr}

	server.Handle("/muse/elements/experimental/concentration", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	server.Handle("/muse/elements/experimental/mellow", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	println("Listening.", addr)
	server.ListenAndServe()
}
