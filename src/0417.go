package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var delay sync.WaitGroup

func main(){
	rand.Seed(time.Now().Unix())
	fmt.Printf("\n作业一\n")
	delay.Add(2)
	go num_one()
	go num_two()
	delay.Wait()
	fmt.Printf("\n作业二\n")
	delay.Add(3)
	go num_one()
	go num_one()
	go num_one()
	delay.Wait()
}

func num_one()  {
	for i:=0;i<99;i++{
		time.Sleep(200)
		fmt.Printf("%d\n",rand.Intn(1000))
	}
    delay.Done()
}

func num_two(){
	for i:=0;i<199;i++{
		time.Sleep(100)
		fmt.Printf("%c\n",rand.Intn(26)+65)
	}
	delay.Done()
}
