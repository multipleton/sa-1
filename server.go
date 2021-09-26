package main

import (
	"fmt"
	"net/http"
	"time"
)

func getTimeString() string {
	return time.Now().Format(time.RFC3339)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := getTimeString()
	fmt.Fprintf(w, currentTime)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}
