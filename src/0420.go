package main

import (
	"fmt"
	"sync"
	"time"
)

var number int = 100
var all sync.WaitGroup
var and sync.RWMutex

func main(){
    all.Add(3)
    go num1()
	go num2()
	go num3()
    all.Wait()
}

func num1(){
	for i:=0;i<100;i++ {
		and.RLock()
		if number<=0{
			and.RUnlock()
			break
		}
		time.Sleep(100*time.Millisecond)
		number--
		fmt.Printf("RL 今天丢了一块钱 还剩%d元;\n",number)
		and.RUnlock()
	}
    all.Done()
}

func num2(){
	for i:=0;i<100;i++ {
		and.RLock()
		if number<=0{
			and.RUnlock()
			break
		}
		time.Sleep(100*time.Millisecond)
		number--
		fmt.Printf(" RL还剩%d元;\n",number)
		and.RUnlock()
	}
   all.Done()
}

func num3(){
	for i:=0;i<100;i++ {
		and.Lock()
		if number<=0{
			and.Unlock()
			break
		}
		time.Sleep(200*time.Millisecond)
		number--
		fmt.Printf("-L 还剩%d元;\n",number)
		and.Unlock()
	}
   all.Done()
}