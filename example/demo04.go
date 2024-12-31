package example

import (
	"fmt"
	"os"
)

// Demo0401 可以直接将字符串写入到标准输出中
func Demo0401() {
	os.Stdout.WriteString("hello world!\n")
}

// Demo0402 Go 有两个内置的函数print，println
func Demo0402() {
	print("hello world!\n")
	println("hello world!")
}

// Demo0403 使用fmt包，它提供了fmt.Println函数
func Demo0403() {
	fmt.Println("hello world!")
}

// Demo0404 使用read直接读文件
func Demo0404() {
	var buf [1024]byte            // 创建一个大小为1024字节的缓冲区
	n, _ := os.Stdin.Read(buf[:]) // 从标准输入读取数据到缓冲区 buf，n 是读取的字节数
	os.Stdout.Write(buf[:n])      // 将读取到的数据写入标准输出
}

func Demo0405() {
	fmt.Printf("%%%s\n", "hello world")

	fmt.Printf("%s\n", "hello world")
	fmt.Printf("%q\n", "hello world")
	fmt.Printf("%d\n", 2<<7-1)

	fmt.Printf("%f\n", 1e2)
	fmt.Printf("%e\n", 1e2)
	fmt.Printf("%E\n", 1e2)
	fmt.Printf("%g\n", 1e2)

	fmt.Printf("%b\n", 2<<7-1)
	fmt.Printf("%#b\n", 2<<7-1)
	fmt.Printf("%o\n", 2<<7-1)
	fmt.Printf("%#o\n", 2<<7-1)
	fmt.Printf("%x\n", 2<<7-1)
	fmt.Printf("%#x\n", 2<<7-1)
	fmt.Printf("%X\n", 2<<7-1)
	fmt.Printf("%#X\n", 2<<7-1)

	type person struct {
		name    string
		age     int
		address string
	}
	fmt.Printf("%v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%+v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%#v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%t\n", true)
	fmt.Printf("%T\n", person{})
	fmt.Printf("%c%c\n", 20050, 20051)
	fmt.Printf("%U\n", '码')
	fmt.Printf("%p\n", &person{})
}
