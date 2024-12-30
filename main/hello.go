package main

import "fmt"

func main() {
	// 声明多个变量
	var xx, yy int = 100, 200
	var s1, s2 string = "hello", "world"
	fmt.Println("xx= ", xx, ", yy= ", yy, ", s1= ", s1, ", s2= ", s2)

	// 多变量声明，多行
	var (
		cc int     = 100
		dd float32 = 3.14
		ee bool    = true
	)
	fmt.Println("cc= ", cc, ", dd= ", dd, ", ee= ", ee)
}
