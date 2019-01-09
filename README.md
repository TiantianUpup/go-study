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

### 项目结构说明
项目结构比较简单  
Basic.go：基础部分的测试代码
MethodAndInterface.go：方法和接口部分的测试代码
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
var paramName [count]type
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
切片是不定长的数组  
语法：
```
var paramName []type 
```
和数组的语法相比它不需要事先指定长度  
切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔  
比如：
```
a[low: high]
```
取得是数组a下标范围为[low, high)的元素  
举个例子：
```
func sliceTest() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[0: 3]
	fmt.Println(slice)
}
```
取得是intArr数组中的0、1、2的值  
注意：切片是数组的引用，修改切片的值会影响到数组中值  
举个例子：
```
func sliceTest2() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[0: 3]
	slice[0] = -1
	fmt.Println(slice)
	fmt.Println(intArr)
}
```
对slice[0]赋值为-1后，打印intArr数组发现下标为0的数组元素变为-1
- 切片的默认行为
切片下界的默认值为0，上界则是该切片的长度  
举个例子：
```
func sliceTest3() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[: 3]
	fmt.Println(slice)
	slice = intArr[1:]
	fmt.Println(slice)
}
```
第一次slice赋值相当于`slice := intArr[: 3]`  
第二次slice赋值相当于`slice := intArr[1: 5]`
- 切片的长度和容量
切片的长度就是它所包含的元素个数  
切片的容量是从它的第一个元素开始数，到其**底层数组元素**末尾的个数  
切片s的长度和容量可通过表达式`len(s)`和`cap(s)`来获取  
举个例子：
```
func sliceTest4() {
	intArr := [5]int{0, 1, 2, 3, 4}
	slice := intArr[: 3]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
```
len为3，因为取得是`intArr`数组下标为[0, 3)的数组，数量为3  
从0到4共有5个元素，所以cap为5
- nil切片  
切片的零值是 nil  
nil切片的长度和容量为0且没有底层数组  
举个例子：
```
func sliceTest5()  {
	var slice []int
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	if slice == nil {
		fmt.Println("nil slice")
	}
}
```
- 用make函数创建切片  
make函数会分配一个元素为零值的数组并**返回一个引用了它的切片**
举个例子
```
a := make([]int, 5)  // len(a)=5
```
要指定它的容量，需向make传入第三个参数
```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```
make函数的第一个参数表示底层数组的类型，第二个参数表示切片的长度，第三个参数表示切片的容量
- 向切片追加元素  
通过append函数向切片追加元素  
append函数介绍：  
```
func append(s []T, vs ...T) []T
```
返回值：[]T。结果是一个包含原切片所有元素加上新添加元素的切片。当s的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。  
切片：T类型的切片S  
追加的值：vs ... T，可以一次性追加多个元素  
举个例子：
```
func sliceTest7()  {
	//申明一个切片
	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5)
    fmt.Println(slice)
}
```
一次性往slice切片中添加了1、2、3、4、5 总共5个元素
##### Range
for循环的range形式可遍历切片或映射  
当使用for循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本
举个例子：
```
func rangeTest() {
	//申明一个切片
	slice := []int{1, 2, 3, 4, 5}
	for i, value := range slice {
		fmt.Println(i, value)
	}
}
```
因为range遍历切片会返回两个值，所以用两个变量i、value分别接收了  
注意：可以将下标或值赋予 _ 来忽略它  
举例说明：
```
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
```
##### 映射
存储键值对  
基本语法：
```
var maps = map[keyType]valueType{
    init Value
}
```
举例说明：
    ```
    func mapTest() {
        intStringMap := map[int]string {
            1: "a",
            2: "b",
        }
    
        fmt.Println(intStringMap[1])
    }
    ```
- 用make函数创建映射  
make函数会返回给定类型的映射，并将其初始化备用  
    基本语法：
    ```
    var maps = make(map[int]string)
    ```
