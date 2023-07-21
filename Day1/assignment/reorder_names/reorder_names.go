package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	name := args[1 : len(args)-1]

	rules := map[string][]int{
		"US": {0, 1, 2},
		"VN": {1, 2, 0},
	}

	order, existed := rules[args[len(args)-1]]
	if !existed {
		order = rules["US"]
	}

	for _, o := range order {
		if o < len(name) {
			fmt.Print(name[o] + " ")
		}
	}
	fmt.Println()
}