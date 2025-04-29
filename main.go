package main

import (
	"net/http"
)

type RequestBody struct {
	Colors []string `json:"colors"`
	Types  []string `json:"types"`
}

var colorMap = map[string]string{
	"White": "produces:w",
	"Blue":  "produces:u",
	"Black": "produces:b",
	"Red":   "produces:r",
	"Green": "produces:g",
}

var typeMap = map[string]string{
	"Fetch": "is:fetchland",
	"Tango": "otag:cycle-tango-land",
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
