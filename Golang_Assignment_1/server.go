package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/Enter", func(http.ResponseWriter, *http.Request) {
		log.Println("Welcome")
	})
	http.HandleFunc("/Execute", func(http.ResponseWriter, *http.Request) {
		log.Println("Your Request is Executing")
	})
	http.HandleFunc("/Help", func(http.ResponseWriter, *http.Request) {
		log.Println("Contact Server for Issues")
	})
	http.HandleFunc("/Wait", func(http.ResponseWriter, *http.Request) {
		log.Println("Your Process is on Hold")
	})
	http.HandleFunc("/Exit", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye")
	})
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Enter a Valid Keyword")
	})
	http.ListenAndServe(":9090", nil)
}
