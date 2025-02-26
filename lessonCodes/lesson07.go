package main

import "fmt"

/**
 * @date:       2025/2/26 14:09
 * @author:     cola1213
 * @description:函数
 */
func main() {
	// 调用函数
	sayHi()
	add(1, 2)
	add2(2, 3)
	add3(4, 5, 6, 7, 8, 9)
}

// 多个参数
func add3(numList ...int) {
	fmt.Println(numList)
}

// 如果参数类型一样，可以省略
func add2(n1, n2 int) {
	fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
}

// 无返回值，接受传入的参数
func add(n1 int, n2 int) {
	fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
}

// 函数是一段封装了特定功能的可重用代码块，用于执行特定的任务或计算
// 函数接受输入（参数）并产生输出（返回值）
func sayHi() {
	fmt.Println("hello world, hi ")
}
