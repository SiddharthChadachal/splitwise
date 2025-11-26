package main

import "fmt"

func calcSplit() {
	billID := getFlag("--bill")

	if billID == "" {
		fmt.Println("Missing --bill")
		return
	}

	b, ok := storage.Bills[billID]
	if !ok {
		fmt.Println("Bill not found:", billID)
		return
	}

	result := b.CalculateSplit()

	fmt.Println("Split for Bill:", billID)
	for pid, amount := range result {
		fmt.Printf("%s => %.2f\n", pid, amount)
	}
}
