package main

func processRates(maxRates, minRates map[string]Rate, valCurs ValCurs) {

	for _, valute := range valCurs.Valute {
		if maxRate, ok := maxRates[valute.CharCode]; !ok || valute.Value > maxRate.Value {
			maxRates[valute.CharCode] = Rate{
				Value:    valute.Value,
				Name:     valute.Name,
				Date:     valCurs.Date,
				CharCode: valute.CharCode,
			}
		}

		if minRate, ok := minRates[valute.CharCode]; !ok || valute.Value < minRate.Value {
			minRates[valute.CharCode] = Rate{
				Value:    valute.Value,
				Name:     valute.Name,
				Date:     valCurs.Date,
				CharCode: valute.CharCode,
			}
		}

		updateAggregatedRates(valute, valCurs)
	}
}

func updateAggregatedRates(valute Valute, valVurs ValCurs) {

	aggregatedRate, exists := AggregatedRates[valute.CharCode]
	if !exists {
		aggregatedRate = AggregatedRate{}
	}

	aggregatedRate.TotalValue += valute.Value
	aggregatedRate.Count++
	AggregatedRates[valute.CharCode] = aggregatedRate
}
