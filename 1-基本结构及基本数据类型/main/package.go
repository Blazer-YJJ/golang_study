/*
 * @Author: YangJijun yangjijuna@gmail.com
 * @Date: 2023-03-14 23:39:44
 * @LastEditors: YangJijun yangjijuna@gmail.com
 * @LastEditTime: 2023-03-14 23:48:29
 * @FilePath: \golang_study\1-基本结构及基本数据类型\main\page.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import fm "fmt" // 使用包别名解决包冲突
// 当有多个包导入时可用以下优雅的方式处理（因式分解关键字）
// import (
// 	"fmt"
// 	"os" // 注意事项：如导入包后不使用会报错
// )

func main() {
	fm.Println("hello package")
}
