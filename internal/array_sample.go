package internal

func ArrayTypeSample() {
	printTitle("ArrayTypeSample()")

	var initArray []int // nil
	if initArray == nil {
		println("array is nil")
	}

	array := []int{} // length == 0の配列
	if array != nil {
		println("array is not nil")
	}

	// var a *int
	// a = array // 配列はpointerではない
}

func ArraySample() {
	printTitle("ArraySample()")

	array := [5]int{}
	// array = [3]int{} // 違う長さで再定義できない (初期値を代入できない)
	array = [5]int{1, 2, 3, 4, 5} // 同じ長さなら再定義できる (初期値を代入できる)
	println(array[0])

	// []intはスライス，[2]intは配列
	slice := []int{1, 2} // 長さを指定せずとも宣言できる
	// array = slice       // これは出来ない
	// slice = array // これも出来ない
	slice = array[:] // これは出来る
	// slice = [3]int{} // これも出来ない
	slice = []int{2, 3}       // これは出来る
	slice = []int{1, 2, 3, 4} // これも出来る
	println(slice[0])

	for i, a := range array[1:] {
		println(i, a)
	}
}

func ArraySizeSample() {
	printTitle("ArraySizeSample()")

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	div := slice[1:6]
	println("slice -> len:", len(slice), "cap:", cap(slice))
	// len -> 2, 3, 4, 5, 6の長さ (divの長さ)
	// cap -> 2, 3, 4, 5, 6, 7の長さ (divの先頭からsliceの終わりまでの長さ)
	println("div -> len:", len(div), "cap:", cap(div))
}
