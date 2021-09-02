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

/*
switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break。
而如果switch没有表达式，它会匹配true
Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
*/
func test3(){
    var grade string = "B"
    var marks int = 90

    switch marks {
        case 90: grade = "A"
        case 80: grade = "B"
        case 50,60,70 : grade = "C"  //case 后可以由多个数值
        default: grade = "D"  
    }

    switch {
        case grade == "A" :
            fmt.Printf("优秀!\n" )     
        case grade == "B", grade == "C" :
            fmt.Printf("良好\n" )      
        case grade == "D" :
            fmt.Printf("及格\n" )      
        case grade == "F":
            fmt.Printf("不及格\n" )
        default:
            fmt.Printf("差\n" );
    }
    fmt.Printf("你的等级是 %s\n", grade );   
}
/*
如需贯通后续的case，就添加fallthrough
*/
type data [2]int
func test4(){
    switch x := 5; x {
    default:
        fmt.Println(x)
    case 5:
        x += 10
        fmt.Println(x)
        fallthrough
    case 6:
        x += 20
        fmt.Println(x)
    }
}
func test5(){
    //case中的表达式是可选的，可以省略。如果该表达式被省略，则被认为是switch true，并且每个case表达式都被计算为true，并执行相应的代码块。
    num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }
}


func test6(){
     /*
    switch的其他写法：
    1.省略switch后的表达式，相当于直接作用在true
    2.case后可以同时匹配多个数据,匹配上任意一个都可以执行该case分支
    3.switch后支持多一条初始化语句
     */
     //1.省略switch后的表达式
     switch  { //相当于true
     case true:
         fmt.Println("true分支被执行了。。")
     case false:
         fmt.Println("false分支被执行了。。。")
     default:
         fmt.Println("以上都没有被执行。。。")
     }
     //2.case后有多个数值
     var c string
     c = "a"
     switch c {
     case "i","o","u","e","a":
         fmt.Println("c是一个元音")
     default:
         fmt.Println("c是辅音。。")
     }
 
     //3.多一条初始化语句
     switch lan:="golang"; lan {
     case "java":
         fmt.Println("java语言。。")
     case "c":
         fmt.Println("c语言。。")
     case "python":
         fmt.Println("python语言。。")
     case "golang":
         fmt.Println("go语言。。")
     default:
         fmt.Println("不知道。。")
     }
}
/*
switch的注意事项
case后的常量值不能重复
case后可以有多个常量值
fallthrough应该是某个case的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误。
*/

/*
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    // 你可以定义任意个数的case 
    default: 
       statement(s);
}
*/

func test7() {
    var x interface{}

    switch i := x.(type) {
        case nil:      
            fmt.Println(" x 的类型 :%T",i)                
        case int:      
            fmt.Println("x 是 int 型")                       
        case float64:
            fmt.Println("x 是 float64 型")           
        case func(int) float64:
            fmt.Println("x 是 func(int) 型")                      
        case bool, string:
            fmt.Println("x 是 bool 或 string 型" )       
        default:
            fmt.Println("未知型")    
    }
}
/*
select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
否则：
如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
*/

func test8(){
    var c1, c2, c3 chan int
    var i1, i2 int
    select {
        case i1 = <-c1:
            fmt.Printf("received ", i1, " from c1\n")
        case c2 <- i2:
            fmt.Printf("sent ", i2, " to c2\n")
        case i3, ok := (<-c3):  // same as: i3, ok := <-c3
            if ok {
                fmt.Printf("received ", i3, " from c3\n")
            } else {
                fmt.Printf("c3 is closed\n")
            }
        default:
            fmt.Printf("no communication\n")
    }    
}
/*
循环语句表示条件满足，可以反复的执行某段代码。
for是唯一的循环语句。(Go没有while循环)
for init; condition; post { }
初始化语句只执行一次。在初始化循环之后，将检查该条件。如果条件计算为true，那么{}中的循环体将被执行，然后是post语句。
post语句将在循环的每次成功迭代之后执行。在执行post语句之后，该条件将被重新检查。如果它是正确的，循环将继续执行，否则循环终止。
所有的三个组成部分，即初始化、条件和post都是可选的。
for condition { } //效果与while相似
for { } //效果与for(;;) 一样

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
*/
func test9(){
    var b int = 15
    var a int

    numbers := [6]int{1, 2, 3, 5} 

    /* for 循环 */
    for a := 0; a < 10; a++ {
        fmt.Printf("a 的值为: %d\n", a)
    }
    for a < b {
        a++
        fmt.Printf("a 的值为: %d\n", a)
        }
    for i,x:= range numbers {
        fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
    }   
}

/*
"贴标签"：
break,continue,多层循环嵌套：默认结束是里层的循环。如果要结束外层循环，那么要给循环贴标签，
break和continue可以跳出指定的循环。这里的"标签名"要满足go语言的标识符的命名
*/
func test10(){
    /*
    break,continue,多层循环嵌套：默认结束是里层的循环。
    "贴标签"
        如果要结束外层循环，那么要给循环贴标签，break和continue可以跳出指定的循环。
     */
     out:for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            if j == 2{
                //break out//结束指定的循环，标签名
                continue out
                //continue
            }
            fmt.Println("i,", i, ",j,", j)
        }
    }
}

func main() {
	test1()
	test2()
    test3()
    test4()
    test5()
    test6()
    test7()
    test8()
    test9()
    test10()
}