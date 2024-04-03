package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Use http.get")
		}

		log.Println("Welcome to homepage")
		outputString, _ := json.Marshal("Welcome to homepage")
		w.Write(outputString)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Use http.get")
		}

		log.Println("Welcome to homepage")
		outputString, _ := json.Marshal("Welcome to homepage")
		w.Write(outputString)
	})

	http.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Println("Use http.get")
		}

		log.Println("Who is root?")
		outputString, _ := json.Marshal("Who is root?")
		w.Write(outputString)
	})

	http.ListenAndServe(":8080", nil)

}
