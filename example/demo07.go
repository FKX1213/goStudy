package example

import "fmt"

/*
*
数组
*/

func Demo0701() {
	// 正确示例
	var nums [5]int

	fmt.Println(nums[0])
	nums[0] = 1
	// 数组元素的数量
	fmt.Println(len(nums))
	// 数组容量
	fmt.Println(cap(nums))
}

func Demo0702() {
	nums := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
	nums1 := nums[1:]
	for i := 0; i < len(nums1); i++ {
		fmt.Printf("%d ", nums1[i])
	}
	fmt.Println()
}

func Demo0703() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	slice[0] = 0
	fmt.Printf("array: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)
}

/**
切片
*/
