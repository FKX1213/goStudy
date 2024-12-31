package main

import (
	"fmt"
	"goStudy/example"
)

func main() {
	fmt.Println("============Demo0801=========")
	demo0801 := example.Demo0801("hello", 123)
	fmt.Println(demo0801)
	fmt.Println("============Demo0802=========")
	demo0802, demo08021 := example.Demo0802("hello", 123)
	fmt.Println(demo0802)
	fmt.Println(demo08021)
	fmt.Println("============Demo0803=========")
	r1, r2 := example.Demo0803("adam", 32)
	fmt.Println("r1 = ", r1, " , r2 = ", r2)
	fmt.Println("============Demo0804=========")
	r3, r4 := example.Demo0804("bob", 345)
	fmt.Println("r3 = ", r3, " , r4 = ", r4)
	fmt.Println("============Demo0805=========")
	r5, r6 := example.Demo0805("david", 9832)
	fmt.Println("r5 = ", r5, " , r6 = ", r6)
}
