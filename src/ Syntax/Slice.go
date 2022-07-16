package main

import "fmt"

func createSlice() {
	//声明方法
	var list1 = []int{}
	var list2 []int
	// make([]Type, size ,cap)
	list3 := make([]int, 3, 7)

	sanguo := [...]string{"吕布", "赵云", 5: "张飞"}
	//数组变量【起始位置:结束位置】不包括结束位置元素
	var list4 = sanguo[1:4]
	list5 := sanguo[:]

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(list4)
	fmt.Println(list1 == nil)    //true
	fmt.Println(len(list1) == 0) //true
	fmt.Println(list5)
}

func changeSlice() {
	var array = [10]string{"诸葛亮", "司马懿", "庞统"}
	s := array[0:2]
	c := s

	fmt.Println(array)
	fmt.Println(s)
	fmt.Println(c)

	c[0] = "周瑜"

	fmt.Println(array)
	fmt.Println(s)
	fmt.Println(c)

}

//追加
func appendSlice() {
	s := []string{"诸葛亮", "司马懿", "庞统"}
	fmt.Printf("len: %d; cap: %d\n", len(s), cap(s))
	s = append(s, "周瑜")
	fmt.Printf("len: %d cap:%d\n", len(s), cap(s))
	s = append(s, "法正", "张昭")
	fmt.Printf("len: %d cap:%d\n", len(s), cap(s))
	// "..."拆箱
	s = append(s, []string{"荀彧"}...)
	fmt.Printf("len: %d cap:%d\n", len(s), cap(s))
}

func mSlice() {
	slice := [][]string{
		{"诸葛亮", "司马懿"},
		{"曹操", "刘备"},
		{"吕布", "赵云", "关羽"},
	}
	fmt.Println(slice)
}

func main() {
	appendSlice()
	changeSlice()
	mSlice()
}
