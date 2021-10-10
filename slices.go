package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	// 切片是数组的view，当数组发生改变时，切片也会一起改变
	s := arr[2:]
	fmt.Println("s = ", s)
	arr[4] = 100
	fmt.Println(arr)
	fmt.Println(s)
}
