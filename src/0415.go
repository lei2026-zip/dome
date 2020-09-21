package main

import (
	"errors"
	"fmt"
)

type Person struct {
	name string
	age int
	address string
}

func JudgeAge(per Person) (bool,error){
	if per.age>=18{
       return true,nil
	}else if 0<per.age{
		return false,nil
	}else{
		var err error =errors.New("年龄不合法")
		return false,err
	}
}

func main(){
    fmt.Printf("\n作业一\n")
	a := Person{"王二狗",0,"南昌"}
	fmt.Println(a)
	b,c := JudgeAge(a)
	if c!=nil{
		fmt.Printf("%s",c)
	}else{
		if b==true {
			fmt.Printf("已成年")
		}else{
			fmt.Printf("未成年")
		}
	}
	//=============
	fmt.Printf("\n作业二\n")
	d := [6]int{1,2,3,4,5,6}
	for i:=0;i<10;i++{
		 if i>=len(d){
		 	panic("超出数组长度")
		 }
	}
}
