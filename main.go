package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//go:embed index.html
var indexHTML string

func main() {
	http.HandleFunc("GET /api/{number}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		number := r.PathValue("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		f := fizzBuzz(n)

		if err := encoder.Encode(f); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("GET /api/fizz/{number}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		number := r.PathValue("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		if err := encoder.Encode(isFizz(n)); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("GET /api/buzz/{number}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		number := r.PathValue("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		if err := encoder.Encode(isBuzz(n)); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("GET /api/fizzbuzz/{number}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		number := r.PathValue("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		if err := encoder.Encode(isFizzBuzz(n)); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, indexHTML)
	})

	http.ListenAndServe(":8080", nil)

	log.Println("Listening on :8080")
}

func isFizzBuzz(n int) bool {
	return n%3 == 0 && n%5 == 0
}

func isFizz(n int) bool {
	return n%3 == 0
}

func isBuzz(n int) bool {
	return n%5 == 0
}

type fizzBuzzResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func fizzBuzz(n int) fizzBuzzResponse {
	if isFizzBuzz(n) {
		return fizzBuzzResponse{
			Message: "FizzBuzz",
			Details: fmt.Sprintf("%d is divisible by 3 and 5", n),
		}
	}
	if isFizz(n) {
		return fizzBuzzResponse{
			Message: "Fizz",
			Details: fmt.Sprintf("%d is divisible by 3", n),
		}
	}
	if isBuzz(n) {
		return fizzBuzzResponse{
			Message: "Buzz",
			Details: fmt.Sprintf("%d is divisible by 5", n),
		}
	}

	return fizzBuzzResponse{
		Message: strconv.Itoa(n),
		Details: fmt.Sprintf("%d is not divisible by 3 or 5", n),
	}
}
