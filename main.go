package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type quote struct {
	Author string `json:"author"`
	Quote string `json:"quote"`
	Permalink string `json:"permalink"`
}

func main() {
	
	res, err := http.Get("http://quotes.stormconsultancy.co.uk/random.json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("Error trying to get quote")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	jsonBody := quote{}
	json.Unmarshal(body, &jsonBody)
	
	s := normalizeQuote(jsonBody.Quote)
	fmt.Printf("\"%s\"\n - %s\n", s, jsonBody.Author)

}

func normalizeQuote(s string) string {
	output := strings.ReplaceAll(s, ". ", ".\n")
	output = strings.ReplaceAll(s, "; ", ";\n")
	output = strings.ReplaceAll(s, ", ", ",\n")
	output = strings.ReplaceAll(s, ": ", ":\n")
	return output
}
