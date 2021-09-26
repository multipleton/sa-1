package main

import (
	"fmt"
	"net/http"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}
