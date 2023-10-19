package main

// 闭包：方法+它绑定的运行上下文
//变量 name +方法本身 func() string

func Closure(name string) func() string {
	return func() string {
		return "hello," + name
	}
}

func ClosureInvoke() {
	c := Closure("大明")
	println(c())
}
