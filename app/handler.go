package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name  string `json:"name" xml:"Name1"`
	Age   int    `json:"age" xml:"Age1"`
	City  string `json:"city" xml:"City1"`
	State string `json:"state" xml:"State1"`
}

func customers(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{
		{"Gowdham", 23, "Puducherry", "Puducherry"},
		{"Kowsalya", 21, "Puducherry", "Puducherry"},
		{"Sathish", 25, "Mangalore", "Karnataka"},
		{"Vignesh", 24, "Puducherry", "Puducherry"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Go routing!!!")
}
