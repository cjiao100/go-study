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

	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	// 一个切片中会包含切片开始位置、切片长度、cap创图
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	// s4 和 s5不再是arr的view
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
}
