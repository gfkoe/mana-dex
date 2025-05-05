package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestBody struct {
	Colors    []string `json:"colors"`
	LandTypes []string `json:"landTypes"`
}

var basicLands = map[string]string{
	"white": "plains",
	"blue":  "island",
	"black": "swamp",
	"red":   "mountain",
	"green": "forest",
}

var colorMap = map[string]string{
	"white": "w",
	"blue":  "u",
	"black": "b",
	"red":   "r",
	"green": "g",
}

var typeMap = map[string]string{
	"fetch":    "is:fetchland",
	"tango":    "is:tangoland",
	"shock":    "otag:shock-land",
	"triomes":  "is:triland",
	"surveil":  "otag:cycle-dual-surveil-land",
	"cycling":  "otag:cycle-akh-dual-cycling-land",
	"verge":    "otag:cycle-verge",
	"bond":     "otag:cycle-bbd-dual-land",
	"pain":     "otag:cycle-pain-land",
	"horizon":  "otag:cycle-horizon-land",
	"check":    "otag:cycle-check-land",
	"slow":     "otag:cycle-slow-land",
	"gates":    "otag:cycle-clb-thriving-gate",
	"thriving": "otag:cycle-jmp-thriving-land",
	"rainbow":  "exotic orchard or spire of industry or mana confluence or city of brass or evolving wilds or terramorphic expanse or myriad landscape or fabled passage",
}

func fetchLands(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var options []string

	for _, t := range req.LandTypes {
		fmt.Println(t)
		if lt, ok := typeMap[t]; ok {
			if t == "fetch" {
				for _, color := range req.Colors {
					if basic, ok := basicLands[color]; ok {
						options = append(options, fmt.Sprintf("%s o:%s", lt, basic))
					}
				}
			} else if t == "rainbow" {
				options = append(options, fmt.Sprintf("%s", lt))
			} else {
				for _, color := range req.Colors {
					options = append(options, fmt.Sprintf("%s produces:%s", lt, string(color)))
				}
			}
		}
	}

	for _, color := range req.Colors {
		if basic, ok := basicLands[color]; ok {
			options = append(options, fmt.Sprintf("!\" %s\"", basic))
		}
	}

	query := strings.Join(options, " or ")
	fmt.Println("Query:", query)
	baseUrl, _ := url.Parse("https://api.scryfall.com/cards/search")
	params := url.Values{}
	params.Add("q", query)
	baseUrl.RawQuery = params.Encode()
	fmt.Println("URL:", baseUrl.String())
	resp, err := http.Get(baseUrl.String())
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
