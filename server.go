package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type model struct {
	Time string `json:"time"`
}

func getTimeString() string {
	return time.Now().Format(time.RFC3339)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := getTimeString()
	result := model{Time: currentTime}
	encoded, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(encoded))
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}
