package main

import (
    "fmt"
)
/* := 初始化声明，相当于C语言中的定义吧，当出现第二次:=时将退化为赋值=操作
    但这个前提是，最少要有一个新的变量被定义，且在同一作用域*/
func main() {
    x := 140
    fmt.Println(&x)
    x, y := 200, "abc"
    fmt.Println(&x, x)
    fmt.Println(y)
    //x := 300// error 
    fmt.Println(x)
}
