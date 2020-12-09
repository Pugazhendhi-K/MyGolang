package main

import "net/http"

func main() {

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello Pugazh welcome to world"))
	})

	http.ListenAndServe(":8084", nil)
}