举例说明:
```
func mapTest1() {
    maps := make(map[int]string)
    maps[0] = "a"
    fmt.Println(maps[0])
}
```
- 修改map
    - 删除
    ```
    delete(maps, key)
    ```
    - 查找
    ```
    maps[key]
    ```
    - 修改
    ```
    maps[key] = newValue
    ```
    - 双赋值检测某个键是否存在
    ```
    elem, ok = m[key]
    ```
    若key在m中，ok为true，否则ok为false  
    若key不在映射中，那么elem是该映射元素类型的零值  
    同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值  

举例说明：
```
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
```
##### 函数值
函数也是值, 它们可以像其它值一样传递  
函数值可以用作函数的参数或返回值  
java中函数是不能作为参数的，这和java很不一样  
举例说明：
```
func funcParamTest(funcParam func(string) string) string {
	return funcParam("hello world")
}
```
- 闭包
Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量  
该函数可以访问并赋予其引用的变量的值  
换句话说，该函数被“绑定”在了这些变量上。
### Go方法和接口
##### 方法
方法就是一类**带特殊的接收者参数**的函数  
方法接收者在它自己的参数列表内，位于func关键字和方法名之间  
语法：
```
func receiverType functionName() returnType {
}
```
举例说明：
```
func (o Orders) functionTest() int {
	return (o.price) * (o.number)
}
```
等同于下面的函数：
```
func functionTest2(o Orders) int {
	return (o.price) * (o.number)
}
```
测试方法：
```
o := Orders{20, 2}
	fmt.Println(o.functionTest())
	//等用于
	fmt.Println(functionTest2(o))
```
觉得接收者参数就是方法参数  
- 接受者参数可以为指针
举例说明
    ```
    func (o *Orders)functionTest3() {
        //将接受者参数扩大十倍
        o.price *= 10
        o.number *= 10
    }
    ```
    指针接收者的方法可以修改接收者指向的值  
  
    如果去掉`*`  
    举例说明：
    ```
    func (o Orders)functionTest4() {
        //扩大十倍
        o.price *= 10
        o.number *= 10
    }

    ```
    接受者参数的值不会改变，因为函数体里面操作的是接受值的拷贝
- 使用指针作为接受者需要注意的点
    - 以指针为接收者的方法被调用时，接收者既能为值又能为指针，自动加上`&`
    - 以值为接收者的方法被调用时，接收者既能为值又能为指针，自动加上`*`
- 为什么使用指针作为接受者
    - 方法能够修改其接收者指向的值
    - 可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效
#### 接口   
接口是一组方法签名定义的集合  
语法：
```
struct interfaceName interface {
    function1()
    ...
    functionn()
}
```
接口的实现和java中非常不一样，go中自我理解是通过接受者参数实现
举例说明：  
接口定义：
```
type print interface {
	printParam()
}

```
print接口里定义了一个printParam的方法需要去实现  
接口实现：
```
func (book Books) printParam()  {
	fmt.Println(book.name)
}
```
由接受者参数Book实现了printParam方法  
注意实现方法的接受者参数不能是基本类型如int、float32这些
- 接口值  
接口值可以看做包含值和具体类型的元组  
值：实现接口的接受者参数的值，可以通过%v获取  
类型：实现接口的接受者参数的类型，可以通过%T获取  
举例说明：
    ```
    func printInterfaec(p Print)  {
        fmt.Printf("%v, %T", p, p)
    }
    ```
- 空接口  
接口中没有定义任何方法  
空接口可保存任何类型的值，因此可以被用来处理未知类型的值  
举例说明：
    ```
    /*
        空接口定义
    */
    type emptyInterface interface {
    
    }
    
    /*
        打印空接口值
    */
    func printEmptyInterface(i interface{})  {
        fmt.Printf("%v, %T\n", i, i)
    }
    
    /*
        空接口测试
    */
    var i interface{} = 44
    printEmptyInterface(i)
    i = "hello world"
    printEmptyInterface(i)
    ```
### Go并发