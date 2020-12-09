package main

import "fmt"

func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []string) []string {
	element := queue[0]
	fmt.Println("Dequeued:", element)
	return queue[1:]
}

func main() {
	var queue []string
	fmt.Println("Customers entering Store in Queue")
	queue = enqueue(queue, "Dennis")
	queue = enqueue(queue, "Aang")
	queue = enqueue(queue, "Jackie")

	fmt.Println("Customers Waiting for Service")
	fmt.Println("Queue:", queue)

	fmt.Println("One Customer got Service and left")
	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	fmt.Println("A new Customer entering the Queue")
	queue = enqueue(queue, "Jake")
	fmt.Println("Queue:", queue)
}
