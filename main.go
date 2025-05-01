package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestBody struct {
	Colors    []string `json:"colors"`
	LandTypes []string `json:"landTypes"`
}

var colorMap = map[string]string{
	"white": "produces%3Aw",
	"blue":  "produces%3Au",
	"black": "produces%3Ab",
	"red":   "produces%3Ar",
	"green": "produces%3Ag",
}

var typeMap = map[string]string{
	"fetch":    "is%3Afetchland",
	"tango":    "otag%3Acycle-tango-land",
	"shock":    "otag%3Ashock-land",
	"triomes":  "is%3Atriland",
	"surveil":  "otag%3Acycle-dual-surveil-land",
	"cycling":  "otag%3Acycle-akh-dual-cycling-land",
	"verge":    "otag%3Acycle-verge",
	"bond":     "otag%3Acycle-bbd-dual-land",
	"pain":     "otag%3Acycle-pain-land",
	"horizon":  "otag%3Acycle-horizon-land",
	"check":    "otag%3Acycle-check-land",
	"slow":     "otag%3Acycle-slow-land",
	"gates":    "otag%3Acycle-clb-thriving-gate",
	"thriving": "otag%3Acycle-jmp-thriving-land",
	"other":    "exotic+orchard+or+spire+of+industry+or+mana+confluence+or+city+of+brass+or+evolving+wilds+or+terramorphic+expanse+or+myriad+landscape+or+fabled+passage",
}

func fetchLands(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Request:", req)

	fmt.Println("Colors:", req.Colors)
	fmt.Println("Land Types:", req.LandTypes)

	var options []string
	for _, color := range req.Colors {
		fmt.Println(color)
		if query, ok := colorMap[color]; ok {
			options = append(options, query)
		}
	}

	for _, t := range req.LandTypes {
		fmt.Println(t)
		if query, ok := typeMap[t]; ok {
			options = append(options, query)
		}
	}

	query := strings.Join(options, "+")
	fmt.Println("Query:", query)
	url := fmt.Sprintf("https://api.scryfall.com/cards/search?q=%s", query)
	fmt.Println("URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	// _, _ = w.ReadFrom(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/api/lands", fetchLands)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
