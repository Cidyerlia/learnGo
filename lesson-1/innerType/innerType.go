package main

import "fmt"

func main() {
	// Go 语言内建类型分为5种，分别是整数类型、浮点数类型、字节类型、复数类型以及布尔类型

	// 整数类型

	// 定义一个Int类型的整数，int 默认是4字节
	var a int = 10
	fmt.Println(a)

	// 定义一个Int8类型的整数
	var b int8 = 10
	fmt.Println(b)

	// 定义一个uint类型的整数
	var c uint = 10
	fmt.Println(c)

	// 浮点数类型，不同于其它语言，浮点数类型名称都是float，没有double
	var d float32 = 10.5
	fmt.Println(d)

	// 字节类型
	var e byte = 'a'
	fmt.Println(e)

	var f rune = '你'
	fmt.Println(f)

	// 复数类型 3+3i
	var g complex64 = 3 + 3i
	fmt.Println(g)

	// 布尔类型
	var h bool = true
	fmt.Println(h)

}
