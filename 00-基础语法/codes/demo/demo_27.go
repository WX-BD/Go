package main

import (
	"fmt"
)

func main() {
	str := "this is a string"
	fmt.Println(&str)
	bytes := []byte(str)
	// 修改字节切片
	bytes = append(bytes, 96, 97, 98, 99)
	// 赋值给原字符串
	str = string(bytes)
	fmt.Println(str)
}
