package main

import "fmt"

func createMap() {
	// make(map[keyType]valueType)
	map1 := make(map[string]int)
	fmt.Println(map1)

	map2 := map[string]string{
		"1": "2",
		"4": "8",
	}
	delete(map2, "1")
	map2["2"] = "4"
	fmt.Println(map2)
	fmt.Println(map2["2"])

	v1, exist := map2["1"]
	fmt.Println(v1)
	fmt.Println(exist)

	v2, exist := map2["2"]
	fmt.Println(v2)
	fmt.Println(exist)

	for key, value := range map2 {
		fmt.Printf("key: %s , value: %s", key, value)
	}
}

func main() {
	createMap()
}
