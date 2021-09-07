/*分支循环练习题*/
package main

import ("fmt" 
		"math/rand"
		"time"
)


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
//先由系统产生一个随机数，介于1-100之间。用户键盘输入正数来猜测数字，比较用户输入的整数和产生的随机数，给与用于提示：猜大了还是猜小了，用于继续猜测，直到猜对为止。

func test6(){
	rand.Seed(time.Now().UnixNano())
    num := rand.Intn(100)+1
	a := 0
	s := "no"
	for i:=1;;i++{
		//1.读取键盘输入
		fmt.Printf("请输入第 %d 次的猜测的数字：",i)
		fmt.Scanln(&a)
		if a > num{
			fmt.Println("猜大了")
		}else if a < num {
			fmt.Println("猜小了")
		}else{
			fmt.Println("猜对了")
		}
		
		if i == 10{
			fmt.Println("你已经没有机会了，是否重新开始游戏？（yes/no）")
			fmt.Scanln(&s)
			if s == "yes" {
				s = "no"
				i = 0;
			}else{
				break
			}
		}else{
			fmt.Printf("你还有%d次机会\n",10 - i)
		}
	}
}

//打印菱形

func test7(){
	var n int
    fmt.Println("请输入菱形的边长：")
    fmt.Scanln(&n)
	for i := 1;i <= n;i++{
		for j := 0; j < n-i;j++{
			fmt.Print(" ")
		}
		for j := 0;j < 2*i -1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1;i <= n;i++{
		for j := 0; j < i;j++{
			fmt.Print(" ")
		}
		for j := 0 ;j < 2*n-1-2*i ;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少，以12个月为例。
//斐波拉契数列问题
func test8(){
	a := 1
	b := 0
	c := 0
	for i:=0;i<12;i++{
		c = a+b
		a = b
		b = c
		fmt.Println(c)
	}
	
}

//猴子分桃：海滩上有一堆桃子，五只猴子来分。第一只猴子把这堆桃子
//平均分为五份，多了一个，这只猴子把多的一个扔入海中，拿走了一份。
//第二只猴子把剩下的桃子又平均分成五份，又多了一个，它同样把多的一个扔入海中，
//拿走了一份，第三、第四、第五只猴子都是这样做的，问海滩上原来最少有多少个桃子？
func test9() {
/*
    猴子分桃：
     */
	 for i := 6; i <= 10000; i += 5 {
        s := i
        count := 0

        for (s - 1) % 5 == 0 {
            count++
            if count == 5 {
                fmt.Println("满足条件的桃子数：", i)
                break
            }
            s = (s - 1) / 5 * 4

        }
    }
}
//打印万年历
//根据输入的年份，月份打印这个月的日历。
func test10(){
	/*
    万年历
    思路：year:=2018,month:=8
    求出2018年7月31日，距离1900年1月1日的总天数，除以7取余数，就是空格的数量
    星期一星期二星期三星期四星期五星期六星期日
     */

    //1.给定年份和月份，可以键盘输入：
    var year int
    fmt.Print("请输入年份：")
    fmt.Scanf("%d",&year)
    var month int
    fmt.Print("请输入月份：")
    fmt.Scanf("%d",&month)
    //2.求出2018年7月31日到1900年1月1日的总天数
    sum := 0 //总天数
    //2.1先累加整年的天数
    for i := 1900; i < year; i++ {
		if i%4 == 0&& i%100 != 0 || i%400 == 0{
			sum += 366
		}else{
			sum += 365
		}
	}
	//2.2求2018年1月到7月的总天数
    day := 0 //每个月的天数
	for i := 1;i < month;i++{
		switch i {
		case 1,3,5,7,8,10,12:
			day = 31
		case 4,6,9,11:
			day = 30
		case 2:
			if i%4 == 0&& i%100 != 0 || i%400 == 0{
				day = 29
			}else{
				day = 28
			}
		}
		sum += day
	}
	//3.计算空格的数量
    kong := sum % 7
	//4.求month月的天数
    day2 :=0//month月的天数
    switch month {
    case 1,3,5,7,8,10,12:
        day2=31
    case 4,6,9,11:
        day2=30
    case 2:
        if year %4==0&&year % 100!=0||year % 400 ==0{
            day2 = 29
        }else {
            day2=28
        }
    }
    //打印
    fmt.Println("一\t二\t三\t四\t五\t六\t日")
    //打印空格
    for i:=0;i<kong;i++{
        fmt.Print("\t")
    }
    //打印数字
    for i:=1;i<=day2;i++{
        fmt.Print(i,"\t")
        //换行
        if (i+kong) % 7 ==0{
            fmt.Println()
        }
    }
	fmt.Println()
}


func main()  {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	test10()
}