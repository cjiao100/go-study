package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // 这样声明出来是空map

	var m3 map[string]int // 这样声明出来是nil, go的nil可以参加到运算中

	fmt.Println(m, m2, m3)
	fmt.Println(m2 == nil)
	fmt.Println(m3 == nil)

	fmt.Println("Traversing map")

	// map的顺序是不定的，无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	// m["xxx"] 会有两个返回值，第一个是取到的value，第二个是是否获取到
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	// 如果获取的key不存在，也不会报错，直接返回一个空
	couseName, ok := m["couse"]
	fmt.Println(couseName, ok)

	fmt.Println("Deleting Values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
