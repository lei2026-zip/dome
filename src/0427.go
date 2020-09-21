package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 chan int
var ch2 chan int
var c,d bool = false,false
func main(){
	fmt.Printf("\n作业一\n")
	ch1 = make(chan int,5)
	ch2 = make(chan int,5)
	e:=0
	go ran()
	go read()
	for{
       e=<-ch2
       fmt.Printf("  其三次方为%d\n",e)
       if c&&d{
       	  break
	   }
	}
	fmt.Printf("程序结束")
}

func ran()  {
	rand.Seed(time.Now().Unix())
   var a int
	for i:=0;i<100 ;i++  {
		a= rand.Intn(1000)
		time.Sleep(20*time.Millisecond)
        ch1 <- a
	}
	d = true
}

func read(){
	a:=0
	for{
		a=<-ch1
		fmt.Printf("随机数为%d",a)
		ch2 <-a*a*a
		if len(ch1)==0&&d{
			break
		}
	}
	c = true
}

