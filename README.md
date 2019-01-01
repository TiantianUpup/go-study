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
