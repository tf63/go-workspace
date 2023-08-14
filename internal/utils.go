package internal

import "fmt"

func printTitle(title string) {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(title)
	fmt.Println("----------------------------------------------------------------")
}

func printSection(section string) {
	fmt.Println()
	fmt.Println("**** " + section + "****")
}
