package main

import "fmt"

/**
 * @date:       2025/2/24 10:00
 * @author:     cola1213
 * @description:输入输出
 */
func main() {
	// 自动换行，中间有空格
	fmt.Println("cola", 12)

	// %s string类型的占位符 \n 是换行符
	fmt.Printf("%s is a boy\n", "bob")
	fmt.Printf("%s is a boy\n", "bob")
	// %d 整数
	fmt.Printf("%d\n", 3)
	// %.2f 保留2位小数（四舍五入）
	fmt.Printf("%.2f\n", 1.2562)
	// 打印空字符串
	fmt.Printf("%#v\n", "")

	// 输入
	var name2 string
	fmt.Scan(&name2) // 这里记住，要在变量的前面加个&, 后面讲指针会提到
	fmt.Println("你输入的名字是", name2)
}
