package main

import "fmt"

const ()

func Array() {

	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{5, 3}
	fmt.Println(arr3)

	arr4 := [5]int{1: 99, 4: 100}
	fmt.Println(arr4)

	arr5 := [...]int{15, 7: 20, 25, 100, 75}
	fmt.Println(arr5)

	fmt.Printf("类型\n arr1: %T;\n arr2: %T;\n arr3: %T;\n arr4: %T;\n arr5: %T;\n ", arr1, arr2, arr3, arr4, arr5)

	//二维数组
	barr := [3][2]string{
		{"刘备", "玄德"},
		{"关羽", "云长"},
		{"张飞", "翼德"},
	}
	fmt.Print(barr)
	//len() 计算数组长度
	fmt.Println("数组长度: ", len(barr))
}

func showArr() {
	arr := [...]int{15, 7: 20, 25, 100, 75}

	for index, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
	}
	//下划线
	for _, value := range arr {
		fmt.Printf("value = %d\n", value)
	}
}

func PrintSanguo() {
	//数组是值类型，不是引用类型
	sanguo := [...]string{"吕布", "赵云", 5: "张飞"}
	copy := sanguo
	copy[2] = "典韦"
	fmt.Println(sanguo)
	fmt.Println(copy)
}

func main() {
	// fmt.Println("----------")
	// Array()
	// showArr()
	PrintSanguo()
}
