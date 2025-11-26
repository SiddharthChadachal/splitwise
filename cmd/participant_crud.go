package main

import (
	"fmt"
	"os"
	"splitwise"
)

func participantCommands() {
	if len(os.Args) < 3 {
		fmt.Println("participant commands: list | get | update | delete")
		fmt.Println("participant list usage: 	 participant list")
		fmt.Println("participant get usage:  	 participant get --bill \"bill id\" --pid \" participant ID \"")
		fmt.Println("participant update usage:  	 participant update --bill \"bill id\" --pid \" participant ID \" --new_name \" new name \"")
		fmt.Println("participant delete usage:  	 participant delete --bill \"bill id\" --pid \" participant ID \"")
		return
	}

	switch os.Args[2] {
	case "list":
		listParticipantCommand()
	case "get":
		getParticipantCommand()
	case "update":
		updateParticipantCommand()
	case "delete":
		deleteParticipantCommand()
	default:
		fmt.Println("Unknown Participant Command")
	}
}

func listParticipantCommand() {
	billID := getFlag("--bill")
	b, ok := storage.Bills[billID]

	if !ok {
		fmt.Println("Participant not Found")
		return
	}

	for _, p := range b.Participants {
		fmt.Printf("%s -> %s\n", p.ID, p.Name)
	}
}

func getParticipantCommand() {
	billID := getFlag("--bill")
	pid := getFlag("--pid")

	b := storage.Bills[billID]

	for _, p := range b.Participants {
		if p.ID == pid {
			fmt.Printf("Participant: %+v\n", p)
			return
		}
	}

	fmt.Println("Participant Not Found")
}

func updateParticipantCommand() {
	billID := getFlag("--bill")
	pid := getFlag("--pid")
	new_name := getFlag("--new_name")
	b := storage.Bills[billID]

	for i, p := range b.Participants {
		if p.ID == pid {
			b.Participants[i].Name = new_name
			fmt.Printf("Updated Participant: %+v\n", p.ID)
			return
		}
	}

	fmt.Println("Participant Not Found")
}

func deleteParticipantCommand() {
	billID := getFlag("--bill")
	pid := getFlag("--pid")

	b := storage.Bills[billID]

	newList := []splitwise.Participant{}
	for _, p := range b.Participants {
		if p.ID != pid {
			newList = append(newList, p)
		}
	}

	b.Participants = newList

	for i := range b.Items {
		fixed := []string{}

		for _, sb := range b.Items[i].SharedBy {
			if sb != pid {
				fixed = append(fixed, sb)
			}
		}

		b.Items[i].SharedBy = fixed
	}

	fmt.Println("Participant Deleted: ", pid)
}
