/*
1.声明和初始化数组
需要指明数组的大小和存储的数据类型。
var variable_name [SIZE] variable_type
2.示例代码：
var balance [10] float32
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
数组的其他创建方式：
  var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
  var a [4] float32 // 等价于：var arr2 = [4]float32{}
  fmt.Println(a) // [0 0 0 0]
  var b = [5] string{"ruby", "王二狗", "rose"}
  fmt.Println(b) // [ruby 王二狗 rose  ]
  var c = [5] int{'A', 'B', 'C', 'D', 'E'} // byte
  fmt.Println(c) // [65 66 67 68 69]
  d := [...] int{1,2,3,4,5}// 根据元素的个数，设置数组的大小
  fmt.Println(d)//[1 2 3 4 5]
  e := [5] int{4: 100} // [0 0 0 0 100]
  fmt.Println(e)
  f := [...] int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
  fmt.Println(f)
*/

/*
1.数组的长度
通过将数组作为参数传递给len函数，可以获得数组的长度。
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is",len(a))

}
*/


package main

import "fmt"


/*
如果您只需要值并希望忽略索引，那么可以通过使用_ blank标识符替换索引来实现这一点。
for _, v := range a { //ignores index  
}
*/
func test1(){
	//使用range遍历数组
	a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}

/*数组是值类型
Go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，
将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。
类似地，当将数组传递给函数作为参数时，它们将通过值传递，而原始数组将保持不变。
*/

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func test2() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}
/*数组的大小是类型的一部分。因此[5]int和[25]int是不同的类型。
因此，数组不能被调整大小。不要担心这个限制，因为切片的存在是为了解决这个问题。
*/

/*切片(Slice)
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

切片是一种方便、灵活且强大的包装器。切片本身没有任何数据。它们只是对现有数组的引用。

切片与数组相比，不需要设定长度，在[]中不用设定值，相对来说比较自由

从概念上面来说slice像一个结构体，这个结构体包含了三个元素：

指针，指向数组中slice指定的开始位置
长度，即slice的长度
最大长度，也就是slice开始位置到数组的最后位置的长度
1.定义切片
var identifier []type
切片不需要说明长度。
或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)

make([]T, length, capacity)
*/






func main() {  
    test1()
	test2()
}