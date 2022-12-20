package main

import "net/http"

func main() {
	http.HandleFunc("/", nil) // hello
	http.ListenAndServe(":8080", nil)
}
