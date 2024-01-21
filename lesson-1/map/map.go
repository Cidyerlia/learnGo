package main

import "fmt"

func main() {
	// 创建map
	m := make(map[string]string)
	if m == nil {
		fmt.Println("emptyMap is nil")
	}
	fmt.Println("m: ", m)

	// 创建一个空map
	var emptyMap map[string]int
	//emptyMap["k1"] = 2  // 不能向空Map里添加键值对，因为系统没有给空map分配空间
	//emptyMap = make(map[string]int) // 分配空间后，才可以使用
	//emptyMap["k1"] = 2
	if emptyMap == nil {
		fmt.Println("emptyMap is nil")
	}

	// 设置键值对
	m["k1"] = "first"
	m["k2"] = "second"
	fmt.Println("m: ", m)

	// 获取键k1值
	v := m["k1"]
	fmt.Println("value: ", v)

	// 使用len() 返回 map中的键值对数量
	fmt.Println("len： ", len(m))

	// 内置的delete移除键值对
	delete(m, "k1")

	// 当获取的map中不存在的键值对，有默认的返回值
	v = m["k3"]
	fmt.Println("v: ", v)

	// 可以通过两个变量接受的方式，判断键是否存在
	_, ok := m["k2"]
	if ok {
		fmt.Println("key k2 exist? ", ok)
	}

	// 遍历是for range
	for key, value := range m {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

}
