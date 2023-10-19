package main

func main() {
	//Invoke()
	//Func5()
	//Func6Invoke()
	//Func7()
	//ClosureInvoke()
	YourNameInvoke()
	//DeferClosure()
	//DeferClosureV1()
	//DeferReturn()
	//DeferReturnV1()
	//DeferClosureLoopV1()
	//println("--------------")
	//DeferClosureLoopV2()
	//println("--------------")
	//DeferClosureLoopV3()

}

func Invoke() {
	str := Func0("小明")
	println(str)

	//str1是新变量，需要使用：=
	str1, err := Func2(1, 2)
	println(str1, err)

	str2, err := Func3(1, 2)
	println(str2, err)

}

//结构：  func 方法名(参数 类型） 返回类型

// 方法1
func Func0(name string) string {
	return "hello," + name
}

// 方法2
func Func1(a, b, c int, d string) (string, error) {
	return "hello", nil
}

// 方法3
func Func2(a, b int) (str string, err error) {
	str = "hello"
	return //上面赋值，可以无返回值，直接return
}

// 方法4
func Func3(a, b int) (str string, err error) {
	str = "hello"
	return "abc", nil //尽管上面赋值，但是以返回值为准
}
