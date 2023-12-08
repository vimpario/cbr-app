package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchData(date time.Time) (ValCurs, error) {
	url := fmt.Sprintf("https://www.cbr-xml-daily.ru/archive/%d/%02d/%02d/daily_json.js", date.Year(), date.Month(), date.Day())

	resp, err := http.Get(url)
	if err != nil {
		return ValCurs{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			// Нет данных для данной даты, возвращаем пустой результат
			return ValCurs{}, nil
		}
		return ValCurs{}, fmt.Errorf("status code error for date %s: %d %s", date.Format("2006-01-02"), resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ValCurs{}, err
	}

	var valCurs ValCurs
	err = json.Unmarshal(data, &valCurs)
	if err != nil {
		return ValCurs{}, fmt.Errorf("error while parse JSON for date %s: %v", date.Format("2006-01-02"), err)
	}

	return valCurs, nil
}
