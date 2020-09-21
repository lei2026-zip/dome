package main

import "fmt"

func main(){
	 //var flo float64 = 48.14141
     //var str string = "qwetyutioioo"
     //fmt.Printf("%.2f %d %s",flo,int(flo),string(int(flo)))
     //fmt.Printf("  %s%s",str[0:3],str[3:len(str)])
     a := 12
     b := 5
     fmt.Print("Number one \n")
     fmt.Printf("a=%d b=%d 加%d 减%d 乘%d 商%d 余%d \n",a,b,a+b,a-b,a*b,a/b,a%b)
     fmt.Print("Number two \n")
     var str string ="wuhanjiayou,zhongguojiayou"
     fmt.Printf("%s\n %s %s %s %s \n",str,str[:5],str[12:20],str[:11],str[12:])
     fmt.Print("Number three \n")
     var G float32 = 6.67259
     fmt.Printf("%f %.2f %.3f",G,G,G)
}