package main

import (
	"fmt"
	//"github.com/thegera4/price_calculator/cmdmanager"
	"github.com/thegera4/price_calculator/filemanager"
	"github.com/thegera4/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index]) //go routine do not return values/errors

		/*if err != nil {
			fmt.Println("Could not process job!")
			fmt.Println(err)
		}*/
	}

	for _, doneChan := range doneChans {
		<-doneChan //wait for the signal of the every channel
	}

}