package main

import "fmt"

func main(){
   fmt.Printf("\n作业一\n")
   n:=9
   sun(n,n-1,1)

   fmt.Printf("\n作业二\n")
   m:=10
   func (a int){
   	out:=1
		for i:=1;i<=a;i++{
			out*=i
		}
		fmt.Printf("%d的阶乘为%d",a,out)
	}(m)
}



func sun(n int,m int,k int){
	for i:=1;i<=k;i++{
		fmt.Printf("%d",i)
	}
	fmt.Printf(" ")
	if m>0{
		k++
		sun(n,m-1,k)
	}
}