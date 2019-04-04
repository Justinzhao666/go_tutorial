// 接口
package main

import "fmt"

// Go的核心编程就是接口编程， 接口是多态的前提：前面的继承可以看到——子类的方法只能自己来去调，没有办法由父类来多态的调用

type usb interface {
	// 接口可以定义一组方法，但是不能定义任何变量
	Start()
	Stop()
}

type Camera struct {
}

func (camera *Camera) Start() {
	fmt.Println("camera start")
}
func (camera *Camera) Stop() {
	fmt.Println("camera stop")
}

type Phone struct {
}

func (phone *Phone) Start() {
	fmt.Println("phone start")
}
func (phone *Phone) Stop() {
	fmt.Println("phone stop")
}

type Computer struct {
}

func (computer *Computer) Work(usb2 usb) {
	usb2.Start()
	usb2.Stop()
}

func main() {
	// 接口使用的简单案例：
	var phone = Phone{}
	var camera = Camera{}
	var computer = Computer{}
	computer.Work(&phone)
	computer.Work(&camera)
}
