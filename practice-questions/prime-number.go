//编写一个函数，输出100以内的所有素数（素数是只能被1和本身整除的数）每行显示5个，并求和

package main

import "fmt"

// 判断n是不是质数
func Prime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//定义一个切片
	var data [] int
	// 求100以内的所有质数,并添加到切片data
	for i := 2; i <= 100; i++ {
		if Prime(i) == true {
			data = append(data,i)
		}
	}
	//输出切片内容
	fmt.Println(data)
    data01 := data[0:5]
	data02 := data[5:10]
	data03 := data[10:15]
	data04 := data[15:20]
	data05 := data[20:]
	fmt.Println("第一组数据为：",data01)
	var a int
	for _,value := range data01{
		a += value
	}
	fmt.Println("第一组的和为：",a)
	fmt.Println("第二组数据为：",data02)
	for _,value := range data02{
		a += value
	}
	fmt.Println("第二组的和为：",a)
	fmt.Println("第三组数据为：",data03)
	for _,value := range data03{
		a += value
	}
	fmt.Println("第三组的和为：",a)
	fmt.Println("第四组数据为：",data04)
	for _,value := range data04{
		a += value
	}
	fmt.Println("第四组的和为：",a)
	fmt.Println("第五组数据为：",data05)
	for _,value := range data05{
		a += value
	}
	fmt.Println("第五组的和为：",a)

}
