/*1.
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
不曾使用的常量，在编译的时候，是不会报错的
显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值
const identifier [type] = value
显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
*/
package main

import "fmt"

func test1() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)
}
/*2.
常量可以作为枚举，常量组
const (
    Unknown = 0
    Female = 1
    Male = 2
)
常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
*/

func test2() {
	const (
        x uint16 = 16
        y
        s = "abc"
        z
    )
	//按照结果来看，%T是类型，%v暂时没看懂
    fmt.Printf("%T,%v\n", y, y) //uint16,16
    fmt.Printf("%T,%v\n", z, z) //string,abc

}

/*3.
iota，特殊常量，可以认为是一个可以被编译器修改的常量
在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
iota 可以被用作枚举值：
const (
    a = iota
    b = iota
    c = iota
)
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
const (
    a = iota
    b
    c
)
*/
func test3() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i) //0 1 2 ha ha 100 100 7 8
}


func main() {
	//test1();
	//test2();
	test3();
}