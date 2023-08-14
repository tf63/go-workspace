## goの配列周り

**arrayの型**
```
	var initArray []int // nil
	if initArray == nil {
		println("array is nil")
	}

	array := []int{} // length == 0の配列
	if array != nil {
		println("array is not nil")
	}
```

- pointerではない
```
	// var a *int
	// a = array // 配列はpointerではない
```

**arrayの長さ**
- 固定長のarrayと可変長のsliceがある
```
    // [2]intは配列
	array := [2]int{}
	// array = [3]int{} // 違う長さで再定義できない (初期値を代入できない)
	array = [2]int{} // 同じ長さなら再定義できる (初期値を代入できる)

	// []intはスライス
	slice := []int{1, 2} // 長さを指定せずとも宣言できる
	// array = slice       // これは出来ない
	// slice = array // これも出来ない
	slice = array[:] // これは出来る
	// slice = [3]int{} // これも出来ない
	slice = []int{2, 3}       // これは出来る
	slice = []int{1, 2, 3, 4} // これも出来る
```

- lenとcapの2つで長さを取得できる
```
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	div := slice[1:6]
    
	println("slice -> len:", len(slice), "cap:", cap(slice))

	// len -> 2, 3, 4, 5, 6の長さ (divの長さ)
	// cap -> 2, 3, 4, 5, 6, 7の長さ (divの先頭からsliceの終わりまでの長さ)
	println("div -> len:", len(div), "cap:", cap(div))
```