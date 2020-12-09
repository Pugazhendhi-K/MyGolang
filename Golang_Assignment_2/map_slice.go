package main

import "fmt"

func main() {

	m := map[string][]string{
		"cat": {"orange", "grey"},
		"dog": {"black"},
	}

	res := append(m["dog"], "brown")
	fmt.Println(res)

	m["fish"] = []string{"orange", "red"}

	fmt.Println(m["fish"])

	for i := range m["fish"] {
		fmt.Println(i, m["fish"][i])
	}
}
