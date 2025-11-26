package main

import "fmt"

func printHelp() {
	fmt.Println("Splitwise CLI Commands:")
	fmt.Println("  help                                  - Show help")
	fmt.Println("  bill                                  - CRUD operations on Bills")
	fmt.Println("  participant                           - CRUD operations on participants")
	fmt.Println("  item                                  - CRUD operations on items")
	fmt.Println("  add-bill --id <id> --tax <percent>    - Create a new bill")
	fmt.Println("  add-participant --bill <id> --id <pid> --name <name>")
	fmt.Println("  add-item --bill <id> --id <iid> --name <name> --price <n> --shared-by p1,p2")
	fmt.Println("  split --bill <id>                     - Calculate split for bill")
}
