package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//func div(a, b int) (q, r int) {
//	return a / b, a % b
//}

// 设定了返回的具体的值后，可以这样写
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d , %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 指针写法
//func swap(a, b *int) {
//	*a, *b = *b, *a
//}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	//fmt.Println(
	//	eval(0, 0, "+"),
	//	eval(3, 4, "*"),
	//)

	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	// 对于这种函数，编译器可以直接添加接受的变量
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
