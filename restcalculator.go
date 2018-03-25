package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc(absPath, abs)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

const absPath = "/abs/"

func abs(w http.ResponseWriter, r *http.Request) {
	// Operation: Absolute Value
	// Example: /abs/-1  ->  1
	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
	// Try to parse input
	input := r.URL.Path[len(absPath):]
	result, err := strconv.Atoi(input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// Calculate absolute value
	if result < 0 {
		result = -result
	}
	// Write a response
	response := Response{"abs", input, strconv.Itoa(result)}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print("Error encoding response: ", err)
	}
}

// Response as specified in https://github.com/aunyks/newton-api
type Response struct {
	Operation  string `json:"operation"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}
