package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	uname := r.FormValue("uname")
	password := r.FormValue("password")
	fmt.Fprintf(w, "User Name = %s\n", uname)
	fmt.Fprintf(w, "Password = %s\n", password)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8084\n")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatal(err)
	}
}
