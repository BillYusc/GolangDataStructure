package main

import "fmt"

/*
	& 放在变量前，取变量的指针，返回的是指针
	* 放在指针前，取指针的变量，返回的是变量
*/
func Pointer1() {
	var x = "111"
	ptr1 := &x

	fmt.Println("x=", x)
	fmt.Println("&x=", &x)
	fmt.Println("ptr1=", ptr1)
	fmt.Println("*ptr1=", *ptr1)
}

func Pointer2() {
	// new 创建一个指针分配好内存，再给指针写入值
	ptr1 := new(string)
	*ptr1 = "紫薯布丁"
	fmt.Println(ptr1)
	fmt.Println(*ptr1)

	//通过定义普通变量，获取指针
	s1 := "兰州烧饼"
	ptr2 := &s1
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	//声明一个指针变量，再从其他变量获取内存地址指针变量
	s2 := "大胆想法"
	var ptr3 *string
	ptr3 = &s2
	fmt.Println(ptr3)
	fmt.Println(*ptr3)

	var ptrptr **string
	s3 := "索然无味"
	ptr4 := &s3
	ptrptr = &(ptr4)
	fmt.Println(ptrptr)
	fmt.Println(ptr4)
	fmt.Println(s3)
}

func Pointer3() {
	var x = "111"
	var ptr *string
	fmt.Println("ptr: ", ptr) // <nil>
	ptr = &x
	fmt.Println("prt:", ptr)
}

func main() {
	Pointer3()
}
