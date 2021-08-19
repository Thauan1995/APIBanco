package main

import (
	"log"
	"net/http"
	"os"
	"site/APIBanco/rest"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	// API Banco
	r.HandleFunc("/banco", rest.APIBancoHandler)

	// Calculo Code
	r.HandleFunc("/banco/calculo", rest.CalculoHandler)

	http.Handle("/", router)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("Defaulting to port %s", port)
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
