package internal

import (
	"fmt"
)

func FuncInputSample() {
	printTitle("FuncInputSample()")

	s := "s"
	printSection("値渡し")
	fmt.Println("before: " + s)
	funcInputFunction(s)       // sを"updated s"という値にしてみる
	fmt.Println("after: " + s) // >>> s

	printSection("参照渡し")
	sp := &s
	fmt.Println("before: " + s)
	funcPointerInputFunction(sp)       // sを"updated s"という値にしてみる
	fmt.Println("after: " + s)         // >>> updated s
	fmt.Println("after (*sp): " + *sp) // >>> updated s
}

func funcInputFunction(_ string) {
	_ = "updated s"
}

func funcPointerInputFunction(s *string) {
	*s = "updated s"
}

// 同じ型の引数は続けて書ける
func _(s1, s2 string, a int) error {
	return nil
}

// オーバーロードはできないっぽい
// func funcInputFunction(s *string) {
// 	s = "updated s"
// }
