//go变量的使用

package main

import "fmt"

//全局变量
var (
	n1 = 100
	n2 = 200
	n3 = "hello"
)

func main() {
	//申明方式1：定义变量，如果不赋值就使用默认值，int 默认值为0 string 为空串
	var i int
	//变量赋值
	i = 10
	var j int = 10
	//变量使用
	fmt.Println("i=", i)
	fmt.Println("j=", j)

	//申明方式2：类型推导，自动识别赋值的类型
	var num = 10.11
	fmt.Println("num=", num)

	//申明方式3：使用:= 来进行申明变量并赋值
	name := "justin"
	fmt.Println("name=", name)

	// ----------------------------------

	// go一次申明定义多个变量
	var a, b, c int
	fmt.Println(a, b, c)

	var d, e, f = 1, 2, 3
	fmt.Println(d, e, f)

	j, h, k := 1, 2, 3
	fmt.Println(j, h, k)

	fmt.Println(n1, n2, n3)

}
