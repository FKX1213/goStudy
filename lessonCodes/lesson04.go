package main

import "fmt"

/**
 * @date:       2025/2/25 16:21
 * @author:     cola1213
 * @description:数组、切片、map
 */
func main04() {
	// map，声明
	var map01 map[string]string
	// 初始化
	map01 = make(map[string]string)
	// 设置值
	map01["name"] = "cola"
	fmt.Println(map01) // map[name:cola]
	// 取值
	fmt.Println(map01["name"])
	// 删除值
	delete(map01, "name")
	fmt.Println(map01)

	map01["sex"] = "no"
	map01["name"] = "bob"
	fmt.Println(map01) // map[name:bob sex:no]
}
