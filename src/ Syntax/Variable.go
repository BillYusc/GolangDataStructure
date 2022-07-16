package main

import (
	"fmt"
	"math"
	"unsafe"
)

func Integer() {
	var num8 int8 = math.MaxInt8
	var num16 int16 = math.MinInt16
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MinInt64

	fmt.Printf("num8 类型：%T ; 大小：%d ; 值： %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16 类型：%T ; 大小：%d ; 值： %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32 类型：%T ; 大小：%d ; 值： %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64 类型：%T ; 大小：%d ; 值： %d\n", num64, unsafe.Sizeof(num64), num64)
}

func UInteger() {
	var num8 uint8 = math.MaxInt8
	var num16 uint16 = math.MaxInt16
	var num32 uint32 = math.MaxInt32
	var num64 uint64 = math.MaxInt64

	fmt.Printf("num8 类型：%T ; 大小：%d ; 值： %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16 类型：%T ; 大小：%d ; 值： %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32 类型：%T ; 大小：%d ; 值： %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64 类型：%T ; 大小：%d ; 值： %d\n", num64, unsafe.Sizeof(num64), num64)
}

func Float() {
	var num1 float32 = math.MaxFloat32
	num2 := math.MaxFloat64
	fmt.Printf("num1 类型 : %T ; 大小：%d ; 值： %f\n", num1, unsafe.Sizeof(num1), num1)
	//%g 科学计数法
	fmt.Printf("num2 类型 : %T ; 大小：%d ; 值： %g\n", num2, unsafe.Sizeof(num2), num2)
}

func main() {
	Integer()
	fmt.Println("--------------")
	UInteger()
	fmt.Println("--------------")
	Float()
}
