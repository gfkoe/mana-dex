package main

import (
	"fmt"
	"net/http"
)

type RequestBody struct {
	Colors []string `json:"colors"`
	Types  []string `json:"types"`
}

var colorMap = map[string]string{
	"White": "produces%3ACw",
	"Blue":  "produces%3ACu",
	"Black": "produces%3ACb",
	"Red":   "produces%3ACr",
	"Green": "produces%3ACg",
}

var typeMap = map[string]string{
	"Fetch":    "is%3ACfetchland",
	"Tango":    "otag%3ACcycle-tango-land",
	"Shock":    "otag%3ACshock-land",
	"Triomes":  "is%3ACtriland",
	"Surveil":  "otag%3ACcycle-dual-surveil-land",
	"Cycling":  "otag%3ACcycle-akh-dual-cycling-land",
	"Verge":    "otag%3ACcycle-verge",
	"Bond":     "otag%3ACcycle-bbd-dual-land",
	"Pain":     "otag%3ACcycle-pain-land",
	"Horizon":  "otag%3ACcycle-horizon-land",
	"Check":    "otag%3ACcycle-check-land",
	"Slow":     "otag%3ACcycle-slow-land",
	"Gates":    "otag%3ACcycle-clb-thriving-gate",
	"Thriving": "otag%3ACcycle-jmp-thriving-land",
	"Other":    "exotic+orchard+or+spire+of+industry+or+mana+confluence+or+city+of+brass+or+evolving+wilds+or+terramorphic+expanse+or+myriad+landscape+or+fabled+passage",
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
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
