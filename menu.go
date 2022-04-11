package main

import (
	"fmt"
)

func main() {
	// var showmenu = menu()
	// showmenu()
	Table := CreateLinkTable()
	insertNode(1, 3, Table)
	insertNode(1, 2, Table)
	insertNode(1, 1, Table)
	printTable(Table)
	fmt.Println("---------------------")
	deleteNode(3, Table)
	printTable(Table)
	fmt.Println(findNode(2, Table).data)
}
func menu() func() {
	var message string = ""
	fun := func() {
		for {
			fmt.Print(">>")
			fmt.Scan(&message)
			switch message {
			case "help":
				fmt.Println("help")
			case "version":
				fmt.Println("version")
			case "quit":
				return
			default:
				fmt.Println("error command")
			}
		}
	}
	return fun
}
