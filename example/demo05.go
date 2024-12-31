package example

import "fmt"

/**
条件判断
*/

func Demo0501() {
	a, b := 1, 2
	if a > b {
		b++
	} else {
		a++
	}
	fmt.Println("a = ", a, " , b = ", b)
}

func Demo0502() {
	if x := 1 + 1; x >= 2 {
		fmt.Println(x)
	}
}

func Demo0503() {
	a, b := 1, 2
	if a<<1%100+3 > b*100/20+6 { // (a<<1%100)+3 > (b*100/20)+6
		b++
	} else {
		a++
	}
	fmt.Println("a = ", a, " , b = ", b)
}

func Demo0504() {
	score := 90
	var ans string
	if score == 100 {
		ans = "S"
	} else if score >= 90 && score < 100 {
		ans = "A"
	} else if score >= 80 && score < 90 {
		ans = "B"
	} else if score >= 70 && score < 80 {
		ans = "C"
	} else if score >= 60 && score < 70 {
		ans = "E"
	} else if score >= 0 && score < 60 {
		ans = "F"
	} else {
		ans = "nil"
	}
	fmt.Println(ans)
}

func Demo0505() {
	score := 90
	var ans string
	if score >= 0 && score < 60 {
		ans = "F"
	} else if score < 70 {
		ans = "D"
	} else if score < 80 {
		ans = "C"
	} else if score < 90 {
		ans = "B"
	} else if score < 100 {
		ans = "A"
	} else if score == 100 {
		ans = "S"
	} else {
		ans = "nil"
	}
	fmt.Println(ans)
}

func Demo0506() {
	str := "a"
	switch str {
	case "a":
		str += "a"
		str += "c"
	case "b":
		str += "bb"
		str += "aaaa"
	default: // 当所有case都不匹配后，就会执行default分支
		str += "CCCC"
	}
	fmt.Println(str)
}

func Demo0507() {
	switch num := f(); { // 等价于 switch num := f(); true {
	case num >= 0 && num <= 1:
		num++
		fmt.Println("case 1 ，num = ", num)
	case num > 1:
		num--
		fmt.Println("case 2， num = ", num)
		fallthrough
	case num < 0:
		num += num
		fmt.Println("case 3， num = ", num)
	}
	fmt.Println("case closed， num = ", num)
}

func f() int {
	return 1
}

func Demo0508() {
	num := 2
	switch { // 等价于 switch true {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
	case num < 0:
		num *= num
	}
	fmt.Println(num)
}

func Demo0509() {
	num := 2
	switch {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
		fallthrough // 执行完该分支后，会继续执行下一个分支
	case num < 0:
		num += num
	}
	fmt.Println(num)
}
