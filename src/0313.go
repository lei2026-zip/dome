package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Printf("作业一_冒泡\n")
	 var b []int = []int{5,9,20,120,450,340}
	 fmt.Printf("%d\n",b)
	 for i:=0;i<len(b)-1;i++{
	 	for j:=0;j<len(b)-1;j++{
	 		if b[j]<b[j+1] {
	 			b[j],b[j+1]=b[j+1],b[j]
			}
		}
		 fmt.Printf("%d\n",b)
	 }
	fmt.Printf("作业二_选择\n")
	var a []int = []int{5,9,20,120,450,340}
	fmt.Printf("%d\n",a)
	for i:=0;i<len(a)-1;i++ {
		for j := i + 1; j < len(b); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
		fmt.Printf("%d\n", a)
	}
	fmt.Printf("作业三_随机数组\n")
	var c [10]int
	times :=0
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	for i:=0;i<10;i++{
		c[i] = rand.Intn(15)
		times++
		for j:=0;j<i;j++{
            if c[i]==c[j]{
            	i--
            	break
			}
		}
	}
	fmt.Printf("结果%d 获取次数%d",c,times)
	fmt.Printf("\n作业四_创建切片\n")
    var d []int=make([]int,5,7)
	 for i:=0;i<len(d);i++{
		 d[i] = rand.Intn(10)
	 }
	 fmt.Printf("%d 长度%d 容量%d",d,len(d),cap(d))
}