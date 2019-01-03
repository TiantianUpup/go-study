# Go语言学习笔记
### 为什么想学习Go语言

### Go简单介绍
Go（又称Golang）是Google开发的一种**静态强类型、编译型、并发型**，并**具有垃圾回收功能**的编程语言。Go的语法接近C语言，但是因为之前是学习java语言的，所以学习的过程中会将go和java进行比较

摘自维基百科

### Go环境安装
- 编辑器
推荐GoLand  
附：[下载地址](https://www.jetbrains.com/go/download/#section=windows)
- go sdk安装  
附：[go sdk下载地址](https://golang.org/dl/)  
windows下如何配置go环境变量，可以[参考这篇文章](https://blog.csdn.net/defonds/article/details/50538077)

### Go基础
##### 程序入口
每个程序都有包组成，test.go中，包为main，程序由main包开始运行，有点类似于java中的`public static void main(String[] args)`，但是在go中为`func main()`
##### 基本类型介绍
基本类型有
```
bool(布尔类型)
string(字符串类型)
int  int8  int16  int32(别名rune)  int64
uint uint8(别名byte) uint16 uint32 uint64 uintptr

float32 float64

complex64 complex128
```

##### 变量
用var申明一个变量，变量类型写在参数名后面，举例说明：
```
var i int
```
或使用短变量申明
```
i : = 2
```
申明了一个int类型的变量  
注意：  
1.变量可以申明在方法里面，也可以申明在包下  
2.对于已经初始化的变量可以省略变量类型
举例说明：
```
var i = 2
```
申明了一个int型的i变量，其值为2  
3.对于短变量申请只能使用在函数内，并且需要进行初始化
变量组的申明，举例说明：
```
var (
    a, b int = 1, 2
    c, d string = "hello", "world"
)
```

##### 常量
用const申明一个常量，举例说明
```
const str = "hello"
```
申明了一个字符串常量，其值为hello

##### 函数使用语法
```
func func_name(x param_type) return_type {
}
```
用func申明这是一个方法，和java函数语法相比：
- 方法返回值写在方法名后面
- 参数类型写在参数后面

其他不同点：
- 支持多值返回
  举例说明
  ```
  func returnMulti(x int, y string) (int, string) {
    var i = x * 2
    return i, y
  }
  ```
  返回了int型和string型的两个变量
- 相同参数类型，除最后一个参数类型不可以省略，前面的参数类型都可以省略
  举例说明：
  ```
  func add(x, y int) int {
      return x + y
  }
  ```
  x和y都为int类型，省略了x的参数类型
- 返回值命名
  举例说明
  ```
  func defineParamName(x int, y string) (i int, str string) {
      i = x
      str = y
      return
  }
  ```
  返回的为i和str这两个变量

##### 循环体
基本的 for 循环由三部分组成，它们用分号隔开：
- 初始化语句：在第一次迭代前执行
- 条件表达式：在每次迭代前求值
- 后置语句：在每次迭代的结尾执行
其中初始化语句和后置语句可以不写在循环中  
如果一个循环中三部分都没有，即`for {}`为死循环
举例说明
```
func forTest1() int {
	sum := 0
	for i:= 0; i < 10; i++ {
		sum += i
	}

	return sum
}

```
和java相比，循环部分没有() 


##### 条件判断
- if  
条件判断
    - 和for循环一样，表达式外无需小括号()  
    举例说明
        ```
        func ifTest1(x int) {
            if x < 0 {
                fmt.Println("x < 0")
            }
        }
        ```
    -  if语句可以在条件表达式前执行一个简单的语句  
    举例说明
        ```
        func ifTest2(x int) {
            if x := -x; x < 0 {
                fmt.Println("x < 0")
            } else {
                fmt.Println("x > 0")
            }
        }
        ```
- switch  
编写一系列if-else的简单写法  
    举个例子
    ```
    func switchTest(str string)  {
        switch str {
        case "hello":
            fmt.Println("hello world")
        case "new":
            fmt.Println("new year")
        default:
            fmt.Println("default")
        }
    }
    ```
    可以在switch中执行一个简单的短语  
    举个例子
    ```
    func switchTest2()  {
        switch str := sayHello(); str {
        case "hello":
            fmt.Println("hello world")
        case "new":
            fmt.Println("new year")
        default:
            fmt.Println("default")
        }
    }
    ```
    switch中的值最后由sayHello()函数决定  
    与java中的不同点： 
    - switch中的值无需为常量，且取值不必为整数  
    - 不需要break，Go自动提供了在这些语言中每个case后面所需的break语句
    - switch中可以执行一个简单的短语
  
    <!-- 没有条件的switch
    相当于switch true，默认是default -->
- defer  
defer语句会将函数推迟到外层函数返回之后执行。  
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用  
举个例子：
    ```
    func deferTest()  {
        defer fmt.Println("world")
        fmt.Println("first print")
        defer fmt.Println("hello")
    }
    
    ```
    运行结果：
    ```
    func deferTest()  {
        defer fmt.Println("world")
        fmt.Println("first print")
        defer fmt.Println("hello")
    }
    ```
##### 指针
指针保存了值的内存地址  
类型 *T 是指向 T 类型值的指针。其零值为 nil。
```
var p *int
```
& 操作符会生成一个指向其操作数的指针。
```
i := 42
p = &i
```
表示p指向了保存42的内存地址,&只能与操作数一起使用
* 操作符表示指针指向的底层值。
```
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

##### 结构体
结构体是字段的集合  
语法如下：
```
type structName struct {
}
```
有点类似于java中的类，但是没有方法
- 结构体字段的访问  
通过`.`进行访问  
举例说明：
```
type Book struct {
	name string
	price float32
}

func structTest() {
	book := Book{"java核心技术", 110.0}
	fmt.Println("name: ", book.name)
	fmt.Println("price: ", book.price)
}

```
通过`.`访问到`name`和`price`
- 结构体文法  
通过`name:`给结构体的部分字段赋值，赋值顺序与定义的顺序无关  
举个例子：  
```
func structTest2() {
	book := Book{price:110.0, name: "java核心技术"}
	fmt.Println("name: ", book.name)
	fmt.Println("price: ", book.price)

	book2 := Book{price:110.0}
	fmt.Println("name: ", book2.name)
	fmt.Println("price: ", book2.price)

	book3 := Book{}
	fmt.Println("name: ", book3.name)
	fmt.Println("price: ", book3.price)
}
``` 
如果不初始化为类型的默认值，如int为0，string为空
##### 数组  
固定个数的一类数  
语法:  
```
var [count]type
```
type：类型  
count：数组的长度
举个例子
```
func arrayTest() {
	var intArr [3]int
	intArr[0] = 1
	fmt.Println(intArr[0])

	intArr2 := [3]int{3, 2, 1}
	fmt.Println(intArr2[0])
}
```
注意：数组短变量申明不能省略变量类型，intArr2变量`:=`右边说明了变量类型为`[3]int`，长度为3的数组
##### 切片
##### Range

### Go方法和接口   
### Go并发
