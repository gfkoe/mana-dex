package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var options []string
	for _, color := range req.Colors {

		if query, ok := colorMap[color]; ok {
			options = append(options, query)
		}
	}

	for _, t := range req.Types {

		if query, ok := typeMap[t]; ok {
			options = append(options, query)
		}
	}

	query := strings.Join(options, "+")
	url := fmt.Sprintf("https://api.scryfall.com/cards/search?q=%s", query)
	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	// _, _ = w.ReadFrom(resp.Body)
	fmt.Println("Fetched lands from Scryfall API")
	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/api/lands", fetchScryfallLands)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
