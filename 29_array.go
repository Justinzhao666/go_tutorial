// 数组
package main

import "fmt"

func main() {

	//1. 初始化数组
	var arr1 [2]int = [2]int{1, 2}
	var arr2 = [2]int{1, 2}
	var arr3 = [...]int{1, 2, 3}               //... 让系统自己去判断大小
	var arr4 = [3]string{1: "tom", 0: "jerry"} // 初始化元素指定下标
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	//2. 数组的遍历
	// 除了for取元素下标外：for-range遍历
	for index, value := range arr4 {
		fmt.Printf("idx=%v,val=%v \n", index, value)
	}

	//注意事项：
	// 数组大小固定不可变，申明时确定
	// 这是一个切片 []里面没有大小
	var arr []int
	fmt.Println(arr)
	// 数组是值类型，需要修改需要使用地址
	swap(&arr2)
	fmt.Println(arr2)

}

// 数组的长度是数组类型的一部分： 传参的必须长度也是相同的 [...] 这种长度也必须是相同的。
func swap(arr *[2]int) {
	var item = (*arr)[0] //(*arr) 获取到数组
	(*arr)[0] = (*arr)[1]
	(*arr)[1] = item
}

func demo() {
	// Go中数组是值类型的
	var arr [6]float64 // 定义数组：初始值都是0，字符的就是""，boolean就是false
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr) //数组变量指向的是元素的是数组的首地址
	fmt.Println(&arr[0])     //数组变量指向的是元素的是数组的首地址
	arr[0] = 1.0
	arr[1] = 2.0
	arr[2] = 3.0
	arr[3] = 4.0
	arr[4] = 5.0
	arr[5] = 6.0
	fmt.Println(arr)
	var sum = 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[0]
	}
	var avg = sum / float64(len(arr)) // 这里直接除以数字是可以的，因为数字会自己转换类型，但是传给变量，go应该是没有拆箱装箱自动转换的
	fmt.Println("sum=", sum)
	fmt.Println("avg=", avg)
}
