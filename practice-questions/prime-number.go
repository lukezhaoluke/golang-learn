//编写一个函数，输出100以内的所有素数，每行显示5个，并求和
//素数：
package main

import "fmt"
func main() {
	/*
	   素数：打印2-100内的素数。只能被2和本身整除。
	*/
	for i := 2; i <= 100; i++ {
		count := 0 //用于统计i被整除的次数
		for j := 2; j < i; j++ {
			if i%j == 0 {
				count++
				break
			}
		}
		// fmt.Println(count)
		if count == 0 {
			fmt.Println(i, "是素数")
		}
	}
}
