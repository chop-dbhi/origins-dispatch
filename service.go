package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	StatusUnprocessableEntity = 422
)

// Single endpoint for receiving the payload request
func handler(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")

	if ct != "application/json" {
		return
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)

		if debug {
			fmt.Println("* Error reading request body")
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	e := EventPayload{}

	if err := json.Unmarshal(b, &e); err != nil {
		log.Println(err)

		if debug {
			fmt.Println("* Error decoding event")
		}

		w.WriteHeader(StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))

		return
	}

	if debug {
		fmt.Printf("* Received '%s' event\n", e.Event)
	}

	// TODO: run this and trigger in go routines once trigger is implemented.
	// Add wait group to confirm the processing finished.
	err = dispatch(&e)

	// Trigger webhooks
	// err = trigger(&e)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func serve() {
	addr := fmt.Sprintf("%s:%d", serveHost, servePort)

	// Register handler with server
	http.HandleFunc("/", handler)

	fmt.Printf("* Serving on http://%s\n", addr)

	if debug {
		fmt.Println("* Debugging enabled")
	}

	log.Fatal(http.ListenAndServe(addr, nil))
}
