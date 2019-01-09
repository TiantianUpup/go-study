package main

import (
	"fmt"
)

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

	//defer测试
	//deferTest()

	//结构体访问测试
	//structTest()

	//structTest2()

	//数组测试
	//arrayTest()

	//切片测试
	//sliceTest()
	//sliceTest2()
	//sliceTest3()
	//sliceTest4()
	//sliceTest5()
	//sliceTest6()
	//sliceTest7()

	//range测试
	//rangeTest()
	//rangeTest2()

	//映射测试
	//mapTest()
	//mapTest1()
	//mapTest2()

	//函数测试
	function := func(str string) string {
		return str
	}

	str := funcParamTest(function)
	fmt.Println(str)
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
	fmt.Println("first Print")
	defer fmt.Println("hello")
}

/*
	定义Book的结构体
*/
type Book struct {
	name  string
	price float32
}

/*
	结构体测试1
*/
func structTest() {
	book := Book{"java核心技术", 110.0}
	fmt.Println("name: ", book.name)
	fmt.Println("price: ", book.price)
}

/*
	结构体测试2
*/
func structTest2() {
	book := Book{price: 110.0, name: "java核心技术"}
	fmt.Println("name: ", book.name)
	fmt.Println("price: ", book.price)

	book2 := Book{price: 110.0}
	fmt.Println("name: ", book2.name)
	fmt.Println("price: ", book2.price)

	book3 := Book{}
	fmt.Println("name: ", book3.name)
	fmt.Println("price: ", book3.price)
}

/*
	数组测试
*/
func arrayTest() {
	var intArr [3]int
	intArr[0] = 1
	fmt.Println(intArr[0])

	intArr2 := [3]int{3, 2, 1}
	fmt.Println(intArr2[0])
}

/*
	切片测试(基本语法)
*/
func sliceTest() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[0:3]
	fmt.Println(slice)
}

/*
	切片测试(测试切片是数组的引用)
*/
func sliceTest2() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[0:3]
	slice[0] = -1
	fmt.Println(slice)
	fmt.Println(intArr)
}

/*
	切片测试(测试切片的默认行为)
*/
func sliceTest3() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[:3]
	fmt.Println(slice)
	slice = intArr[1:]
	fmt.Println(slice)
}

/*
	切片测试(测试切片的长度和容量)
*/
func sliceTest4() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[:3]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

/*
	切片测试(测试nil切片)
*/
func sliceTest5() {
	var slice []int
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	if slice == nil {
		fmt.Println("nil slice")
	}
}

/*
	切片测试(make函数创建切片)
*/
func sliceTest6() {
	slice := make([]int, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
}

/*
	切片测试(切片元素的添加)
*/
func sliceTest7() {
	//申明一个切片
	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)
}

/*
	range测试(基本语法)
*/
func rangeTest() {
	//申明一个切片
	slice := []int{1, 2, 3, 4, 5}
	for i, value := range slice {
		fmt.Println(i, value)
	}
}

/*
	range测试(使用_省略)
*/
func rangeTest2() {
	//申明一个切片
	slice := []int{1, 2, 3, 4, 5}
	//省略下标
	for _, value := range slice {
		fmt.Println(value)
	}
	//省略值
	for i := range slice {
		fmt.Println(i)
	}
}

/*
	map测试(基本语法)
*/
func mapTest() {
	intStringMap := map[int]string{
		1: "a",
		2: "b",
	}

	fmt.Println(intStringMap[1])
}

/*
	map测试(测试make函数创建映射)
*/
func mapTest1() {
	maps := make(map[int]string)
	maps[0] = "a"
	fmt.Println(maps[0])
}

/*
	map测试(修改map)
*/
func mapTest2() {
	maps := make(map[int]string)
	maps[0] = "a"
	maps[1] = "b"
	fmt.Println(maps[0])
	//修改
	maps[0] = "c"
	fmt.Println(maps[0])
	//删除
	delete(maps, 0)
	fmt.Println(maps[0])
	//查找
	value, i := maps[0]
	fmt.Println(i, value)

	value, i = maps[1]
	fmt.Println(i, value)
}

/*
	函数值测试
*/
func funcParamTest(funcParam func(string) string) string {
	return funcParam("hello world")
}
