package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type RunRequest struct {
	Code string `json:"code"`
}

type RunResponse struct {
	Result string `json:"result"`
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	var req RunRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resp := RunResponse{Result: req.Code}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/run", withCORS(runHandler))

	logged := logMiddleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening to http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, logged))
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll((r.Body))
		r.Body = io.NopCloser(io.MultiReader(bytes.NewReader(body)))

		log.Println(">>>", r.Method, r.URL.Path)
		log.Println("Payload", string(body))

		next.ServeHTTP(w, r)
	})
}
