package main

import "fmt"

// 声明单个常量

const PI = 3.14

// 声明多个常量
const (
	StatusOk  = 200
	NotFound  = 404
	Forbidden = 403
)

func main() {
	fmt.Println("圆周率：", PI)
	fmt.Println("状态码OK: ", StatusOk)
	fmt.Println("状态码NotFound: ", NotFound)
	fmt.Println("Forbidden: ", Forbidden)

	// 常量用来枚举
	const (
		Red = iota // 默认是0开始
		Green
		Blue
	)

	fmt.Println("红色的枚举值：", Red)
	fmt.Println("绿色的枚举值：", Green)
	fmt.Println("蓝色的枚举值：", Blue)

	// 常量表达式可以执行任意精度的计算
	const b = 3e20 / PI
	fmt.Println("b: ", b)

}
