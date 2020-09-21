package servehttp

import (
	"DateStuct"
	"fmt"
	"html/template"
	"net/http"
)


func Servehtml(w http.ResponseWriter,data []DateStuct.News){
	temp,err :=template.ParseFiles("./html/main.html")
	if err!=nil{
		fmt.Println(err)
		return
	}
	temp.Execute(w,data)
}
