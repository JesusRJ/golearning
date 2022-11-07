package main

import (
	"log"
	"net/http"
	"time"
)

type Mux struct{}

func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Xupa que é de uva"))
}

func main() {

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 7 * time.Second,
		Handler:      Mux{},
	}

	log.Fatal(s.ListenAndServe())
}
