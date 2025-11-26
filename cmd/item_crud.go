package main

import (
	"fmt"
	"os"
	"splitwise"
)

func itemCommands() {
	if len(os.Args) < 3 {
		fmt.Println("item commands: list | get | update | delete")
		fmt.Println("item list usage: 	 item --bill \" bill ID\"")
		fmt.Println("item get usage:  	 item get --bill \"bill id\" --item_id \" item ID \"")
		fmt.Println("item update usage:  	 participant update --bill \"bill id\" --pid \" participant ID \" --new_name \" new name \"")
		fmt.Println("item delete usage:  	 participant delete --bill \"bill id\" --pid \" participant ID \"")
		return
	}

	switch os.Args[2] {
	case "list":
		listItemCommand()
	case "get":
		getItemCommand()
	case "update":
		updateItemCommand()
	case "delete":
		deleteItemCommand()
	default:
		fmt.Println("Unknown Item Command")
	}
}

func listItemCommand() {
	billID := getFlag("--bill")
	b := storage.Bills[billID]

	for _, it := range b.Items {
		fmt.Printf("%s -> %s (%.2f)\n", it.ID, it.Name, it.Price)
	}
}

func getItemCommand() {
	billID := getFlag("--bill")
	itemID := getFlag("--item_id")
	b := storage.Bills[billID]

	for _, it := range b.Items {
		if it.ID == itemID {
			fmt.Printf("Item: %+v\n", it)
			return
		}
	}

	fmt.Println("Item not found")
}

func updateItemCommand() {
	billID := getFlag("--bill")
	itemID := getFlag("--item_id")
	name := getFlag("--name")
	price := getFlag("--price")
	sharedBy := getFlag("--shared-By")

	os.Args = []string{"cmd", "add-item",
		"--bill", billID,
		"--id", itemID,
	}

	if name != "" {
		os.Args = append(os.Args, "--name", name)
	}

	if price != "" {
		os.Args = append(os.Args, "--price", price)
	}

	if sharedBy != "" {
		os.Args = append(os.Args, "--shared-By", sharedBy)
	}

	addItem()
}

func deleteItemCommand() {
	billID := getFlag("--bill")
	itemID := getFlag("--item_id")
	b := storage.Bills[billID]

	newItems := []splitwise.Item{}
	for _, it := range b.Items {
		if it.ID != itemID {
			newItems = append(newItems, it)
		}
	}

	b.Items = newItems
	fmt.Println(" Items deleted: ", itemID)
}
