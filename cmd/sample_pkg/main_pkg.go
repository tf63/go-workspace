package main

import (
	"fmt"

	sample "github.com/tf63/go_workspace/pkg/sample_pkg"
)

func main() {
	sl := []int{10, 20, 30}
	r := sample.Calculate(sl...)
	fmt.Println(r)
}
