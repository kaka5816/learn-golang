package main

// defer,延迟调用。先定义的后执行，后定义的先执行
func Defer() {
	defer func() {
		println("第一个defer")
	}()
	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	a := 0
	defer func() {
		println(a)
	}()
	a = 1
	//输出1,直接调用函数输出a后，再赋值
}

func DeferClosureV1() {
	a := 0
	defer func(a int) {
		println(a)
	}(a)
	a = 1
	//输出0，直接传参a=0
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}
