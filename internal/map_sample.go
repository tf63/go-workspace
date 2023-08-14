package internal

func MapSample() {
	printTitle("MapSample()")

	m := map[string]int{"a": 1}
	println("defined", m["a"])

	m["a"] = 0
	println("update value", m["a"])

	m["b"] = 2
	println("new key", m["b"])

	delete(m, "b")
	println("deleted key", m["b"]) // 0

	for key, value := range m {
		println("key:", key, "value:", value) // 削除したbは表示されない
	}
}
