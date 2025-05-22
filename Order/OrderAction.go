package main

import (
	"fmt"
	"sampleProject/db"
)

func main() {
	db1, err := db.GetConnection()
	if err != nil {
		fmt.Println("Error connecting to DB:", err)
	}
	defer db1.Close()
	orderList := db.GetOrderList(db1)
	for _, order := range orderList {
		fmt.Println("Account Id: ", order.GetAcctId())
		fmt.Println("Account User: ", order.GetAcctUser())
		fmt.Println("Currency Pair: ", order.GetCurrPair())
		fmt.Println("Deal Qty: ", order.GetDealQty())
		fmt.Println("Countrer Qty: ", order.GetCtrQty())
		fmt.Println("--------------------------------------------------")
	}
}
