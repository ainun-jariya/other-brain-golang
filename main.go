package main

import (
	"encoding/json"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/run", runHandler)
	http.ListenAndServe(":"+port, nil)
}
