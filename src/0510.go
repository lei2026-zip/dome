package main

import (
	"fmt"
	"math"
	"net/http"
)
func string_to_number(a []string) int{
	var len = len(a[0])
	var num int = 0
	//b:=strings.Split(a,"")
	for i:=0;i<len;i++{
		num = (int(a[0][i])-48)*int(math.Pow10(int(len-i-1))) + num
	}
	return num
}
func serve(w http.ResponseWriter,r *http.Request){
    url := r.URL
    r.ParseForm()
	date:=r.Form
	if len(date)==0{
		return
	}
    fmt.Printf("path:%s\n",url.Path)
    fmt.Printf("scheme:%s\n",url.Scheme)
    fmt.Printf("from:%s\n",r.Form)
    brith_year := string_to_number(date["brith_year"])
	brith_month := string_to_number(date["brith_month"])
	brith_day := string_to_number(date["brith_day"])
    if len(date["username"][0])>12{
         w.Write([]byte("姓名不符合规范,请重新填写"))
	}else if brith_year>2020||(brith_year==2020&&brith_month>5)||(brith_year==2020&&brith_month==5&&brith_day>9){
		w.Write([]byte("太着急了,还没到日子呢"))
	}else if len(date["ID"][0])!=18{
		w.Write([]byte("身份证长度不符合"))
	}else{
		w.Write([]byte("恭喜您注册成功!"))
	}
	fmt.Printf("\n姓名: %s",date["username"])
	fmt.Printf("\n出生日期 %d年%d月%d日",brith_year,brith_month,brith_day)
}
func main(){
	fmt.Printf("服务器开始运行")
	http.HandleFunc("/",serve)
	 err :=http.ListenAndServe("127.0.0.1:8080",nil)
	 if err!=nil{
	 	fmt.Printf("ListenAndServe:%s",err)
		 return
	 }
	return
}
