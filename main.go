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

var colorMap = map[string]string{
	"white": "t:land !\"plains\" or produces:w t:land",
	"blue":  "t:land !\"island\" or produces:u t:land",
	"black": "t:land !\"swamp\" or produces:b t:land",
	"red":   "t:land !\"mountain\" or produces:r t:land",
	"green": "t:land !\"forest\" or produces:g t:land",
}

var typeMap = map[string]string{
	"fetch":    "is:fetchland",
	"tango":    "otag:cycle-tango-land",
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

	query := strings.Join(options, " ")
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
