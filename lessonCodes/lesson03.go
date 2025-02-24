package main

import "fmt"

/**
 * @date:       2025/2/24 10:25
 * @author:     cola1213
 * @description:数据类型
 */
func main() {
	var n1 uint8 = 2
	var n2 uint16 = 2
	var n3 uint32 = 2
	var n4 uint64 = 2
	var n5 uint = 2
	var n6 int8 = 2
	var n7 int16 = 2
	var n8 int32 = 2
	var n9 int64 = 2
	var n10 int = 2

	fmt.Println(n1, n2, n3, n4, n5, n6, n7, n8, n9, n10)

	var c1 = 'a'
	var c2 = 97
	fmt.Println(c1, c2)           // 直接打印都是数字
	fmt.Printf("%c %c\n", c1, c2) // 以字符的格式打印

	var r1 rune = '中'
	fmt.Printf("%c\n", r1)

	var s string = "枫枫知道"
	fmt.Println(s)

	fmt.Println("枫枫\t知道")              // 制表符
	fmt.Println("枫枫\n知道")              // 回车
	fmt.Println("\"枫枫\"知道")            // 双引号
	fmt.Println("枫枫\r知道")              // 回到行首
	fmt.Println("C:\\pprof\\main.exe") // 反斜杠

	var s3 = `今天
天气
真好
`
	fmt.Println(s3)

	var a1 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a1) // 0
	fmt.Printf("%#v\n", a2) // 0
	fmt.Printf("%#v\n", a3) // ""
	fmt.Printf("%#v\n", a4) // false
}
