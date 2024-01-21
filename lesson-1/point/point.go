package main

import "fmt"

func incrementValue(val int) {
	val++
}

func incrementPoint(val *int) {
	*val++
}

func appendSlice(slice []int, value int) {
	slice = append(slice, value)
}

func appendSlicePoint(slice *[]int, value int) {
	*slice = append(*slice, value)
}

func addEmpty(m map[string]int, key string, value int) {
	m[key] = value
}

type Person struct {
	Name string
	Age  int
}

func birthday(p *Person) {
	p.Age++
}

func main() {
	a := 10
	fmt.Println("original 'a' value: ", a)
	// 传值
	incrementValue(a)
	fmt.Println("'a' after incrementValue: ", a)

	// 传引用
	incrementPoint(&a)
	fmt.Println("'a' after incrementPoint: ", a)

	// 引用传递 slice
	s := []int{1, 2, 3}
	appendSlice(s, 4)
	fmt.Println("s: ", s)
	appendSlicePoint(&s, 4)
	fmt.Println("s: ", s)

	// 引用传递Map
	m := make(map[string]int)
	addEmpty(m, "k1", 1)
	fmt.Println("m: ", m)

	// 引用传递结构体
	jiamu := Person{Name: "jiamu", Age: 18}
	birthday(&jiamu)
	fmt.Println(jiamu)

}
