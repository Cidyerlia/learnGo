package main

import "fmt"

func main() {
	// 创建一个空切片，长度和容量都为0
	var s []int
	fmt.Println("Empty slice: ", s)

	// 用make创建一个指定类型、长度和容量的切片
	s = make([]int, 3) // 创建一个长度为3的切片, 默认element是0
	fmt.Println("make创建的切片： ", s)

	// 切片初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("init slice: ", s1)

	// 向切片中添加元素
	s = append(s, 3, 4, 5)
	fmt.Println("append element: ", s)

	// 切片中获取子切片, 获取index 1 - 3的切片
	s2 := s1[1:3] // 左闭右开
	fmt.Println("sub slice s2: ", s2)

	s2[1] = 100
	fmt.Println("sub slice s2: ", s2)
	fmt.Println("origin slice s1: ", s1)

	// 切片的零值是 nil
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("nilSlice is nil")
	}

	// 遍历切片和遍历数组的方式一致，也是建议使用 for range方式遍历

	// 切片长度和容量
	sl := []int{1, 1, 1, 1, 1, 1, 1, 1}
	// 使用 len() 函数获取长度
	fmt.Println("Length of s ", len(sl))

	// 使用cap() 函数获取容量
	fmt.Println("Capacity of s ", cap(sl))

	// 改变切片的元素会影响原始数组
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[:4]
	mySlice[0] = 100
	fmt.Println("myArray: ", myArray)
	fmt.Println("mySlice: ", mySlice)

	// 切片的复制操作
	copyS := make([]int, len(sl))
	copy(copyS, sl)
	fmt.Println("Copy sl to copyS ", copyS)

	// 创建一个动态数组，实现动态添加元素
	dynamicSlice := []int{}
	for i := 0; i < 10; i++ {
		dynamicSlice = append(dynamicSlice, i)
	}
	fmt.Println("dynamicSlice is ", dynamicSlice)

}
