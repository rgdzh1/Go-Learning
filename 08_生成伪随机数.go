package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()      // 获取时间
	nano := now.UnixNano() // 以当前时间戳作为种子
	rand.Seed(nano)
	for m := 0; m <= 10; m++ {
		i := rand.Intn(10) // 生成[0~10) 0到9之间的伪随机数
		fmt.Println(i)
	}
}
