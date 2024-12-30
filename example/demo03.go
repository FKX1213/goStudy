package example

import "fmt"

// 模拟枚举类型
type Season uint8

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func Demo03() {
	fmt.Println(Spring) // 输出 0
	fmt.Println(Summer) // 输出 1
	fmt.Println(Autumn) // 输出 2
}

func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	}
	return ""
}
