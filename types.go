package main

type ValCurs struct {
	Date   string            `json:"Date"`
	Valute map[string]Valute `json:"Valute"`
}

type Valute struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
}

type Rate struct {
	Value    float64
	Name     string
	Date     string
	CharCode string
}

type AggregatedData struct {
	TotalValue float64
	Count      int
}

type AggregatedRate struct {
	TotalValue float64
	Count      int
}

var maxRates = make(map[string]Rate)
var minRates = make(map[string]Rate)
var AggregatedRates = make(map[string]AggregatedRate)
