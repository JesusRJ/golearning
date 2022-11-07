package main

import (
	"net/http"
)

func main() {

	// Por baixo dos panos o go tem um DefaultServeMux (https://pkg.go.dev/net/http#ServeMux)
	// que relaciona todas as handlefunc
	http.HandleFunc("/info", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		// Usando assim "[]byte("ok")" não aloca o buffer na heap (não consome memória)
	})

	http.ListenAndServe(":8080", nil)
}
