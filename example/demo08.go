package example

import "fmt"

func Demo0801(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 100
}

func Demo0802(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 111, 222
}

func Demo0803(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 100
	r2 = 100
	return
}

func Demo0804(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 101
	r2 = 202
	return
}

func Demo0805(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	r1 = 104
	r2 = 209
	return
}
