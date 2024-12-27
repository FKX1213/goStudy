package main

import "fmt"

func main() {
	// 方式一：声明一个变量，默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("tyoe of a = %T\n", a)

	// 方式二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("tyoe of b = %T\n", b)

	// 方法三：初始化的时候, 省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("tyoe of c = %T\n", c)

	// 方法四：(常用) 省去var关键字，直接自动匹配
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("tyoe of d = %T\n", d)
}
