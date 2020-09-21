package main

import "fmt"

func main(){
  fmt.Printf("\n作业一\n")
  a:=[5]int{23,12,65,78,33}
   P := &a
   fmt.Printf("数组的地址%p \nP变量的类型%T",&a,P)
   fmt.Printf("\n作业二\n")
   var st [6]string = [6]string{"123","12","asd","bb","cc","dd"}
   var Pst [6]*string
   for i:=0;i<len(st);i++{
   	   Pst[i] = &st[i]
   }
   fmt.Printf("%p",Pst)
   *Pst[5]="ss"
   fmt.Printf("\n%s",st,)
   fmt.Printf("\n作业三\n")
   b :=10
   c :=22
   d := chang3(b,c)
   fmt.Printf("%d+%d=%d",b,c,*d)
}
func chang3(a int,b int) *int{
     var c int = a+b
     return &c
}