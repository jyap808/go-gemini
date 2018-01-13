package gemini

type Ticker struct {
	Bid  float64 `json:"bid,string"`  // Innermost bid
	Ask  float64 `json:"ask,string"`  // Innermost ask
	Last float64 `json:"last,string"` // The price at which the last order executed
}
