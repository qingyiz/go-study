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
//九九乘法表
func test2(){
	for i:=1;i<10;i++{
		for j := 1;j<=i;j++{
			fmt.Printf("%d*%d= %d ",i,j,i*j)
		}
		fmt.Println();
	}
}
//水仙花数 
/*
水仙花数：三位数：[100,999]
每个位上的数字的立方和，刚好等于该数字本身。那么就叫水仙花数
*/
func test3(){
	for i := 100;i < 1000; i++{
		a := i/100
		b := (i/10)%10
		c := i%10
		if(a*a*a + b*b*b + c*c*c == i){
			fmt.Print(i);
			fmt.Printf(" ");
		}
	}
	fmt.Println();

}

//百元百鸡：一百元钱，买100只鸡。公鸡：每只5元，母鸡：每只3元，小鸡：1元3个。有多少种买法？

func test4(){
	var count int = 0
	for i := 0;i < 20;i++{
		for j := 0;j < 33;j++{
			k := 100-i-j
				if 5*i+3*j+(k/3) == 100 && k%3 == 0{
					fmt.Printf("公鸡:%v,母鸡:%v,小鸡:%v",i,j,k)
					fmt.Println()
					count ++
				}

		}
		
	}
	fmt.Println(count);
}
//打印出2-100之间的素数
//素数：也叫质数，只能被1和本身整除的数。
func test5(){
	for i:=2;i<=100;i++{
        flag := true //标记i是否是素数
        //循环验证是否是素数
        for j := 2; j < i; j++ {
            if i%j == 0 {
                flag = false //不是素数
                break
            }
        }
        //打印结果
        if flag == true{
            fmt.Println(i,"是素数")
        }
    }
}

func main()  {
	test1()
	test2()
	test3()
	test4()
	test5()
}