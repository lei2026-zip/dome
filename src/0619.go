package main

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"player"
)

//var index int
const name  = 23
const a  = iota
func main(){
	date,err := player.GetLoldate()
	if !err{
		fmt.Println("英雄信息获取错误")
		return
	}
	//========发送英雄皮肤主页=
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/login.html")
		if err!=nil{
			fmt.Println(err)
			return
		}
		temp.Execute(writer,"")
	})
	//这里发送相关英雄的技能详表
	http.HandleFunc("/heroSpells", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		i:= request.FormValue("id")
		ind,er :=forrangeID(date,i)
		//index = ind
		if !er{
			fmt.Println("检索失败，无此英雄ID",ind)
			return
		}
		temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/skills.html")
		if err!=nil{
			fmt.Println(err)
			return
		}
		temp.Execute(writer,date[ind])
	})
	//这里接收直接的请求，并发送本项目主页
	http.HandleFunc("/lol", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		i:= request.FormValue("id")
		p:= request.FormValue("password")
		str,er :=player.Switch_userpassword(i,p)
		if !er{
			writer.Write([]byte(str))
			fmt.Println("登入失败",str)
			return
		}else{
			fmt.Println("密码正确登入成功")
		}
		temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/herolist.html")
		if err!=nil{
			fmt.Println(err)
			return
		}
		temp.Execute(writer,date)
	})
	http.HandleFunc("/back", func(writer http.ResponseWriter, request *http.Request) {
		temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/herolist.html")
		if err!=nil{
			fmt.Println(err)
			return
		}
		temp.Execute(writer,date)
	})
	//这里接收并发送发送相关英雄相惜信息的请求
	http.HandleFunc("/herodate", func(writer http.ResponseWriter, request *http.Request) {
		database,err :=sql.Open("mysql","root:lch123456@tcp(127.0.0.1:3306)/mydate?charset=utf8")
		defer database.Close()
		if err != nil{
			fmt.Println("数据库打开失败0")
			log.Fatal(err.Error())
			return
		}
		request.ParseForm()
		//temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/hero.html")
		//if err!=nil {
		//	fmt.Println(err)
		//	return
		//}
		i:= request.FormValue("id")
		ind,er :=forrangeID(date,i)
		//index = ind
		if !er{
			fmt.Println("检索失败，无此英雄ID",ind)
			return
		}
		temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/hero.html")
		if err!=nil{
			fmt.Println(err)
			return
		}
		temp.Execute(writer,date[ind])
		//temp.Execute(writer,date[ind])
	})
	http.HandleFunc("/input", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		i:= request.FormValue("id")
		p:= request.FormValue("password")
		dat := player.User{i,p}
		err:= player.Addsave_date(dat)
		if !err{
			writer.Write([]byte("用户名已存在..."))
		}else {
			writer.Write([]byte("注册成功..."))
		}
	})
	http.Handle("/",http.FileServer(http.Dir("src/doc")))
	http.HandleFunc("/regeist", func (writer http.ResponseWriter, request *http.Request) {
		temp,err :=template.ParseFiles("C:/Users/lei17/Desktop/GOproject/awesomeProject/src/regeist.html")
		if err!=nil{
			fmt.Println(err)
			return
		}
		temp.Execute(writer,"")
	})
	fmt.Println("Every is ok ! You can do that !")
	er := http.ListenAndServe("127.0.0.1:8080",nil)
	if er!=nil{
		log.Fatal(er)
		return
	}
	return
}
//============================================
//通过Heroid的值检索英雄在英雄列表中的index值====
//============================================
func forrangeID(list []player.HeroList,i string)(int,bool){
	for index,value:= range list{
		if value.Hero.Heroid == i{
               return index,true
		}
	}
	return 0,false
}

