// Package Gemini is an implementation of the Gemini API in Golang.
package gemini

import (
	"encoding/json"
	"strings"
)

const (
	API_BASE                   = "https://api.gemini.com/" // Gemini API endpoint
	API_VERSION                = "v1"                      // Gemini API version
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                        // HTTP client timeout
)

// New return a instanciate gemini struct
func New(apiKey, apiSecret string) *Gemini {
	client := NewClient(apiKey, apiSecret)
	return &Gemini{client}
}

// gemini represent a gemini client
type Gemini struct {
	client *client
}

// GetTicker is used to get the current ticker values for a market.
func (b *Gemini) GetTicker(market string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "pubticker/"+strings.ToLower(market), "", false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &ticker); err != nil {
		return
	}
	return
}
