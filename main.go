package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//go:embed index.html
var indexHTML string

var ErrorNumberTooLarge = errors.New("number too large")
var ErrorNumberInvalid = errors.New("invalid number")

func main() {
	http.HandleFunc("GET /api/{number}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		n, err := getNumber(r)
		if err != nil {
			response := getErrorResponse(err)
			w.WriteHeader(response.StatusCode)
			if err := encoder.Encode(response); err != nil {
				log.Printf("Error encoding JSON: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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

		n, err := getNumber(r)
		if err != nil {
			response := getErrorResponse(err)
			w.WriteHeader(response.StatusCode)
			if err := encoder.Encode(response); err != nil {
				log.Printf("Error encoding JSON: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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

		n, err := getNumber(r)
		if err != nil {
			response := getErrorResponse(err)
			w.WriteHeader(response.StatusCode)
			if err := encoder.Encode(response); err != nil {
				log.Printf("Error encoding JSON: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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

		n, err := getNumber(r)
		if err != nil {
			response := getErrorResponse(err)
			w.WriteHeader(response.StatusCode)
			if err := encoder.Encode(response); err != nil {
				log.Printf("Error encoding JSON: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func getNumber(r *http.Request) (int, error) {
	num := r.PathValue("number")
	// iterate over string for non-numeric characters
	for _, c := range num {
		if c < '0' || c > '9' {
			return 0, ErrorNumberInvalid
		}
	}

	if len(num) > 7 {
		return 0, ErrorNumberTooLarge
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return 0, ErrorNumberInvalid
	}

	return n, nil
}

type errorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}

func getErrorResponse(err error) errorResponse {
	if errors.Is(err, ErrorNumberInvalid) {
		return errorResponse{
			Error:      "Invalid number. Please upgrade to a paid plan to use imaginary, non-real, or non-numeric numbers.",
			StatusCode: http.StatusBadRequest,
		}
	} else if errors.Is(err, ErrorNumberTooLarge) {
		return errorResponse{
			Error:      "Free tier is limited to 7 digits. Please upgrade to a paid plan to use larger numbers.",
			StatusCode: http.StatusPaymentRequired,
		}
	} else {
		return errorResponse{
			Error:      "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}
	}
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
