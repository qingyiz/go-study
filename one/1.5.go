/*1.
基本数据类型
1.bool
2.Number types:
	int8,int16,int32,int64,int
	uint8,uint16,uint32,uint64,uint
	float32,float64
	complex64,complex128
	byte
	rune
3.string

1.布尔型bool
布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true

2.数值型
2.1 整形
int8
有符号 8 位整型 (-128 到 127)
int16
有符号 16 位整型 (-32768 到 32767)
int32
有符号 32 位整型 (-2147483648 到 2147483647)
int64
有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
uint8
无符号 8 位整型 (0 到 255)
uint16
无符号 16 位整型 (0 到 65535)
uint32
无符号 32 位整型 (0 到 4294967295)
uint64
无符号 64 位整型 (0 到 18446744073709551615)
int和uint:根据底层平台，表示32或64位整数。除非需要使用特定大小的整数，否则通常应该使用int来表示整数。
大小:32位系统32位，64位系统64位。
范围:-2147483648到2147483647的32位系统和-9223372036854775808到9223372036854775807的64位系统。
2.2浮点型
float32
IEEE-754 32位浮点型数
float64
IEEE-754 64位浮点型数
complex64
32 位实数和虚数
complex128
64 位实数和虚数

2.3其他
byte
类似 uint8
rune
类似 int32
uint
32 或 64 位
int
与 uint 一样大小
uintptr
无符号整型，用于存放一个指针
3 字符串型
字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本
    var str string
    str = "Hello World"
4 派生类型(复合类型)
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) 通道类型(Channel )
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map类型
*/