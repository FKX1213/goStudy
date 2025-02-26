package main

import (
	"fmt"
	"time"
)

/**
 * @date:       2025/2/26 10:23
 * @author:     cola1213
 * @description:for语句，例如求1+2+...+100的和
 */
func main06() {
	// for 循环
	var result = 0
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)

	// 死循环，每隔1秒打印当前的时间
	for false {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}

	// golang 没有 while 循环，如果需要使用 while 循环，是由 for 变体
	i := 0
	sum := 0
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)

	// do - while
	j := 0
	sum2 := 0
	for {
		sum2 += j
		j++
		if j > 100 {
			break
		}
	}
	fmt.Println(sum2)

	// 遍历切片
	s := []string{"A", "B", "C", "D", "E"}
	for index := range s {
		fmt.Println(index, s[index])
	}

	// 遍历map
	map01 := map[string]string{
		"A": "A",
		"B": "B",
		"C": "C",
		"D": "D",
		"E": "E",
	}
	for key, value := range map01 {
		fmt.Println(key, value)
	}

	// break:跳过循环
	// continue:跳过本轮循环

}
