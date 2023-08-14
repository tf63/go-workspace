## goに関するメモ

### 環境
**プロジェクトの作成**
```
    go mod init github.com/tf63/go_workspace
    go mod tidy
```

### メモ

**変数の宣言**
- `:=`による宣言は関数の中でしか使えない

**関数呼び出し**
```
	// internal.varSample() // キャメルケースではダメ
	internal.VarSample() // パスカルケースじゃないといけない
```

**stringの定義**
- 参照型ではないっぽい
- デフォルトで`""`が入ってるっぽい
```
	var s3 string
	// nilではないっぽい
	// if s3 == nil {

	// }
	if s3 == "" {
		fmt.Println(`s3 is ""`)
	}
```

**オーバーロード**

できない

**値渡しと参照渡し**

普通の挙動
```
    func FuncInputSample() {
        printTitle("FuncInputSample()")

        s := "s"
        printSection("値渡し")
        fmt.Println("before: " + s)
        funcInputFunction(s)       // sを"updated s"という値にしてみる
        fmt.Println("after: " + s) // >> s

        printSection("参照渡し")
        sp := &s
        fmt.Println("before: " + s)
        funcPointerInputFunction(sp)       // sを"updated s"という値にしてみる
        fmt.Println("after: " + s)         // >> updated s
        fmt.Println("after (*sp): " + *sp) // >> updated s
    }
```

**ポインタ型の初期値**
- `nil`になっているっぽい
```
	var s *string // pointer型はデフォルトでnilになっている?
	if s == nil {
		fmt.Println("s is nil") // >>> s is nil
	}
```

**`nil`を参照すると?**
- invalid memory address
```
    fmt.Println("*s: " + *s) // >>> panic: runtime error: invalid memory address or nil pointer dereference
```

**標準出力**
- built-inの`println`と`fmt.Println`を使う (ひとまず)
- 引数はany型

- `println`はフォーマットができない
- `fmt.Println`はフォーマットができる?
- フォーマットするならprintf

- カンマ区切りで入力すると異なる型でもスペース区切りで出力してくれる
```
    a := 1
    println("a:", a)
```

**rune**
- `char`の代わり
- `a := 'a'`とすると`rune`型になる

**定数**
- constで宣言する
- :=では宣言できない
```
	const s = "const variable"
	// const s := "const variable" // :=ではダメ
```

**constとpointer**
```
	sp := &sc // const型の変数をpointerで指せない
	const sp = &s // pointerをconstで宣言できない
```

**switch**
- breakが無くとも1つしか評価されない
```
	switch a {
	case "a":
		println("a")
	case "b":
		println("b")
	default:
		println("default")
	}
```

- bool型なら引数を省略できる
```
	switch {
	case a == "a":
		println("a (cmp)")
	case a == "b":
		println("b (cmp)")
	default:
		println("default (cmp)")
	}
```

- 複数trueになる場合は一番上のものだけが評価される

**Defer**
- 関数の終了時に評価してくれる

```
    func DeferSample() {
        defer println("defered comment 1")
        defer println("defered comment 2")
        defer println("defered comment 3")

        println("comment1")
        println("comment2")
        println("comment3")
    }
```

結果
```
    comment1
    comment2
    comment3
    defered comment 3
    defered comment 2
    defered comment 1
```

- Defer句で変数の代入とかは出来ない

**struct**
- structの変数の初期値はstring, intの初期値となる
```
    type Person struct {
        name string
        age  int
    }
```

**struct pointer**
- メンバアクセスがデフォルトでアロー演算子のように働く
```
	p := Person{name: "aa", age: 30}
	pp := &p
	(*pp).name = "bb"
	(*pp).age = 10

	// これはできない
	// *pp.name = "bb"
	// *pp.age = 10

	// 参照せずともクラスみたいにも使える
	pp.name = "bb"
	pp.age = 10

	// 等価
	*&pp.name = "bb"
	*&pp.age = 10

	println(pp.name, pp.age) // 更新できる

	pd := &Person{name: "aa", age: 30} // 宣言せずともクラスみたいにpointer型を定義できる
	pd.name = "bb"
	pd.age = 10
```
- 普通はできない
```
	s := "s"
	sp := &s

	// sp = "updated s" // できない
	*sp = "updated s"
```

**structの値渡し，参照渡し**
- 思っている挙動と同じ
```
	p := Person{name: "aa", age: 10}
	funcStructInput(p) // 値渡し
	println("(value) name:", p.name, "age:", p.age) // 更新できない
	funcStructPointerInput(&p) // 参照渡し
	println("(pointer) name:", p.name, "age:", p.age) // 更新できる
```

**structのreceiver**
- value receiverとpointer receiverがある
```
    func (p Person) changeNameWithValue(name string) {
        p.name = name
    }

    func (p *Person) changeNameWithPointer(name string) {
        p.name = name
    }
```

- value receiverはフィールドを更新しないが，pointer receiverはフィールドを更新する
```
	println("value receiver")
	p.changeNameWithValue("updated")
	println(p.intro()) // not updated

	println("pointer receiver")
	p.changeNameWithPointer("updated")
	println(p.intro()) // updated
```

**structの埋め込み**
- 継承構文はないが, structを埋め込むことで同じようなことができる
```
    type User struct {
        Person
    }
```