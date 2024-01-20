package main

import "fmt"

func modifyArray(arr [5]int) {
	arr[0] = 100
	fmt.Println("modify arr: ", arr)
}

func main() {

	// 定义长度为5的数组，默认值为0
	var a [5]int
	fmt.Println("array-a: ", a)

	// 设置数组特定位置的值
	a[4] = 100
	fmt.Println("array-a: ", a)

	// 获取数组的长度
	fmt.Println("array-a的长度是：", len(a))

	// 使用数组字面量初始化并赋值
	b := [5]int{12, 3, 4, 5, 6}
	fmt.Println("array-b: ", b)

	// 二维数组声明和初始化
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("array-twoD: ", twoD)

	// 声明数组省略长度，编译器自动计算
	c := [...]int{6, 7, 8}
	fmt.Println("array-c: ", c)

	// 遍历数组 -- 不介绍了，一般都是for range遍历

	// 重点：数组传参是传值！！
	d := [5]int{1, 1, 1, 1, 1}
	fmt.Println("before modify: ", d)
	modifyArray(b)
	fmt.Println("after modify: ", d)
}
