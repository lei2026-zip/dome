package main

import (
	"fmt"
)

func main(){
	var a,b int
	a = 3
	b= 3
	fmt.Println("\n作业1")
	fmt.Println("\n加减乘除：",a+b,a-b,b*a,a/b)
	var flo float32 = 3.14
	var bo bool = true
	str := "lei"
	fmt.Println("\n作业2")
	fmt.Printf("\n %f %t %s",flo,bo,str)
    fmt.Println("\n作业3")
	fmt.Printf("\n %d %d %d %d %d %d %d ",Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday)
	fmt.Println("\n作业4")
	var z,x,_,c int = 1,2,3,4
	fmt.Printf("\n 1 2 3 4")
	fmt.Printf("\n %d %d %d ",z,x,c)
	fmt.Println("\n作业5")
	fmt.Println("\n月份" ,January ,February, March, April, May ,June ,July ,August ,September, October, November,December)
	fmt.Println("\n作业6")
	fmt.Println("\n前面已定义a,b所以用大写形式来代替")
	fmt.Printf("\n A+B=%d",A+B)
}

const(
	Monday int8 = 1
	Tuesday int8 = 2
	Wednesday int8 = 3
	Thursday int8 = 4
	Friday int8 = 5
	Saturday int8 = 6
	Sunday int8  = 7

	January string = "一月"
	February string = "二月"
	March string = "三月"
	April string = "四月"
	May string = "五月"
	June string = "六月"
	July string = "七月"
	August string = "八月"
	September string = "九月"
	October string = "十月"
	November string = "十一月"
	December string = "十二月"

	A int8 = 3
	B int8 = 5
	)
