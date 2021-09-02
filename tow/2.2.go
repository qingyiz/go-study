/*分支循环练习题*/
package main

import "fmt"


//1 交换两个数

func test1()  {
    /*
    9.交换2个变量的值。
	a := 3
	b := 2
    */
	a := 3
	b := 2
	//方法一：借助第三个变量
	temp := a
	a = b
	b = temp
	fmt.Println(a, b)
	//方法二：同时交换
	a = 3
	b = 2
	a, b = b, a
	fmt.Println(a, b)
	//方法三：求和
	a = 3
	b = 2
	a = a + b // a=5 b=2
	b = a - b //a=5 b=3
	a = a - b // a2,b=3
	fmt.Println(a, b)
	//方法四：
	a = 3
	b = 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
	
}

func main()  {
	test1()
}