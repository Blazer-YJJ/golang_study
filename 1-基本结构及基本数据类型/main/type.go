/*
 * 类型可以是基本数据类型：
 * int float bool string
 * 结构化的类型：
 * struct array slice map channel
 * 描述类型行为：
 * interface
 */
package main

import "fmt"

func main() {
	type (
		IZ  int
		FZ  float64
		STR string
	) // 自定义类型
	var num IZ = 5
	var sat STR = "HELLO"
	var foo FZ = 3.1415926

	fmt.Println(num)
	fmt.Println(sat)
	fmt.Println(foo)
}
