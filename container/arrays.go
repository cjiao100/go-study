package main

import "fmt"

// 值传递
//func printArray(arr [5]int) {
//	for i, v := range arr {
//		fmt.Println(i, v)
//	}
//}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	// 二维函数
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)

	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(i, arr3[i])
	}

	// 等价上面写法
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(&arr1)
	// [3]int 和 [5]int 在go中会被认为是不同的数据类型
	//printArray(arr2)
	printArray(&arr3)

	fmt.Println(arr1, arr3)
}
