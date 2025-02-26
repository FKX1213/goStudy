package main

import "fmt"

/**
 * @date:       2025/2/26 9:56
 * @author:     cola1213
 * @description:判断语句
 */
func main05() {
	fmt.Println(1 == 1 && 1 == 0) //false
	fmt.Println(1 == 1 && 1 == 1) //true

	fmt.Println("================")
	fmt.Println(1 == 1 || 1 == 0) //true
	fmt.Println(1 == 0 || 1 == 0) //false

	fmt.Println("================")
	fmt.Println(!(1 == 1)) // false
	fmt.Println(!(1 == 0)) // true

}
