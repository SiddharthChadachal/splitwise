package main

import (
	"fmt"
	"os"
	"strconv"
)

func billCommands() {
	if len(os.Args) < 3 {
		fmt.Println("bill commands: list | get | update | delete")
		fmt.Println("bill list usage: 	 bill list")
		fmt.Println("bill get usage:  	 bill get --bill \"bill id\"")
		fmt.Println("bill update usage:  	 bill update --bill \"bill id\" --new_tax \" your new tax \"")
		fmt.Println("bill delete usage:  	 bill delete --bill \"bill id\"")
		return
	}

	switch os.Args[2] {
	case "list":
		listBillCommand()
	case "get":
		getBillCommand()
	case "update":
		updateBillCommand()
	case "delete":
		deleteBillCommand()
	default:
		fmt.Println("Unknown Bill Command")
	}
}

func listBillCommand() {
	for id := range storage.Bills {
		fmt.Println("Bill: ", id)
	}
}

func getBillCommand() {
	billID := getFlag("--bill")
	b, ok := storage.Bills[billID]

	if !ok {
		fmt.Println("Bill not found")
		return
	}

	fmt.Printf("Bill %s: %+v\n", billID, b)
}

func updateBillCommand() {
	billID := getFlag("--bill")
	tax := getFlag("--new_tax")

	b, ok := storage.Bills[billID]
	if !ok {
		fmt.Println("Bill not found")
	}

	if tax != "" {
		tax, err := strconv.ParseFloat(tax, 64)

		if err == nil {
			b.TaxPercent = tax
			fmt.Println("tax updated")
		}
	}
}

func deleteBillCommand() {
	billID := getFlag("--bill")
	_, ok := storage.Bills[billID]

	if !ok {
		fmt.Println("Bill not found")
		return
	}

	delete(storage.Bills, billID)

	fmt.Println("Bill deleted: ", billID)
}
