package main

import "fmt"

/**
 * @date:       2025/2/26 10:23
 * @author:     cola1213
 * @description:for语句，例如求1+2+...+100的和
 */
func main() {
	// for 循环
	var result = 0
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)
}
