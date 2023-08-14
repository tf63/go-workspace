package internal

import "strconv"

type Person struct {
	name string
	age  int
}

func StructSample() {
	printTitle("StructSample")

	person := Person{}
	println(person.name, person.age) // 初期値は反映されてる?
}

func StructPointerSample() {
	printTitle("StructPointerSample")

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
	// *&pp.name = "bb"
	// *&pp.age = 10

	println(pp.name, pp.age) // 更新できる

	pd := &Person{name: "aa", age: 30} // 宣言せずともクラスみたいにpointer型を定義できる
	pd.name = "bb"
	pd.age = 10
}

func StructFuncSample() {
	printTitle("StructFuncSample")

	p := Person{name: "aa", age: 10}
	funcStructInput(p)
	println("(value) name:", p.name, "age:", p.age) // 更新できない
	funcStructPointerInput(&p)
	println("(pointer) name:", p.name, "age:", p.age) // 更新できる
}

func funcStructInput(p Person) {
	p.name = "updated"
	p.age = 100
}

func funcStructPointerInput(p *Person) {
	p.name = "updated"
	p.age = 100
}

func (p Person) intro() string {
	return p.name + " " + strconv.Itoa(p.age)
}

func (p Person) changeNameWithValue(name string) {
	p.name = name
	_ = p.name
}

func (p *Person) changeNameWithPointer(name string) {
	p.name = name
}

func StructMethodSample() {
	printTitle("StructMethodSample")

	p := Person{name: "aa", age: 10}
	println("method呼び出し")
	println(p.intro())

	println("value receiver")
	p.changeNameWithValue("updated")
	println(p.intro()) // not updated

	println("pointer receiver")
	p.changeNameWithPointer("updated")
	println(p.intro()) // updated

}

type User struct {
	Person
}

func StructExtendSample() {
	printTitle("StructExtendSample")

	u := User{}
	u.changeNameWithPointer("user")
	println(u.intro())
}
