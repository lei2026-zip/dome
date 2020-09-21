package main

import (
	"fmt"
	"sync"
	"time"
)
var ch chan int
var wit sync.WaitGroup
var cl sync.Mutex
var stock int = 200

func main(){
	ch = make(chan int)
    wit.Add(2)
   go sale1()
    go sale2()
    wit.Wait()
}


func sale1(){
	for i:=0;i<1000;i++{
		 time.Sleep(2*time.Millisecond)
         stock =<- ch
		fmt.Printf("\n 当前库存%d",stock)
	}
	wit.Done()
}
func sale2(){
	for i:=0;i<1000;i++{
		time.Sleep(2*time.Millisecond)
        stock++
        ch <- stock
		fmt.Printf("\n 当前库存%d",stock)
	}
	wit.Done()
}