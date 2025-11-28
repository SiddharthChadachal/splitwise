package main

import "fmt"

func printHelp() {
	fmt.Print("Splitwise Commands guide:\n\n")
	fmt.Print("splitwise <OPTIONS> <COMMAND> [ARGS...]\n\n")
	fmt.Println("OPTIONS:")
	fmt.Println("  help                                  - Show help")
	fmt.Println("  bill                                  - CRUD operations on Bills")
	fmt.Println("  participant                           - CRUD operations on participants")
	fmt.Print("  item                                  - CRUD operations on items\n\n")

	fmt.Println("COMMANDS:")
	fmt.Println("  add-bill --id <id> --tax <percent>							- Create a new bill")
	fmt.Println("  add-participant --bill <id> --id <pid> --name <name>					- Add participant to bill")
	fmt.Println("  add-item --bill <id> --id <iid> --name <name> --price <n> --shared-by p1,p2		- Add item to bill")
	fmt.Println("  split --bill <id>                     						- Calculate split for bill")
}
