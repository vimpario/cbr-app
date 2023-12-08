package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func(daysAgo int) {
			defer wg.Done()

			date := time.Now().AddDate(0, 0, -daysAgo-i-1)
			valCurs, err := fetchData(date)
			if err != nil {
				log.Fatal(err)
			}
			mu.Lock()
			processRates(maxRates, minRates, valCurs)
			mu.Unlock()
		}(i)

	}

	wg.Wait()

	writeValues(maxRates, "Максимальные")
	writeValues(minRates, "Минимальные")
	writeAggregatedValues(AggregatedRates, "Сумма курсов")

}
