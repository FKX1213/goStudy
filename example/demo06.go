package example

import "fmt"

func Demo601() {
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}
}

func Demo602() {
	for i, j := 1, 2; i < 100 && j < 1000; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}

func Demo603() {
	num := 1
	for num < 100 {
		num *= 2
	}
	fmt.Println(num)
}

func Demo604() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i <= j {
				fmt.Printf("%d*%d = %2d  ", i, j, i*j)
			}
		}
		fmt.Println()
	}
}

func Demo605() {
	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, value)
	}
}
