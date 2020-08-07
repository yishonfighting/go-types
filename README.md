# goTypes
decription types system in go 

```
A type system consists of :
Basic types — Included in the programming language and available to any program written in that language. Go has various basic types(int8 , uint8 ( byte ), int16 , uint16 , int32 ( rune ), uint32) etc.
Type constructors — Way for a programmer to define new types.
Eg. Pointer to T, where T is a Type or Struct {a: T}
Type Inference — The compiler can infer the type of a variable or a function without us having to explicitly specify it. Go has Uni-directional type inference.
Type Compatibility — Which assignments are allowed by the type system? a int; b int8; a = b; ?
How to determine if two types are equal? In Go assignability is what mostly determines whether types can be used interchangeably. We will look at this later in details.
```
