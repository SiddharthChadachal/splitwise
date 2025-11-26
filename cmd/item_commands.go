package main

import (
	"bufio"
	"fmt"
	"os"
	"splitwise"
	"strconv"
	"strings"
)

func addItem() {
	billID := getFlag("--bill")
	id := getFlag("--id")
	name := getFlag("--name")
	priceStr := getFlag("--price")
	shared := getFlag("--shared-by")

	if billID == "" || id == "" || name == "" || priceStr == "" {
		fmt.Println("Missing: --bill --id --name --price")
		return
	}

	b, ok := storage.Bills[billID]
	if !ok {
		fmt.Println("Bill not found:", billID)
		return
	}

	// Check for duplicate item ID
	for idx, existing := range b.Items {
		if existing.ID == id {
			fmt.Printf("Item already with ID: %s -> (%s - %.2f). Wanna update it ? (y,n): ",
				existing.ID, existing.Name, existing.Price)

			reader := bufio.NewReader(os.Stdin)
			ans, _ := reader.ReadString('\n')
			ans = strings.TrimSpace(strings.ToLower(ans))

			if ans != "y" {
				fmt.Println("Please use a different item ID.")
				return
			}

			_ = idx // allow update
		}
	}

	// Clean shared list: Split by comma (NOT character)
	rawParts := strings.Split(shared, ",")
	raw := []string{}
	for _, x := range rawParts {
		x = strings.TrimSpace(x)
		if x != "" {
			raw = append(raw, x)
		}
	}

	// Resolve IDs
	resolvedIDs := []string{}
	for _, entry := range raw {
		id, warn := resolveParticipant(b, entry)
		if warn != "" {
			fmt.Println("Warning:", warn)
			fmt.Println("Item NOT added. Fix Shared list.")
			return
		}
		resolvedIDs = append(resolvedIDs, id)
	}

	// Parse price
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Invalid price.")
		return
	}

	// Update or add new item
	updated := false
	for idx, it := range b.Items {
		if it.ID == id {
			b.Items[idx] = splitwise.Item{
				ID:       id,
				Name:     name,
				Price:    price,
				SharedBy: resolvedIDs,
			}
			fmt.Println("Item Updated:", name)
			updated = true
			break
		}
	}

	if !updated {
		b.AddItem(splitwise.Item{
			ID:       id,
			Name:     name,
			Price:    price,
			SharedBy: resolvedIDs,
		})
		fmt.Println("Item added:", name)
	}
}
