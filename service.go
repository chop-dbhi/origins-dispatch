package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

const (
	StatusUnprocessableEntity = 422
)

var getPayload = map[string]string{
	"name":    "Origins Dispatch Service",
	"version": version,
}

// Single endpoint for receiving the payload request
func handler(w http.ResponseWriter, r *http.Request) {
	debug := viper.GetBool("debug")

	// Respond to GET requests with name and version.
	if r.Method == "GET" {
		h := w.Header()

		h.Set("Content-Type", "application/json; charset=\"utf-8\"")

		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(getPayload)
		w.Write(b)

		return
	}

	// Response to non-POST methods
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ct := r.Header.Get("Content-Type")

	if ct != "application/json" {
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

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
	debug := viper.GetBool("debug")
	host := viper.GetString("serve_host")
	port := viper.GetInt("serve_port")

	addr := fmt.Sprintf("%s:%d", host, port)

	// Register handler with server
	http.HandleFunc("/", handler)

	fmt.Printf("* Serving on http://%s\n", addr)

	if debug {
		fmt.Println("* Debugging enabled")
	}

	log.Fatal(http.ListenAndServe(addr, nil))
}
