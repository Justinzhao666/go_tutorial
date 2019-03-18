// go的数据类型

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 基本数据类型
	var t = 100
	// 获取变量类型和字节大小
	fmt.Printf("t's type is: %T; \nt's size is: %d; \n", t, unsafe.Sizeof(t))

	// 整形---------------------------------------------
	//有符号整数
	var i int = 1
	var i1 int8 = -128 // -128~127  1byte=8位
	var i2 int16 = 2   // 2byte
	var i3 int32 = 2   // 4byte
	var i4 int64 = 2   // 8byte
	fmt.Println(i, i1, i2, i3, i4)

	//无符号整数
	var j uint = 1
	var j1 uint8 = 1 // 0-255 没有符号
	var j2 uint16 = 1
	var j3 uint32 = 1
	var j4 uint64 = 1
	var j5 byte = 1
	fmt.Println(j, j1, j2, j3, j4, j5)

	var r rune = 12        //等价int32 表示一个Unicode码
	var ch rune = '\u8d75' // 好像也不能表示中文
	fmt.Println(ch)
	fmt.Println(r)

	//存字符用byte，等价unit8，表示的数字范围也一样，只是他还可以接收字符而已
	var c byte = 'c'
	fmt.Println(c)

	// 浮点型---------------------------------------------
	// 浮点数都是有符号的，所以第一位都是符号位
	var k float32 = -123.00009  //单精度
	var k1 float64 = -123.00009 //双精度 - 推荐直接使用
	fmt.Println(k, k1)          // 两者所能表示的精度不一样，小的精度可能丢失
	var k2 = .1                 // 等价于 0.1
	fmt.Printf("%T", k2)        // 给一个小数，go会默认为float64
	//支持科学计数法
	var k3 = 1.2345e2
	fmt.Println(k3)

	var b bool = true
	fmt.Println(b)

	var s string = "abc"
	fmt.Println(s)

	//复杂数据类型

}
