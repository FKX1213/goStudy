package example

// 常量的声明需要用到const关键字，常量在声明时就必须初始化一个值，并且常量的类型可以省略
const name string = "Jack" // 字面量

const msg = "hello world" // 字面量

const num = 1 // 字面量

const numExpression = (1+2+3)/2%100 + num // 常量表达式

// 批量声明常量可以用()括起来以提升可读性，可以存在多个()达到分组的效果
const (
	Count = 1
	Name  = "Jack"
)

const (
	Size = 16
	Len  = 25
)

// 在同一个常量分组中，在已经赋值的常量后面的常量可以不用赋值，其值默认就是前一个的值
const (
	A = 1
	B // 1
	C // 1
	D // 1
	E // 1
)

// 内置的常量标识符

const (
	Num  = iota * 2 // 0
	Num1            // 2
	Num2            // 4
	Num3            // 6
	Num4            // 8
)

const (
	Num0  = iota<<2*3 + 1 // 1 第一行
	Num01 = iota<<2*3 + 1 // 13 第二行
	_                     // 25 第三行
	Num03                 //37 第四行
	Num04 = iota          // 4 第五行
	_                     // 5 第六行
	Num05                 // 6 第七行
)
