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
	*Shock - otag:shock-land
*Triomes - is:triland
*Surveil - otag:cycle-dual-surveil-land
*Cycling - otag:cycle-akh-dual-cycling-land
Verge - otag:cycle-verge
Bond - otag:cycle-bbd-dual-land
Pain - otag:cycle-pain-land
Horizon - otag:cycle-horizon-land
Check - otag:cycle-check-land
Slow - otag:cycle-slow-land
Gates - otag:cycle-clb-thriving-gate
Thriving - otag:cycle-jmp-thriving-land
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
