package internal

import (
	"fmt"
)

func PrintSample() {
	printTitle("PrintSample()")

	a := 1
	fmt.Println("fmt.Println")
	fmt.Println("a", a)

	println("(built-in) println")
	println("a", a)

	fmt.Printf("%s: %d\n", "a", a)
}

func ForSample() {
	printTitle("ForSample()")

	for i := 0; i < 5; i++ {
		println("i:", i)
	}

	colors := []string{"red", "blue", "green", "yellow"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
}

func SwitchSample() {
	printTitle("SwitchSample")

	a := "a"

	switch a {
	case "a":
		println("a")
	case "b":
		println("b")
	default:
		println("default")
	}

	switch {
	case a == "a":
		println("a (cmp)")
	case a == "b":
		println("b (cmp)")
	default:
		println("default (cmp)")
	}

	// 複数trueになる場合
	switch {
	case true:
		println("true")
	case a == "a":
		println("a (cmp)")
	case a == "b":
		println("b (cmp)")
	default:
		println("default (cmp)")
	}
}

func DeferSample() {
	printTitle("DeterSample")

	defer println("defered comment 1")
	defer println("defered comment 2")
	defer println("defered comment 3")
	println("comment1")
	println("comment2")
	println("comment3")
}
