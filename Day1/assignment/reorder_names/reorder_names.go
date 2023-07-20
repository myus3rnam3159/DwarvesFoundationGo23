package main

import(
	"fmt"
	"os"
)

func main(){

	args := os.Args;
	name := map[int]string{}

	for i:= 1; i < len(args) - 1; i++{
		name[i] = args[i]
	}

	rules := map[string][]int {

		"US": {1, 2, 3},
		"VN": {2, 3, 1},
	}

	order, existed := rules[ args[len(args) - 1] ]
	if !existed {
		order = rules["US"]
	}

	for _,o := range order {

		if value, existed := name[o]; existed{
			fmt.Print(value + " ")
		}
	}
	fmt.Println()
}
