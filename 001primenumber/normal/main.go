package main

//标准计算素数的方法，1不是素数，2是最小的素数
import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for j := 1; j <= 80000; j++ {
		for i := 2; i <= j; i++ {
			if j%i == 0 && i < j {
				break
			}
			if i == j {
				//fmt.Printf("素数=%d\n", j)
			}
		}

	}
	end := time.Now().Unix()
	fmt.Printf("运算结束,整体耗时=%d", end-start)
}
