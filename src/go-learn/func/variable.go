package main

//不定参数。最后一个参数可以传入任意多个值

func YourName(name string, aliases ...string) {
	if len(aliases) > 0 {
		println(aliases[0])
	}
}

func YourNameInvoke() {
	YourName("张三")
	YourName("李四", "李五")
	YourName("王五", "王六")
}
