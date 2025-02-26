package main

import "fmt"

/**
 * @date:       2025/2/26 9:56
 * @author:     cola1213
 * @description:判断语句
 */
func main() {
	var age int
	fmt.Println("请输入你的年龄:")

	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}

}
