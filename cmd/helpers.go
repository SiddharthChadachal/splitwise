package main

import (
	"fmt"
	"splitwise"
	"strings"
)

func resolveParticipant(bill *splitwise.Bill, value string) (string, string) {
	value = strings.TrimSpace(strings.ToLower(value))
	fmt.Println(value)
	for _, p := range bill.Participants {
		if strings.ToLower(p.ID) == value {
			return p.ID, ""
		}
	}

	matches := []string{}

	for _, p := range bill.Participants {
		if strings.ToLower(p.Name) == value {
			matches = append(matches, p.ID)
		}
	}

	if len(matches) == 1 {
		return matches[0], ""
	}

	if len(matches) > 1 {
		return "", "Multiple matches found for " + value
	}

	return "", "No matches found for " + value
}
