package main

import "fmt"

/**
 * @date:       2025/2/25 16:21
 * @author:     cola1213
 * @description:数组
 */
func main() {
	// 数组在声明是长度只能是一个常量，不能是变量
	var array = [5]int{1, 2, 4, 3, 5}
	fmt.Println(array)

	// 数组的使用
	fmt.Println(array[0])
	array[0] = 0
	fmt.Println(array[0])
	fmt.Println(len(array))
}
