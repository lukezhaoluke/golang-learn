//循环打印输入的月份的天数【使用continue实现】，要有判断输入的月份是否错误的语句
//分析:首先判断输入的年份是否为闰年，
//判断输入月份的数值是否为1-12的数字
//闰年2月份为29天，否则输出28天

package main

import (
	"fmt"
)

//判断是否为闰年函数
func leap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	var year int
	var month int
	for {
		fmt.Println("请输入年份")
		_, _ = fmt.Scanln(&year)
		fmt.Println("请输入月份")
		_, _ = fmt.Scanln(&month)
		if month < 1 || month > 12 {
			fmt.Println("输入月份错误，请重新输入")
			continue
		}
		if month == 2 {
			if leap(year) == true {
				fmt.Printf("%v年%v月为29天\n", year, month)
			} else {
				fmt.Printf("%v年%v月为28天\n", year, month)
			}
		} else if month == 4 || month == 6 || month == 9 || month == 11 {
			fmt.Printf("%v年%v月为30天\n", year, month)
		} else {
			fmt.Printf("%v年%v月为31天\n", year, month)
		}

	}
}
