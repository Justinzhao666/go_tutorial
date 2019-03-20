package main

import (
	"../model" // 相对路径或者GoPath相对路径或者绝对路径都可以
	"fmt"
)

func main() {

	fmt.Println(model.Public_val)
	//fmt.Println(model.private_val) 无法访问私有变量
}
