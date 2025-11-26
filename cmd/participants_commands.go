package main

import (
	"fmt"
	"splitwise"
)

func addParticipant() {
	billID := getFlag("--bill")
	id := getFlag("--id")
	name := getFlag("--name")

	if billID == "" || id == "" || name == "" {
		fmt.Println("Missing: --bill --id --name")
		return
	}

	b, ok := storage.Bills[billID]
	if !ok {
		fmt.Println("Bill not found:", billID)
		return
	}

	b.AddParticipant(splitwise.Participant{
		ID:   id,
		Name: name,
	})

	fmt.Println("Participant added:", name)
}
