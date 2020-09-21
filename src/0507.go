package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Printf("开始运行")
	http.HandleFunc("/userlogin",func(w http.ResponseWriter,re *http.Request){
		fmt.Printf("\n已发送回复")
		w.Write([]byte("欢迎访问用户登录功能"))
	})
   err:= http.ListenAndServe("127.0.0.1:9001",nil)
	if err!= nil{
		fmt.Println(err)
	}

}
