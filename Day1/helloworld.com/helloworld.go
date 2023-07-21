package main

import (
	"fmt" // standard package for format and output operation
	"net/http" // provide tools and functions for web app development
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, World!") // write the string Hello World to the response
	fmt.Println()
}

func main(){
	http.HandleFunc("/", helloHandler) // setup first
	http.ListenAndServe(":8080", nil) // run after
}