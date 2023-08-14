package internal

import (
	"fmt"
)

func VarSample() {
	printTitle("VarSample()")

	// 変数の宣言
	var s1 string = "s1"
	s2 := "s2"

	var s3 string
	// nilではないっぽい
	// if s3 == nil {

	// }
	if s3 == "" {
		fmt.Println(`s3 is ""`)
	}

	fmt.Println("s1, s2, s3: " + s1 + s2 + s3)
}

func ConstSample() {
	printTitle("ConstSample()")

	s := "variable"
	const sc = "const variable"
	// const sc := "const variable" // :=ではダメ

	// sp := &sc // const型の変数をpointerで指せない
	// const sp = &s // pointerをconstで宣言できない

	println(s)
	println(sc)
}

func StringSample() {
	printTitle("StringSample()")

	s1 := "abcd"
	s2 := `abcd`
	r1 := 'a'

	println(s1, s2, r1)
}

func NilSample() {
	printTitle("NilSample()")

	var s *string

	// pointer型はデフォルトでnilになっている?
	if s == nil {
		fmt.Println("s is nil") // >>> s is nil
	}

	// nilを参照してみる
	// fmt.Println("*s: " + *s) // >>> panic: runtime error: invalid memory address or nil pointer dereference
}

func PointerSample() {
	printTitle("PointerSample()")

	s := "s"
	sp := &s

	// sp = "updated s" // できない
	*sp = "updated s"
}
