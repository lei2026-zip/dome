package main

import (
	"fmt"
	"strings"
)

func main(){
	//tolower转小写
	//toupper转大写
	//Trim去除开头和末尾指定内容
	//Trimspace 开头和末尾去空格
    //
	str := "    addwqdadascasdwfgbnjmyjtryd.abc"
	fmt.Printf("%s\n",str)
	str2 := strings.Replace(str,"a","b",10)
	a := strings.Count(str,"s")
	b :=strings.HasSuffix(str,"ryd")  //判断后戳
   fmt.Printf("%s\n %d %t",str2,a,b)
	str3 := strings.Split(str,"d")
	fmt.Printf("\n%s-%s-%s\n",str3[1],str3[2],str3[3])
	d := hah(3,2)
   fmt.Printf("%d",d)
 	e:="中国"
 	f :=[]byte(e)
	fmt.Printf("len=%d",f)
}
func hah(a int,b int)(int){
 var c int= a+b
	return  c
}