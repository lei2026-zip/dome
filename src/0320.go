package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Printf("作业一\n")
	var str string = "我爱你中国,I LOVE CHINA"
	strings.Split(str," ")
	fmt.Printf("长度%d\n",len(str))
    for _,value:=range str{
    	fmt.Printf("%c  ",value)
	}
	//===========================
	fmt.Printf("\n作业二\n")
    var a []byte = make([]byte,4,4)
    a =[]byte{97,98,99,100}
  //  ax := string(a)
    for i:=0;i<4;i++{
		fmt.Printf("%c->%d ",a[i],a[i])
	}
    //============================
	fmt.Printf("\n作业三\n")
    str2 := "WUHANJIAYOUZHONGGUOJIAYOU"
    var b bool = strings.Contains(str2,"ZHONGGUO")
    if b {
		fmt.Printf("存在 位置为%d", strings.Index(str2, "ZHONGGUO"))
	}else{
		fmt.Printf("不存在ZHONGGUO")
	}
}
