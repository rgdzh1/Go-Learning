package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串是一个字节切片
	// 1. 字符串是不允许修改的
	// str1 := "hello"
	// str1[0]='H' // cannot assign to str1[0] 字符串是不允许修改的
	// 2.字符串长度,返回的是字节个数,一个英文字母代表一个字节,一个中文长度为三个字节
	str2 := "hello"
	str3 := "hello world"
	str4 := "hello world你好"
	fmt.Println(len(str2))
	fmt.Println(len(str3))
	fmt.Println(len(str4))
	// 3.获取字符串中第一个字节
	fmt.Println(str4[0])        // 打印的是字节
	fmt.Printf("%c\n", str4[0]) // 这样打印的是字符
	a := 'h'
	b := 104 // 104 ASCII码对应的就是c
	fmt.Printf("%c,%c,%c\n", str4[0], a, b)
	// 5.字符串与字节转化
	slice1 := []byte{65, 66, 67, 68} // 字节数组转为字符串
	str5 := string(slice1)           // 将字节数组转为字符串
	fmt.Println(str5)

	str6 := "abcde"
	slice2 := []byte(str6) // 将字符串转为字节数组
	fmt.Println(slice2)
	// string API
	str7 := "hello world"
	fmt.Println(strings.Contains(str7, "hell"))    // 是否包含
	fmt.Println(strings.ContainsAny(str7, "abdh")) // 表示是否包含任意一个字符字符
	strings.Count(str7, "llo")                     // 统计substr在str7中出现的次数
	strings.HasPrefix(str7, "h")                   // 判断是否以指定内容开始
	strings.HasSuffix(str7, "d")                   // 判断是否以指定内容结尾
	strings.Index(str7, "l")                       // 查找substr在str7中的位置,如果不存在就返回-1
	strings.IndexAny(str7, "abcdh")                // 查找chars中任意的字符,第一个出现在str7中的位置
	strings.LastIndex(str7, "l")                   //查找substr在str7中最后一次出现的位置

	// 字符串拼接
	ss1 := []string{"aaa", "bbb", "ccc"}
	str8 := strings.Join(ss1, "-") //字符串拼接,以"-"
	fmt.Println(str8)
	// 切割
	ss2 := "aaa,bbb,ccc"
	slice3 := strings.Split(ss2, ",")
	fmt.Println(slice3)

	ss3 := strings.Repeat("hello", 5)
	fmt.Println(ss3)
	ss4 := "lvlvlvlvlvlvl"
	ss5 := strings.Replace(ss4, "l", "*", 2) // n:表示2次,如果全部替换就使用-1表示.
	fmt.Println(ss5)

	strings.ToLower(ss4) //转小写
	strings.ToUpper(ss4) //转大写

	// 字符串截取
	ss6 := "123QWEASDZXCV"
	ss7 := ss6[:5]
	fmt.Println(ss7)
	ss8 := ss6[5:]
	fmt.Println(ss8)
}
