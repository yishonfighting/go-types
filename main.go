package main

import (
	"fmt"
	"go-test/goTypes/goType"
)

// golang 类型定义可以分为三类
// named types: 内置类型,包括 int string int64等
// unnamed types:复合类型，包括数组，结构体，指针，函数，接口等
// underlying types: 每个类型均有一个基础的类型T ,如果T 声明了boolean / string / int / goType literal ，那么其基础类型就是T 本身

func main() {

	//类型一致性
	//两种类型要么相同，要么不相同。
	//一个命名类型一定和其它类型都不同。
	//如果他们基础类型的字面量在结构上是等价的，他们就是相同的类型。
	//因此，预先声明的命名类型 int, int64, 等都是不一致的。

	//demo 只列举一种编译通过的情况
	unnamedTypeExchange()
}

//unnamedType ：基础类型一致即可进行类型转换
func unnamedTypeExchange() {
	var testOne goType.UnnamedStructOne

	var testTwo goType.UnnamedStructTwo

	testOne = goType.UnnamedStructOne{
		Name: "here is one",
	}

	testTwo = goType.UnnamedStructTwo(testOne)

	fmt.Println("unnamedType To unnamedType:")
	fmt.Println(testOne, testTwo)
}
