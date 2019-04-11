// channel
package main

import (
	"fmt"
	"time"
)

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
	fmt.Println(intChan) // 是引用类型

	// 写入管道
	intChan <- 10
	intChan <- 11
	intChan <- 12
	//intChan <- 13 管道不能超出定义的容量
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	// 管道取出
	var num = <-intChan
	fmt.Println(num)
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	num = <-intChan
	//num = <- intChan
	//num = <- intChan  没有数据的时候再取报错 deadlock

	// 关闭管道
	// 关闭后不能再从管道写数据了，但是可以读数据
	intChan <- 10
	intChan <- 11
	close(intChan)
	//intChan <- 13  panic: send on closed channel
	fmt.Println(<-intChan)

	// 遍历管道
	// 使用for-range,不可以for循环，应为队列是在不断变的
	// 遍历如果管道没有close，遍历完后报deadlock错误
	// 遍历如果管道close了，正常遍历后退出遍历
	for v := range intChan {
		fmt.Print(" ", v)
	}
	fmt.Println()
	demoChannel()

}

func writeChannel(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readChannel(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read", v)
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func demoChannel() {
	// 代码一环套一环的： read 套 write的close channel才会退出，main 又套在read 结束后close channel才结束
	var intChan = make(chan int, 50) // go的读写频率不一致不会有问题，底层有处理
	var exitChan = make(chan bool, 1) // 单独定义一个管道判断是否诚信读取完成， 其实一个标志位也可以啊？
	go writeChannel(intChan)
	go readChannel(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}

	}
}
