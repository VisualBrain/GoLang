//http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

// StockItems is ...
type StockItems struct {
	Status           string
	Name             string
	Symbol           string
	LastPrice        float64
	Change           float64
	ChangePercent    float64
	Timestamp        string
	MSDate           int
	MarketCap        int
	Volume           int
	ChangeYTD        float64
	ChangePercentYTD float64
	High             float64
	Low              float64
	Open             float64
}

func main() {
	runtime.GOMAXPROCS(8)
	stockSymbols := []string{
		"googl",
		"aapl",
		"msft",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}
	start := time.Now()
	numComplete := 0
	for _, symbol := range stockSymbols {
		go func(symbol string) {
			response, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			defer response.Body.Close()

			body, _ := ioutil.ReadAll(response.Body)
			quote := new(StockItems)
			xml.Unmarshal(body, &quote) //everything in raw xml is converted into struct items
			fmt.Printf("%s:%.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}
	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("%-29s%s", "Execution time is:", elapsed)
}
