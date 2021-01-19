package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const message = "siemanko"

type Dog struct {
	Name string
	Age  int
}

func main() {

	newDog := Dog{
		Name: "Laos",
		Age:  2,
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(newDog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
