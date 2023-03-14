/*
 * @Author: YangJijun yangjijuna@gmail.com
 * @Date: 2023-03-14 23:51:26
 * @LastEditors: YangJijun yangjijuna@gmail.com
 * @LastEditTime: 2023-03-15 06:58:37
 * @FilePath: \golang_study\1-基本结构及基本数据类型\main\type.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
