package main

import "fmt"

/**
 * @date:       2025/2/24 9:13
 * @author:     cola1213
 * @description:变量定义
 */
func main() {
	var name1, name2, name3 string = "cola", "cola2", "cola3"
	fmt.Println(name1, name2, name3)

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

	b1, b2 := 1, 2
	fmt.Println(b1, b2)

	fmt.Println(c1, c2)
}

var (
	c1 = 1
	c2 = 2
)
