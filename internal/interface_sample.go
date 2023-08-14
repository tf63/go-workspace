package internal

func InterfaceSample() {
	printTitle("InterfaceSample()")

	// 空のインターフェースは任意の値を代入できる
	var obj interface{}
	obj = "a"
	println(obj) // >>> (0x97180,0xc6098)

	obj = 2
	println(obj) // >>> (0x96b00,0xc5fd8)

	// 型アサーション
	value, _ := obj.(int)
	println(value) // >>> 2

	// 型 switch
	switch v := obj.(type) {
	case int:
		println("int", v) // >>> int 2
	case string:
		println("string", v)
	default:
		println("other")
	}
}

type People interface {
	intro() string
}

type Student struct {
	name string
}

type Teacher struct {
	name string
}

func (s Student) intro() string {
	return "student " + s.name
}

func (t Teacher) intro() string {
	return "teacher " + t.name
}

var _ People = Student{} // コンパイラによるチェック
var _ People = Teacher{}

func IntefaceImplementSample() {
	printTitle("InterfaceImplementSample()")

	// アサーションによる型チェック
	student := Student{}
	_, ok := interface{}(student).(People)
	if ok {
		println("Student is People")
	}

	peopleMethodCall(student)

	teacher := Teacher{}
	peopleMethodCall(teacher)
}

func peopleMethodCall(p People) {
	println(p.intro())
}
