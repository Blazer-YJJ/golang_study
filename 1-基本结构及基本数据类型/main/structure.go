/*
 * @Author: YangJijun yangjijuna@gmail.com
 * @Date: 2023-03-15 05:43:00
 * @LastEditors: YangJijun yangjijuna@gmail.com
 * @LastEditTime: 2023-03-15 06:59:48
 * @FilePath: \golang_study\1-基本结构及基本数据类型\structure.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
//  例子结构总体思路：
// 	1.在完成包的import之后，开始对常量变量及类型的定义或声明；
// 	2.若存在init函数，则对该函数进行定义
// 	3.若当前包是main包，则定义main函数
// 	4.然后定义其余函数，先是类型的方法，按照main函数中先后调用的顺序来定义相关函数
//  程序执行顺序：
// 	1.按顺序导入所有呗main引用的其他包，然后每个包执行流程如下：
// 	2.若改包又导入有其他包，则从第一步开始递归，但每个包只会被导入一次
// 	3.后以相反的顺序在每个包中初始常量与变量，若含有init函数则调用该函数
// 	4.完成以上步骤后，main也执行同样的过程，最后调用main函数执行程序。
package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // 初始化包
	// ...
}

func main() {
	var a int
	Func1()
	fmt.Println(a)
}

func (t T) Method1() {
	// ...
}

func Func1() { // 导出函数Func1
	// ....
}
