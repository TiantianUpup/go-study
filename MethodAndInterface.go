package main

import "fmt"

func main() {
	//方法测试
	//o := Orders{20, 2}
	//fmt.Println(o.functionTest())
	//等用于
	//fmt.Println(functionTest2(o))
	//o.functionTest4()
	//fmt.Println(o.price, o.number)
	//o.functionTest3()
	//fmt.Println(o.price, o.number)

	//接口测试
	//books := Books{10, "go in action"}
	//books.printParam()
	//接口值测试
	//printInterface(books)

	//空接口测试
	var i interface{} = 44
	printEmptyInterface(i)
	i = "hello world"
	printEmptyInterface(i)
}

/*
	结构体
*/
type Orders struct {
	price  int
	number int
}

/*
	方法测试
*/
func (o Orders) functionTest() int {
	return (o.price) * (o.number)
}

/*
	方法测试二
*/
func functionTest2(o Orders) int {
	return (o.price) * (o.number)
}

/*
	方法测试三
*/
func (o *Orders) functionTest3() {
	//扩大十倍
	o.price *= 10
	o.number *= 10
}

/*
	方法测试四
*/
func (o Orders) functionTest4() {
	//扩大十倍
	o.price *= 10
	o.number *= 10
}

/*
	接口定义
*/
type Print interface {
	printParam()
}

/*
	接口实现类
*/
func (book Books) printParam() {
	fmt.Println(book.name)
}

/*
	结构体定义
*/
type Books struct {
	price int
	name  string
}

func printInterface(p Print) {
	fmt.Printf("%v, %T", p, p)
}

/*
	空接口定义
*/
type emptyInterface interface {
}

func printEmptyInterface(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}
