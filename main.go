// TODO: create structs for time parsing and then save output
// Refs:
// https://stackoverflow.com/questions/67827329/get-array-of-nested-json-struct-in-go
// https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	FTEndpoint = "https://markets-data-api-proxy.ft.com/research/webservices/securities/v1/quotes"
)

var QuoteSymbols = map[string]string{
	"FTSE100": "FTSE:FSI",
	"DJIA":    "DJI:DJI",
	"GBPUSD":  "GBPUSD",
	"Gold":    "GC.1:CMX",
}

const timeFmt = "2006-01-02T15:04:05"

type RespTime struct {
	time.Time
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (rt *RespTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	time_, err := time.Parse(timeFmt, s)
	rt.Time = time_
	return err
}

type Data struct {
	Items []DataItems `json:"items"`
}

type Quote struct {
	DailyRet float64 `json:"change1DayPercent"`
}

type DataItems struct {
	SymbolInput string `json:"symbolInput"`
	Quote       Quote  `json:"quote"`
}

func ftQuoteEndpoint(symbols []string) string {
	return FTEndpoint + "?symbols=" + strings.Join(symbols, ",")
}

// Query FT for market data quotes.
// Returns a map of returned data for the queried symbols
func QueryFTMktData(symbols []string) (map[string]Quote, error) {
	endpoint := ftQuoteEndpoint(symbols)
	req, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(req.Body)

	resp := struct {
		Data    Data     `json:"data"`
		GenTime RespTime `json:"timeGenerated"`
	}{}

	err = d.Decode(&resp)

	if err != nil {
		return nil, err
	}

	data_map := make(map[string]Quote, len(symbols))

	for _, quote := range resp.Data.Items {
		symbol := quote.SymbolInput
		data_map[symbol] = quote.Quote
	}

	return data_map, nil
}

func main() {
	ft_symbols := make([]string, len(QuoteSymbols))
	InvQuoteSymbols := make(map[string]string, len(QuoteSymbols))

	i := 0
	for dispSym, ftSym := range QuoteSymbols {
		ft_symbols[i] = ftSym
		i++

		InvQuoteSymbols[ftSym] = dispSym
	}

	respQuotes, err := QueryFTMktData(ft_symbols)
	CheckErr(err)

	dispsymRets := make(map[string]float64, len(QuoteSymbols))

	for ftSym, ftQuote := range respQuotes {
		dispSym := InvQuoteSymbols[ftSym]
		dispsymRets[dispSym] = ftQuote.DailyRet
	}

	outJSON, err := json.Marshal(dispsymRets)
	CheckErr(err)

	fmt.Println(string(outJSON))
}
