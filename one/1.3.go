/*空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃
_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值*/
package main

import (
    "fmt"
)
func main() {
	var x int //单变量定义 var name type 
	var y,z int //多变量定义 var name1, name2, name3 type
	_,x = 4,5
	fmt.Println(y,z)//变量定义要使用
    fmt.Println(x)
}
/*
1.变量必须先声明，才能够使用，而且每个变量只能被声明一次。
2.因为go是强类型语言，赋值类型要对应
3.name := value，这种声明变量的方式，不能在函数外部使用
4.默认值：也叫零值。
*/