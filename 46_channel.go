// channel
package main

import "fmt"

//为什么需要channel
//主线程无法确定是否所有的协程都执行完了

// channel
//本质：就是一个数据结构-队列
//队列本身是线程安全的，多个协程访问的时候不需要加锁
//channel是有类型的，string类型只能放string类型数据 -- 废话
// channel是引用类型，必须初始化make后才能写入数据

func main() {

	// 定义管道
	var intChan chan int // 定义一个int的channel
	intChan = make(chan int, 3)
	fmt.Println(intChan)  // 是引用类型

	// 写入管道
	intChan <- 10
	intChan <- 11
	intChan <- 12
	//intChan <- 13 管道不能超出定义的容量
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	// 管道取出
	var num = <- intChan
	fmt.Println(num)
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	num = <- intChan
	num = <- intChan
	//num = <- intChan  没有数据的时候再取报错 deadlock

}
