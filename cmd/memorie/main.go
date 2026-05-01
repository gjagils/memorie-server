package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getenv("PORT", "8090")
	addr := ":" + port

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)

	log.Printf("memorie-server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
