package main

import "fmt"

func main() {

	// 基本的for
	for i := 0; i < 5; i++ {
		fmt.Printf("第%d次循环\n", i)
	}

	//  for省略前置语句
	j := 5
	for ; j > 0; j-- {
		fmt.Printf("第%d次循环\n", j)
	}

	// for循环只包含条件表达式，类似while语句
	k := 0
	for k < 5 {
		fmt.Printf("第%d次循环\n", k)
		k++
	}

	//无限循环 相当于while true
	//for {
	//	fmt.Println("进入死循环")
	//}

	// for range 遍历数组
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	// for range 遍历字符串
	for index, char := range "hello" {
		fmt.Printf("index： %d , charactor is %c\n", index, char)
	}

	// 遍历map
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("key is %s and value is %d\n", key, value)
	}

	// 遍历通道
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for n := range ch {
		fmt.Println("从通道取值：", n)
	}

}
