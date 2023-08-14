## インターフェース周り

### structがインターフェースを実装できているか
- コンパイラによるチェック
```
    var _ People = Student{} // コンパイラによるチェック
```

- 実行時のチェック
```
	// アサーションによる型チェック
	student := Student{}
	_, ok := interface{}(student).(People)
	if ok {
		println("Student is People")
	}
```