// 文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 程序中文件是以流的形式来操作的
// 写入内存叫做输入流，内存写入文件叫做输出流
// go操作文件使用os.File，其中File是一个结构体

func main() {

	//#1 打开文件
	// open 会返回一个file对象，和一个错误信息
	// file 还可以叫file指针，文件句柄等
	file, e := os.Open("./41_file")
	if e != nil {
		fmt.Println("open file error:", e)
	}
	fmt.Println(file)
	defer file.Close()
	//#2 读取文件(带缓存方式) - 返回一个reader
	reader := bufio.NewReader(file)
	//交由reader去读取文件
	for {
		s, e := reader.ReadString('\n') // 读取到换行就结束
		fmt.Print(s)
		if e == io.EOF { // 表示文件读取到末尾了
			break
		}

	}
	fmt.Println("\n----------------一次性读取文件----------------")
	// 一次性读取文件, 这种没有显示的open操作就不要我们去close文件
	bytes, e := ioutil.ReadFile("./41_file")
	if e != nil {
		fmt.Println("read all error:", e)
	}
	fmt.Printf("%v", string(bytes))

	//#3 关闭文件

	//e = file.Close()
	//if e != nil {
	//	fmt.Println("file close error:", e)
	//}

	writeFile()

}

func writeFile() {
	// 写文件  文件路径，打开方式，linux权限
	file, e := os.OpenFile("41_writeFile", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("write file error:", e)
		return
	}
	defer file.Close()
	str := "justin write this"
	writer := bufio.NewWriter(file) // 带缓存的写入
	writer.WriteString(str)         // 这一步其实是先将文件写入缓存的
	writer.Flush()                  // 将缓存刷新到文件中
}
