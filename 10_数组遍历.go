package main

import "fmt"

func main() {
	var arr1 = [...]int{0, 1, 2, 3, 4, 5}
	// 第一种for循环遍历
	for n := 0; n < len(arr1); n++ {
		fmt.Println(arr1[n])
	}
	// 第二种遍历
	for index, value := range arr1 {
		fmt.Printf("下标是: %d,数值是: %d\n", index, value)
	}
	// 数组类型
	arr2 := [2]int{0, 5}
	arr3 := [2]int{0, 1}
	arr2 = arr3
	fmt.Println(arr2)
	fmt.Println(arr3)
	arr3[1] = 15555
	fmt.Println(arr2) // [0 1] arr3中元素改变不影响arr2中元素值, arr2=arr3是值传递,就是copy了一份值, 和指针无关.
	fmt.Println(arr3) // [0 15555]

	arr4 := [2]int{0, 5}
	arr5 := [2]int{0, 5}
	fmt.Println(arr4 == arr5)
}
