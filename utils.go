package main

import "fmt"

func writeValues(rates map[string]Rate, title string) {
	fmt.Printf("%v\n", title)
	for _, rate := range rates {
		fmt.Printf("%v: %.2f на дату %v\n", rate.CharCode, rate.Value, rate.Date)
	}
}

func writeAggregatedValues(aggregatedRates map[string]AggregatedRate, title string) {
	fmt.Printf("%v\n", title)
	for charCode, aggregatedRate := range aggregatedRates {
		average := aggregatedRate.TotalValue / float64(aggregatedRate.Count)
		fmt.Printf("%v: Среднее=%.2f\n", charCode, average)
	}
}
