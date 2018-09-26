package test3

// package main

import (
	"fmt"
	"net/http"
	"strings"
)

// SayhelloName, a route
func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("usr_long", r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello zouhl")

}

/*
func main() {
	http.HandleFunc("/", SayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
*/