package main

import (
	"net/http"
)
//作业一===================================================
type router struct {
}

func (a *router)ServeHTTP(w http.ResponseWriter,r *http.Request){
	ur :=r.URL
	if ur.Path=="/userinfo"{
		w.Write([]byte("查询用户信息"))
		return
	}else if  ur.Path=="/student"{
		w.Write([]byte("查询学生信息"))
	}else{
		http.NotFound(w,r)
	}
	return
}

func main(){
	Rout := router{}
    http.ListenAndServe("192.168.1.108:8000",&Rout)
	//作业一=================================================================
	//作业二=================================================================
	http.Handle("/",http.FileServer(http.Dir("C:")))
	http.ListenAndServe("192.168.1.108:8000",nil)
}
