package main

import (
	"fmt"
	"html/template"
	"net/http"
	"player"
)



func Getdate(w http.ResponseWriter,r *http.Request){
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
	id := date["uid"]  //获取输入的UID
	if len(id)==0{
		fmt.Println()
		return
	}
	data := player.GetUpDate(id[0])  //此处获取对应UP的信息
	fmt.Println(data)
	temp,err:= template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/get.html")
     if err!=nil{
		w.Write([]byte("解析文件出现错误...."))
		return
	 }
	if data.Name==""{
		temp.Execute(w,"抱歉无该用户...")
	}else {
		str := "UID:"+string(data.UID)+" 名字:"+string(data.Name)+" 性别:"+string(data.Sex)+" 头像链接:"+string(data.Face)+" 关注:"+string(data.Following)+" 粉丝:"+string(data.Follower)
		temp.Execute(w,str)
	}
}
func login(w http.ResponseWriter,r *http.Request){
	url := r.URL
	r.ParseForm()
	date:=r.Form
	if len(date)==0{
		return
	}
	fmt.Printf("path:%s\n",url.Path)
	fmt.Printf("scheme:%s\n",url.Scheme)
	fmt.Printf("from:%s\n",r.Form)
	id :=  date["id"]
	password := date["password"]
	fmt.Println(id,password)
	temp,err:= template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/get.html")
	if err!=nil{
		w.Write([]byte("解析文件出现错误...."))
		return
	}
	temp.Execute(w,"")
}

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
	http.HandleFunc("/login",login)
	http.HandleFunc("/UID",Getdate)

	// err := http.ListenAndServe("127.0.0.1:8080",nil)
	s := http.Server{
		Addr:            "127.0.0.1:8080"  ,
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


