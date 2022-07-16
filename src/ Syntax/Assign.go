package main

import (
	"fmt"
)

//常量定义
const (
	LIUBEI = iota * 10
	GUANYU
	ZHANGFEI
)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出
)

func main() {
	var a int
	var b = 100.0
	var c, d = "你好", 99.9
	e := "iota"
	var f string = "111"

	fmt.Printf("a = %d , a 的类型%T\n", a, a)
	fmt.Printf("b = %f , b 的类型%T\n", b, b)
	fmt.Printf("c = %s , c 的类型%T\n", c, c)
	fmt.Printf("d = %f , d 的类型%T\n", d, d)
	fmt.Printf("e = %s , e 的类型%T\n", e, e)
	fmt.Printf("f = %s , f 的类型%T\n", f, f)

	fmt.Printf("LIUBEI = %d a的类型%T\n", LIUBEI, LIUBEI)
	fmt.Printf("GUANYU = %d a的类型%T\n", GUANYU, GUANYU)
	fmt.Printf("ZHANGFEI = %d a的类型%T\n", ZHANGFEI, ZHANGFEI)

}
