package main

import (
	"fmt"
	"os"

	"splitwise"
)

var storage *splitwise.Storage

const storagePath = "data.json"

// flag parser
// utility: get flag value after a flag like: --name Alice
func getFlag(flag string) string {
	for i, arg := range os.Args {
		if arg == flag && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
	}
	return ""
}

func main() {
	// 1. Load persistent storage
	s, err := splitwise.LoadStorage(storagePath)
	if err != nil {
		fmt.Println("Error loading storage:", err)
		return
	}
	storage = s

	// 2. No command? show help
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "bill":
		billCommands()
	case "participant":
		participantCommands()
	case "item":
		itemCommands()
	case "help":
		printHelp()
	case "add-bill":
		addBill()
	case "add-participant":
		addParticipant()
	case "add-item":
		addItem()
	case "split":
		calcSplit()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		printHelp()
	}

	// 3. Save changes
	err = storage.Save(storagePath)
	if err != nil {
		fmt.Println("Failed to save:", err)
	}

}
