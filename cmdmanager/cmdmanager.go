package cmdmanager

import "fmt"

type CMDmanager struct {
}

func (cmd CMDmanager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prizes. Confirm every price with ENTER. Press 0 to check the results:")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDmanager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDmanager {
	return CMDmanager{}
}