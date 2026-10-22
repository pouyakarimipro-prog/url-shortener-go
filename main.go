package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Store struct {
	sync.RWMutex
	urls map[string]string
}

var store = Store{urls: make(map[string]string)}

func shorten(url string) string {
	h := sha256.Sum256([]byte(url))
	return base64.RawURLEncoding.EncodeToString(h[:])[:7]
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	log.Println("Starting service on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
