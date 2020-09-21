package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	requst := "https://www.baidu.com"
	clien,err := http.Get(requst)
	if err!=nil{
		fmt.Printf("http get err %s",err)
		return
	}
	byte,err2 := ioutil.ReadAll(clien.Body)
	if err2!=nil{
		fmt.Printf("ioutil.ReadAll err %s",err2)
		return
	}
	fmt.Println(clien.Body)
	fmt.Printf("%s",byte)

}
