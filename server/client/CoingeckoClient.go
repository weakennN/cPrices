package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Coin struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"current_price"`
}

const API_URL = "https://api.coingecko.com/api/v3"

func GetRates(symbols []string) ([]Coin, error) {
	url := API_URL + "/coins/markets/?vs_currency=usd&ids=" + strings.Join(symbols, ",")

	res, err := http.Get(url)

	if err != nil {
		log.Printf("Error sending request: %v", err)
		return []Coin{}, err
	}
	// TODO add defer when you learn what exactly it is
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		log.Printf("Unexpected response status: %d, body: %s", res.StatusCode, body)
		return []Coin{}, fmt.Errorf("unexpected response status: %d", res.StatusCode)
	}

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		log.Printf("Error reading response body: %v", readError)
		return []Coin{}, readError
	}

	var coins []Coin
	err = json.Unmarshal(body, &coins)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return []Coin{}, err
	}

	return coins, nil
}
