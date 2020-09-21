package main

import "fmt"

func main(){
	var day int = 0
	fmt.Printf("作业一 请输入一个日期：")
	fmt.Scanf("%d",&day)
	if day<=31 {
		switch day {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
			fmt.Printf("\n上旬\n")
		case 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
			fmt.Printf("中旬\n")
		default:
			fmt.Printf("下旬\n")
		}
	}else{
		fmt.Printf("输入的数据有误")
	}

	fmt.Printf("作业二 输出1-100中同时是6的倍数又是4的倍数的数：")
	for i:=0 ;i<100;i++{
		if (i%6==0)&&(i%4==0) {
			fmt.Printf("%d ",i)
		}
	}
	fmt.Printf("作业四 打印特殊图形：\n")
	var a,b,a1,b1 int =4,0,4,0
   for i:=0;i<9;i++{
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
	  if i<=3{
		  a--
		  b++
	  }else {
	  	b--
	  	a++
	  }
	  fmt.Printf("\n")
   }
}
