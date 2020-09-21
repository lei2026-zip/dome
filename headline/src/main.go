package main

import (
	"DateStuct"
	"fmt"
	"net/http"
	"postdate"
	"servehttp"
)

var str []string = []string{"top","shehui","guonei","guoji","yule","tiyu","junshi","keji","caijing","shishang"}
var NewsList []DateStuct.News

func main(){
	fmt.Println("running Program... ")
    NewsList,err :=  postdate.Get_All_Headline(str)
    if !err{
    	fmt.Println("数据获取失败！")
	//	return
	}
	fmt.Println(NewsList)
	http.Handle("/",http.FileServer(http.Dir("./html")))
    http.HandleFunc("/toutiao", func(writer http.ResponseWriter, request *http.Request) {
		servehttp.Servehtml(writer,NewsList)
	})
    http.ListenAndServe("127.0.0.1:8080",nil)
}
