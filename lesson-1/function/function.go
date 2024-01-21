package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func closure() func() int {
	sum := 0
	return func() int {
		sum++
		return sum
	}
}

func applyFunc(a int, b int, f func(a int, b int) int) int {
	return f(a, b)
}

func createAdder() func(int) int {
	return func(addend int) int {
		return addend + 5
	}
}

func main() {
	// 调用add函数
	result := add(2, 3)
	fmt.Println("result: ", result)

	// 匿名函数直接执行
	func(message string) {
		fmt.Println(message)
	}("hello world")

	// 匿名函数赋值给变量
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println("sum : ", sum(3, 5))

	increment := closure()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	// 函数作为参数
	sum1 := applyFunc(10, 20, add)
	fmt.Println("10 + 20 = ", sum1)

	// 函数作为返回值
	adder := createAdder()
	fmt.Println("40 + 5 = ", adder(40))
}
