package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建map
	var map1 map[int]string         // map1属于nil map,没有初始化,不能向其中插入数据
	var map2 = make(map[int]string) // 创建一个初始化过的map
	var map3 = map[int]string{10: "小明", 20: "小红", 30: "小兰"}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	// 使用map时候的安全的做法
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}
	// 遍历map, map是无序的,这一点需要注意
	map4 := make(map[int]string)
	map4[0] = "小明"
	map4[1] = "小张"
	map4[2] = "小华"
	map4[3] = "小丽"
	map4[4] = "小王"
	map4[5] = "小天"
	map4[6] = "小静"
	for k, v := range map4 {
		fmt.Printf("键 %d 值 %s\n", k, v)
	}
	// 按照有序的方式打印map中的数据
	keys := make([]int, 0, len(map4))
	for k, _ := range map4 {
		// 切片append之后要将内存地址值重新赋值给切片, 因为切片底层是一个动态的数组
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, _ := range keys {
		fmt.Println(map4[i])
	}
	// map与切片结合使用
	map5 := map[string]string{"姓名": "张三", "年龄": "40"}
	map6 := map[string]string{"姓名": "李晓", "年龄": "20"}
	map7 := map[string]string{"姓名": "西控", "年龄": "80"}
	slice1 := make([]map[string]string, 0, 3) // 这个切片的数据类型是 map[string]string, 这种与以前的有点不大相同
	slice1 = append(slice1, map5)
	slice1 = append(slice1, map6)
	slice1 = append(slice1, map7)
	for _, v := range slice1 {
		for key, value := range v {
			fmt.Printf("%s %s \n", key, value)
		}
	}
}
