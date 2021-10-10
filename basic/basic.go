package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3

// bb := true  在函数外不能这样写

// 批量声明
var (
	cc = 4
	dd = "ddd"
)

func variableZeroValue() {
	// 不赋值时会有默认值
	var a int
	var s string

	fmt.Printf("%d  %q\n", a, s)
}

func variable() {
	// 如果有多个同类型的值可以直接这样赋值
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	// 如果赋值可以明确类型，可以省略掉类型
	var a, b, c, d = 3, 4, "abc", true

	fmt.Println(a, b, c, d)
}

func variableShorter() {
	// := 可以表示声明+赋值
	a, b, c, d := 3, 4, "abc", true

	fmt.Println(a, b, c, d)
}

// 验证欧拉公式
func euler() {
	c := 3 + 4i

	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int

	// 强制类型转化
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4

	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
		_
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang, javascript)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")

	variableZeroValue()

	variable()

	variableTypeDeduction()

	variableShorter()

	fmt.Println(aa, cc, dd)

	euler()

	triangle()

	consts()

	enums()
}
