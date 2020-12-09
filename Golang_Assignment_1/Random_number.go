package main

import "fmt"
"math/rand"
"time" 
func init() {

    rand.Seed(time.Now().UnixNano())
}
func main() {

    messages := make(chan string)
    for i := 0; i < 5; i++ {

		go func() { messages <-  rand.Intn(20) }()

		msg := <-messages
		fmt.Println(msg)
    }

    
}