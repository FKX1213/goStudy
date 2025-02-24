package main

import "fmt"

/**
 * @date:       2025/2/24 9:13
 * @author:     cola1213
 * @description:变量定义
 */
// main 包才能执行函数
func main() {
	// 先定义，再赋值
	var name string
	name = "coco cola" // go 中，声明的变量，必须使用，否则报错：Unused variable 'name'
	fmt.Println(name)

	// 声明变量的时候就赋值
	var name1 string = "pepsi cola"
	fmt.Println(name1)

	// 但是其实，声明的时候，可以省略类型，由赋的值直接推断是什么类型
	var name2 = "diet coke"
	fmt.Println(name2)

	// 短声明符号赋值
	name3 := "zero coke"
	fmt.Println(name3)

	// 零值
	var name4 string
	fmt.Println(name4 + "==")

	fmt.Println(age)
}

var age = 12

func hello() {
	fmt.Println("hello")
	fmt.Println(age)
}
