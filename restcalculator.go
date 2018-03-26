package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"math"
	"encoding/json"
)

var operations = []Operation{
	{"abs", absolute},
	{"log", logarithm},
}

func main() {
	log.Print("Starting restcalculator")
	for _, operation := range operations {
		http.Handle(operation.endpoint(), operation)
	}
	log.Fatal(http.ListenAndServe(":5000", nil))
}

type Operation struct {
	Name    string
	Execute func(expression string) string
}

func (operation Operation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
	// Read input from path
	input := r.URL.Path[len(operation.endpoint()):]
	// Execute operation
	result := operation.Execute(input)
	// Write a response
	resp := response{operation.Name, input, result}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Print("Error encoding response: ", err)
	}
}

func (operation Operation) endpoint() string {
	return "/" + operation.Name + "/"
}

// Response as specified in https://github.com/aunyks/newton-api
type response struct {
	Operation  string `json:"operation"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

func absolute(expression string) (string) {
	number, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return "NaN"
	}
	// Calculate absolute value
	if number < 0 {
		number = -number
	}
	return strconv.FormatFloat(number, 'f', -1, 64)
}

func logarithm(expression string) (string) {
	// f(x) = log_b(x)
	variables := strings.Split(expression, "|")
	if len(variables) != 2 {
		return "invalid syntax"
	}
	base, err := strconv.ParseFloat(variables[0], 64)
	if err != nil {
		return "NaN"
	}
	x, err := strconv.ParseFloat(variables[1], 64)
	if err != nil {
		return "NaN"
	}
	// Calculate logarithm value
	f := math.Log(x) / math.Log(base)
	return strconv.FormatFloat(f, 'f', -1, 64)
}
