package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	rules := map[string][]int{
		"US": {0, 2, 1},
		"VN": {1, 2, 0},
	}

	order, existed := rules[args[len(args)-1]]
	if !existed {
		fmt.Println("Country code not supported")
		return
	}

	for i, idx := range order {
		if idx < len(args)-1 {
			fmt.Print(args[idx])
			if i < len(order)-1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}
