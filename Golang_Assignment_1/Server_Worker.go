package main

import (
	"fmt"
	"time"
)

func server(v string) {
	for i := 0; i < 5; i++ {
		fmt.Println(v)
		time.Sleep(20 * time.Millisecond)
	}
}

func main() {
	go server("Worker Request")
	go server("Server Resspond")
	go server("Connected")
	time.Sleep(time.Second)
	fmt.Println("done")
	fmt.Println("Success")
}
