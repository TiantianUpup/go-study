package main

import "fmt"

func main() {
	//fmt.Println("Hello World")
	//fmt.Println(sayHello())

	//测试add函数
	//fmt.Println(add(1, 2))

	//测试returnMulti多值返回函数
	//fmt.Println(returnMulti(1, "hello"))

	//测试指定返回值的名
	fmt.Println(defineParamName(1, "hello"))
}

/*
   函数语法测试
*/
func sayHello() string {
	return "hello"
}

/*
	相同参数类型最后一个可以省略
*/
func add(x, y int) int {
	return x + y
}

/*
	多值返回
*/
func returnMulti(x int, y string) (int, string) {
	var i = x * 2
	return i, y
}

/*
	指定返回值的名称
*/
func defineParamName(x int, y string) (i int, str string) {
	i = x
	str = y
	return
}
