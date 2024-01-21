package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = "hello world"
	s2 := "你好，世界"

	// 查看长度
	fmt.Println("The length of s1: ", len(s1))
	fmt.Println("The length of s2: ", len(s2))

	// 单字符打印
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%x ", s2[i])
	}
	fmt.Println()
	for _, r := range s2 {
		fmt.Printf("%x ", r)
	}
	fmt.Println()

	// 字符串拼接
	s3 := s1 + s2
	fmt.Println(s3)

	// 字符串比较
	if s1 == "hello world" {
		fmt.Println("s1 equals to 'hello world'")
	}

	// 字符串分割
	parts := strings.Split(s3, " ")
	for _, part := range parts {
		fmt.Println(part)
	}

	// 字符串包含
	if strings.Contains(s3, "世界") {
		fmt.Println("s3 contains '世界'")
	}

	// 字符串替换，替换所有匹配项(n = -1)
	news3 := strings.Replace(s3, "hello", "hi", -1)
	fmt.Println(news3)

	// 子字符串判断
	isPrefix := strings.HasPrefix(s3, "hello")
	isSuffix := strings.HasSuffix(s3, "世界")
	fmt.Println("Prefix: ", isPrefix, "Suffix: ", isSuffix)

	// 中文字符串截取
	r := []rune(s2)

	// 单字符打印
	for i := 0; i < len(r); i++ {
		fmt.Printf("char %d: %c\n", i, r[i])
	}

	// 截取中文字符串的前两个文字
	substr := string(r[:2])
	fmt.Println("substr: ", substr)

}
