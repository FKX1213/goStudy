package main

import (
	"fmt"
	"sort"
)

/**
 * @date:       2025/2/25 16:21
 * @author:     cola1213
 * @description:数组、切片
 */
func main() {
	// 切片
	var nums []int
	nums = append(nums, 1)
	nums = append(nums, 3)

	fmt.Println(nums[0])
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// 定义一个字符串切片
	var list []string
	fmt.Println(list == nil) // true

	// 定义一个字符串切片
	var list1 = make([]string, 0)
	fmt.Println(list1, len(list1), cap(list1))
	fmt.Println(list1 == nil) // false

	list2 := make([]int, 2, 2)
	fmt.Println(list2, len(list2), cap(list2))

	// 切片是数组切出来的
	var nums01 = [3]int{1, 2, 3}
	nums02 := nums01[0:2] // 1,2
	fmt.Println(nums02)

	// 切片排序
	nums03 := []int{1, 2, 7, 3, 8, 6, 9}
	sort.Ints(nums03) // 默认是从小到大
	fmt.Println(nums03)
	sort.Sort(sort.Reverse(sort.IntSlice(nums03))) // 反转，从大到小
	fmt.Println(nums03)
}
