package goType

// 每个类型均有一个基础的类型T
// 如果T 声明了boolean / string / int / goType literal ，那么其基础类型就是T 本身

type UnderlyingInt NamedInt //基础类型为 int

type UnderlyingSlice UnnamedSlice

type UnderlyingMap map[NamedString]int // 基础类型为 map[NamedString]int ，而非map[string]int
