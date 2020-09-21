package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.Handle("/",http.FileServer(http.Dir("C:/Users/lei17/Desktop/HTML系列/记事框2.html")))
	http.HandleFunc("/w",func(w http.ResponseWriter,re *http.Request){
		fmt.Printf("\n已发送回复")
		w.Write([]byte("<!DOCTYPE html><html ><head><meta charset=\"UTF-8\"><title>Title</title></head><body><form action=\"http://127.0.0.1:8080/login\" method=\"post\" > 用户名<input type=\"text\" name=\"username\" ><br> 密码<input type=\"password\" name=\"password\"><br><br><input type=\"submit\" value=\"登录\"></body></html>")) })
    http.
	http.ListenAndServe("127.0.0.1:8090",nil)
}


