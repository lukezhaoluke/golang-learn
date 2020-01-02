//编写一个函数，判断打鱼还是晒网
//中国有句俗话叫"三天打渔两天晒网"，如果从1990年1月1日开始执行三天打鱼两天晒网。如何判断在以后的某一天中是打渔还是晒网
//分析
//1)计算从1990年1月1日与指定日期相差多少天；需要注意的是年份有闰年，闰年2月为29天，平年为28天
//  因此通过时间差来计算天数，一天的unix时间戳为：86400
//2)由于“打鱼”和“晒网”的周期为5天，所以将计算出的天数用5去除；
//3)根据余数判断他是在“打鱼”还是在“晒网”；
//若 余数为0，1，2，则他是在“打鱼”
//否则 是在“晒网”
package main

import (
	"fmt"
	"time"
)

func GetTimeArr(start, end string) int64 {
	loc, _ := time.LoadLocation("Local")
	// 字符串转成时间戳
	startUnix, _ := time.ParseInLocation("2006-01-02", start, loc)
	endUnix, _ := time.ParseInLocation("2006-01-02", end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	//计算相差天数
	data := (endTime - startTime) / 86400
	if data % 5 == 0 || data % 5 == 1 || data % 5 == 2 {
		fmt.Printf("距离1990-01-01,已经过去%v天,今天去打渔\n", data)
	} else {
		fmt.Printf("距离1990-01-01,已经过去%v天,今天去晒网\n", data)
	}
	return data
}

func main() {
	var s string
	fmt.Println("请按照1990-01-01的格式输入时间")
	_, _ = fmt.Scanln(&s)
	GetTimeArr("1990-01-01", s)
}
