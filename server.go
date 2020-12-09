package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		/*	d, _ := ioutil.ReadAll(r.Body)

			fmt.Fprintf(rw, "Hello %s", d) */
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye World")
	})
	http.ListenAndServe(":8084", nil)
}
