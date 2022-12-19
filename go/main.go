// TODO: create structs for time parsing and then save output
// Refs:
// https://stackoverflow.com/questions/67827329/get-array-of-nested-json-struct-in-go
// https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	FTEndpoint = "https://markets-data-api-proxy.ft.com/research/webservices/securities/v1/quotes"
)

var QuoteSymbols = []string{"FTSE:FSI", "DJI:DJI"}

type Data struct {
	Quotes []Quotes `json:"items"`
}

type Quotes struct {
	SymbolInput string `json:"symbolInput"`
}

func ftQuoteEndpoint(symbols []string) string {
	return FTEndpoint + "?symbols=" + strings.Join(symbols, ",")
}

// Query FT for market data quotes.
// Returns a dict of returned data for the queried symbols
func QueryFTMktData(symbols []string) (string, error) {
	endpoint := ftQuoteEndpoint(symbols)
	req, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	d := json.NewDecoder(req.Body)

	resp := struct {
		RespData Data      `json:"data"`
		GenTime  time.Time `json:"timeGenerated"`
	}{}

	err = d.Decode(&resp)

	if err != nil {
		return "", err
	}

	return "", nil

}

func main() {
	resp, err := QueryFTMktData(QuoteSymbols)

	fmt.Println(resp)
	fmt.Println(err)
}
