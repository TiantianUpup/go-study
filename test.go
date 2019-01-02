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
	//fmt.Println(defineParamName(1, "hello"))

	//测试for循环
	//fmt.Println(forTest1())

	//if测试
	//ifTest1(-1)
	//ifTest2(-1)

	//switch测试
	//switchTest("hello")
	//switchTest("hello world")
	//switchTest2()

	deferTest()
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

/*
	go中的循环体
*/
func forTest1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}

/*
	go if测试1
*/
func ifTest1(x int) {
	if x < 0 {
		fmt.Println("x < 0")
	}
}

/*
	go if测试2
*/
func ifTest2(x int) {
	if x := -x; x < 0 {
		fmt.Println("x < 0")
	} else {
		fmt.Println("x > 0")
	}
}

/*
	switch测试
*/
func switchTest(str string) {
	switch str {
	case "hello":
		fmt.Println("hello world")
	case "new":
		fmt.Println("new year")
	default:
		fmt.Println("default")
	}
}

/*
	switch测试二
*/
func switchTest2() {
	switch str := sayHello(); str {
	case "hello":
		fmt.Println("hello world")
	case "new":
		fmt.Println("new year")
	default:
		fmt.Println("default")
	}
}

/*
	defer测试
*/
func deferTest() {
	defer fmt.Println("world")
	fmt.Println("first print")
	defer fmt.Println("hello")
}
