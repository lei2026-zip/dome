package main

import "fmt"

func main(){
   var a int16 = 26
   var out int16 =0
   out = a%2 + a/2%2*10 + a/4%2*100 + a/8%2*1000+a/16%2*10000
   fmt.Printf("作业一  十进制转二进制:")
   fmt.Printf("%d %d \n",a,out)
    a=31
    out = 0
    out = a%8 + a/8%8*10
	fmt.Printf("作业二  十进制转八进制:")
	fmt.Printf("%d %d \n",a,out)
    a=1
    out =2
	fmt.Printf("作业三  运算符运用:")
	fmt.Printf(" a>out=%t a<out=%t a>=out=%t a<=out=%t a==out=%t a!=out=%t \n",a>out,out>a,a>=out,a<=out,a==out,a!=out)
	a = 4
	var  b int16 = 3
	res1 := a<b&&b/2==1&&a%3!=0
	res2 :=	(a+b)*3<a<<2||(a-b)>0
    fmt.Printf("作业四  a:=4,b:=3求表达式的值:\n")
	fmt.Printf("1.res1:=a<b&&b/2==1&&a余3!=0 %t \n",res1)
	fmt.Printf("2.res2:=(a+b)*3<a<<2||(a-b)>0 %t \n",res2)
}
//  %T  返回对象数据类型  \t  空格   