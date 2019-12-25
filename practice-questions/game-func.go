//编写一个函数，随机猜数游戏，随机生成一个1-100的整数，有十次机会
//如果第一次猜中，就提示"你真是个天才"
//如果第2-3次猜中，提示"你很聪明，赶上我了"
//如果第4-9次猜中，提示"一般般"
//如果最后一次猜中，提示"可算猜对了"
//一次都没猜中，提示"说你点啥好呢"
//分析：需要随机生成数字，并赋值给一个全局变量
//比较变量是否一致，输出对应的结果

package main

import (
	"fmt"
	"math/rand"
	"time"
)
//生成随机数字
func num(number int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(number)
}
func main()  {
	var result int
	p := num(100)
	fmt.Printf("结果为%v\n", p)
	for frequency :=1; frequency <=10; frequency++{
		fmt.Print("开始猜数字游戏，请输入:")
		_, _ = fmt.Scanf("%v", &result)
		if frequency ==1 && result == p {
			fmt.Println("你真是个天才")
			break
		}else if frequency >=2 && frequency <= 3 && result == p {
			fmt.Println("你很聪明，赶上我了")
			break
		}else if frequency >=4 && frequency <= 9 && result == p {
			fmt.Println("一般般")
			break
		}else if frequency == 10 && result == p {
			fmt.Println("可算猜对了")
			break
		}else if frequency ==10 && result !=p {
			fmt.Println("说你点啥好呢")
		}else {
			fmt.Printf("哎呀，猜错了，只有十次机会哦，你已经使用了%v次机会。\n", frequency)
		}
    }
}