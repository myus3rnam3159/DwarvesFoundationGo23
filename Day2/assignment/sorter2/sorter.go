package main

import (
	sortUtils "SortUtils"
	fl "flag"
	"fmt"
)

func main(){
	var flags = map[string]sortUtils.SortFunc {
		"int": sortUtils.SortIntegers,
		"string": sortUtils.SortStrings,
		"mix": sortUtils.SortMixs,
	}

	for flagKey:= range flags{
		fl.Bool(flagKey, false, flagKey + " input list")
	}
	fl.Parse()

	var args = fl.Args()

	for k, f := range flags{
		if fl.Lookup(k).Value.String() == "true"{
			f(&args)
			break
		}
	}
	fmt.Print("Output: ")

	for _, e := range args{
		fmt.Print(e + " ")
	}
	fmt.Println()
}