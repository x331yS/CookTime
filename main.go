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
	http.HandleFunc("/search/", cook_index.SearchHandler)
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		fmt.Println(err)
	}
}
