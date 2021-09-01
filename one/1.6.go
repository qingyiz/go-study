package main

import "fmt"
/*
if 布尔表达式 {
}
*/
func test1() {
   var a int = 10

   if a < 20 {
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}

/*
如果其中包含一个可选的语句组件(在评估条件之前执行)，则还有一个变体。它的语法是
if statement; condition {  
}

if condition{
}
*/

func test2() {
	if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even") 
    }  else {
        fmt.Println(num,"is odd")
    }
}

func main() {
	test1()
	test2()
}