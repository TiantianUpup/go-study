# Go语言学习笔记
### 为什么想学习Go语言

### Go简单介绍
Go（又称Golang[3]）是Google开发的一种**静态强类型、编译型、并发型**，并**具有垃圾回收功能**的编程语言。Go的语法接近C语言，但是因为之前是学习java语言的，所以学习的过程中会将go和java进行比较

摘自维基百科

### Go环境安装
- 编辑器
推荐GoLand  
附：[下载地址](https://www.jetbrains.com/go/download/#section=windows)
- go sdk安装  
附：[go sdk下载地址](https://golang.org/dl/)  
windows下如何配置go环境变量，可以[参考这篇文章](https://blog.csdn.net/defonds/article/details/50538077)

### Go基础
###### 程序入口
每个程序都有包组成，test.go中，包为main，程序由main包开始运行，有点类似于java中的`public static void main(String[] args)`，但是在go中为`func main()`
###### 基本类型介绍
基本类型有
```
bool(布尔类型)
string(字符串类型)
int  int8  int16  int32(别名rune)  int64
uint uint8(别名byte) uint16 uint32 uint64 uintptr

float32 float64

complex64 complex128
```

###### 变量
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

###### 常量
用const申明一个常量，举例说明
```
const str = "hello"
```
申明了一个字符串常量，其值为hello

###### 函数使用语法
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
