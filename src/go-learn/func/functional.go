package main

// 方法作为变量赋值
func Func4() {
	myFunc3 := Func3
	s, err := myFunc3(1, 2)
	println(s, err)
	_, _ = myFunc3(2, 3)
}

// 局部方法
func Func5() {
	fn := func(name string) string {
		return "hello," + name
	}
	str := fn("小王")
	println(str)
}

// 方法作为返回值
func Func6() func(name string) string {
	return func(name string) string {
		return "hello," + name
	}
}

func Func6Invoke() {
	fn := Func6()
	str := fn("小张")
	println(str)
}

// 匿名方法发起调用
func Func7() {
	fn := func(name string) string {
		return "hello," + name
	}("小李")
	println(fn)
}
