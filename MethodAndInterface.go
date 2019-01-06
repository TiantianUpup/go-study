package main

import "fmt"

func main() {
	//方法测试
	o := Orders{20, 2}
	//fmt.Println(o.functionTest())
	//等用于
	//fmt.Println(functionTest2(o))
	o.functionTest4()
	fmt.Println(o.price, o.number)
	o.functionTest3()
	fmt.Println(o.price, o.number)
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
