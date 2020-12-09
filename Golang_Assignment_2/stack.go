package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
	fmt.Println(str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack
	fmt.Println("Rule: Last in First Out")
	fmt.Println("Pushing Values into Stack")
	stack.Push("I'm First ")
	stack.Push("I'm Second ")
	stack.Push("I'm Third ")
	fmt.Println("Poping Values out of Stack")
	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
