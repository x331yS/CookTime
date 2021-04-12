package main

import (
	"./pkg/cook-index"
	"fmt"
	"net/http"
)



func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//Main page handler
	http.HandleFunc("/", cook_index.IndexHandler)
	http.HandleFunc("/about/", cook_index.AboutHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
