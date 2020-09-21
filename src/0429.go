package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main(){
	fmt.Printf("\n开始监听\n")

	http.HandleFunc("/login",func( writer http.ResponseWriter ,requst *http.Request){
		var na,name,numb string
		var value string
		var a,b,c,d bool
		 ur1 := requst.URL
		 requst.ParseForm()
		 _,fe:= requst.MultipartReader()
		if fe != nil {
			fmt.Println(fe)
		}
		fmt.Println("\n\nForm ->",requst.Form)
		 fmt.Println("host",ur1.Host)
		 fmt.Println("path",ur1.Path)
		 fmt.Println("header",ur1.Scheme)
		 fmt.Println("postFrom",requst.PostForm)
		 str := ur1.Path
		 x:=strings.Split(str,"/")
		 fmt.Printf("长度%d %s",len(x),x)
		 str = x[len(x)-1]  //获取截取的最后一段
		 a =strings.Contains(str,"name")
		 b =strings.Contains(str,"numb")
		 c =strings.Contains(str,"&")
		if !a&&!b&&!c{
			    fmt.Println("请求格式无效")
				writer.Write([]byte("请求格式无效"))
				return
		}
		fmt.Printf("\nstr %s",str)
		 str2 := strings.Split(str,"&")
		fmt.Printf("\nstr1 %s str2 %s",str2[0],str2[1])
		d =strings.Contains(str2[0],"=")
		if !d{
			fmt.Println("请求格式无效")
			writer.Write([]byte("请求格式无效"))
			return
		}
		na,value = str_split(str2[0])
		if na=="name"{
			name =value
		}else{
			numb =value
		}
		d =strings.Contains(str2[1],"=")
		if !d{
			fmt.Println("请求格式无效")
			writer.Write([]byte("请求格式无效"))
			return
		}
		na,value = str_split(str2[1])
		if na=="name"{
			name =value
		}else{
			numb =value
		}
		if name=="hello"&&numb=="123456"{
			writer.Write([]byte("恭喜登录成功"))
		 }else{
			writer.Write([]byte("用户不存在登录失败"))
		 }

	})
    err:=	http.ListenAndServe("127.0.0.1:8080",nil)
	fmt.Println("接收完成",err)
}

func str_split(st string)(string,string){
	a := strings.Split(st,"=")
	return a[0],a[1]
}