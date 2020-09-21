package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "http://192.168.10.100:8080/Day33_Servlet/aa.jpeg"
	fmt.Printf("\n作业一\n")
	fmt.Printf("%s文件路劲:",str)
	ind := strings.LastIndex(str,"/")
	str2 := str[ind+1:]
	ind = strings.LastIndex(str2,".")
	str3 := str2[ind+1:]
	str2 = str2[:ind]
	fmt.Printf("\n该路径文件名为:%s 扩展名为:%s",str2,str3)

	fmt.Printf("\n作业二\n")
    st :="hello hello hello word"
    times := strings.Count(st,"hello")
    fmt.Printf("\n%s \n该字符串\"heello\"出现的次数为:%d",st,times)

	fmt.Printf("\n作业三\n请输入一个年份:")
    var a int
    for i:=0;i<3;i++{
		fmt.Scanf("%d",&a)
		if huihui(a){
			fmt.Printf("该年是闰年\n")
		}else{
			fmt.Printf("该年不是闰年\n")
		}
	}
}

func huihui(a int)(bool,){
   if ((a%100!=0)&&(a%4==0))||a%400==0{
	   return true
   }else{
	   return false
   }
	
}