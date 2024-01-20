package main

import "fmt"

func main() {

	// 基本的if语句
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	}

	// if-else
	if num > 0 {
		fmt.Println(num, "是正数")
	} else {
		fmt.Println(num, "是负数")
	}

	// if-else if
	if num > 10 {
		fmt.Println(num, "大于10")
	} else if num == 10 {
		fmt.Println(num, "等于10")
	} else {
		fmt.Println(num, "小于10")
	}

	// 用if语句定义局部变量
	if scope := 24; scope > num {
		fmt.Println("scope > num")
	}

	// switch case 是个值
	n := 3
	switch n {
	case 1:
		fmt.Println("n == 1")
	case 3:
		fmt.Println("n == 3")
	default:
		fmt.Println("n == xx")
	}

	// switch case 是个表达式
	switch {
	case n%2 == 0:
		fmt.Println("n是奇数")
	default:
		fmt.Println("n是偶数")
	}
}
