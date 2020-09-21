package main

import "fmt"

func main(){
	fmt.Printf("作业一:\n")
     for i:=0;i<6;i++{
     	for k:=0;k<6-i;k++ {
			fmt.Printf(" ")
		}
     	for j:=0;j<9;j++{
     		fmt.Printf("*")
		}
       fmt.Printf("\n")
	 }
	fmt.Printf("作业二:\n")
     k:=0
     for i:=0;i<10;i++{
     	k++
         for j:=0;j<k;j++{
         	if j==0||j==k-1||i==9{
               fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			 }
		 }
		 fmt.Printf("\n")
	 }

	fmt.Printf("作业三-实心正三角形:\n")
	var a,b,a1,b1 int =4,0,4,0
	for i:=0;i<5;i++{
		a1=a
		b1=b*2+1
		for j:=0;j<9;j++{
			if a1>0 {
				fmt.Printf(" ")
				a1--
			} else if b1 > 0 {
				fmt.Printf("*")
				b1--
			}
		}
			a--
			b++
		fmt.Printf("\n")
	}

	fmt.Printf("作业四-1-1000水仙花数：\n")
	x,y,z:=0,0,0
	for i:=0;i<1000;i++{
		x=i/100
		y=i%100/10
		z=i%100%10
		if z*z*z+x*x*x+y*y*y==i{
			fmt.Printf("%d \n",i)
		}
	}

}
