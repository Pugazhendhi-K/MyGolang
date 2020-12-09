package main

import (
	"fmt"
	"time"
)

func master(c chan string) {
	for i := 0; ; i++ {
		c <- "request"
	}
}

func worker(c chan string) {
	for i := 0; ; i++ {
		c <- "respond"
	}
}

func result(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 5)
	}
}

func main() {
	var c chan string = make(chan string)

	go master(c)
	go worker(c)
	go result(c)

	var input string
	fmt.Scanln(&input)
}
