package main

import (
	"net/http"
)

type APIResponse struct {
	Message string `json:"message"`
}

func fetchScryfallLands(w http.ResponseWriter, r *http.Request) {
	// Fetch the lands from Scryfall API
	resp, err := http.Get("https://api.scryfall.com/cards/search?q=type:land")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	// _, _ = w.ReadFrom(resp.Body)
}

func main() {
	http.HandleFunc("/api/lands", fetchScryfallLands)
	http.ListenAndServe(":8080", nil)
}
