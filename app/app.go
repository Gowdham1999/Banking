package main

import (
	"log"
	"net/http"
)

func Start() {
	//Route definition
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", customers)

	//Starting Server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
