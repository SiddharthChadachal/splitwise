package main

import (
	"fmt"
	"splitwise"
	"strconv"
)

func addBill() {
	id := getFlag("--id")
	tax := getFlag("--tax")

	if id == "" {
		fmt.Println("Missing --id")
		return
	}

	taxPercent := 0.0
	if tax != "" {
		taxPercent, _ = strconv.ParseFloat(tax, 64)
	}

	storage.Bills[id] = &splitwise.Bill{
		ID:         id,
		TaxPercent: taxPercent,
	}

	fmt.Println("Bill created:", id)
}
