package main

import (
	"fmt"
	_ "go-sql-driver/mysql"
	"html/template"
	"net/http"
	"player"
)



func getdate(w http.ResponseWriter,r *http.Request){
	url := r.URL
	r.ParseForm()
	date:=r.Form
	if len(date)==0{
		w.Write([]byte("请求为空"))
		return
	}
	fmt.Printf("path:%s\n",url.Path)
	fmt.Printf("scheme:%s\n",url.Scheme)
	fmt.Printf("from:%s\n",r.Form)
	id := date["id"]  //获取输入的UID
	if len(id)==0{
		fmt.Println("空")
		return
	}
	data := []player.MovieDate{}
	var er bool
	data ,er = player.GetDouBan_date()  //此处获取对应UP的信息
	if !er{
		fmt.Println("douban信息存储或获取错误2")
		w.Write([]byte("douban信息存储或获取错误2...."))
		return
	}
	//fmt.Println(data)
	temp,err:= template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/get.html")
	if err!=nil{
		w.Write([]byte("解析文件出现错误...."))
		return
	}
	temp.Execute(w,data)
}
//func login(w http.ResponseWriter,r *http.Request){
//	url := r.URL
//	r.ParseForm()
//	date:=r.Form
//	if len(date)==0{
//		return
//	}
//	fmt.Printf("path:%s\n",url.Path)
//	fmt.Printf("scheme:%s\n",url.Scheme)
//	fmt.Printf("from:%s\n",r.Form)
//	id :=  date["id"]
//	password := date["password"]
//	fmt.Println(id,password)
//	temp,err:= template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/get.html")
//	if err!=nil{
//		w.Write([]byte("解析文件出现错误...."))
//		return
//	}
//	temp.Execute(w,"")
//}

func main(){
	fmt.Printf("服务器开始运行")
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		temp,err:= template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/login.html")
		if err!=nil{
			writer.Write([]byte("解析文件出现错误...."))
			return
		}
		temp.Execute(writer,"")
	})
	http.HandleFunc("/login",getdate)

	// err := http.ListenAndServe("127.0.0.1:8080",nil)
	s := http.Server{
		Addr:            "127.0.0.1:8090",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err:= s.ListenAndServe()
	defer s.Close()
	if err != nil{
		fmt.Printf("ListenAndServe:%s",err)
		return
	}
	return
}


