package player

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
	"log"
)

type User struct {
	Name string
	Password string
}

func Create_table(database *sql.DB)bool{
	//检查数据库是否已拥有此数据表
	row:=database.QueryRow("select COUNT(lol_user.index_id) count from lol_user")
	var count int
	err2 := row.Scan(&count)
	fmt.Println("==============================================================================================================================")
	fmt.Println(count)
	if err2 != nil {
		_, err :=database.Exec("create table "+"lol_user"+"( index_id int not null primary key auto_increment,name varchar(20),user_password varchar(20))"+"charset=utf8;")
		if err != nil{
			fmt.Println("数据库操作执行失败")
			log.Fatal(err.Error())
			return false
		}
		return false
	}else{
		fmt.Println("表格已存在")
		return true
	}
}

func Addsave_date(date User)bool{
	database,err :=sql.Open("mysql","root:lch123456@tcp(127.0.0.1:3306)/mydate?charset=utf8")
	defer database.Close()
	if err != nil{
		fmt.Println("数据库打开失败")
		log.Fatal(err.Error())
		return false
	}
	Create_table(database)
	qur,err2:= database.Query("select *from lol_user where name")
	if err2 != nil {
		log.Fatal(err.Error())
		return false
	}
	var dat []string
	str := ""
	for {
		if qur.Next(){
			qur.Scan(&str)
			dat = append(dat,str)
		}else{
			break
		}
	}
	for i:=0;i<len(dat);i++{
		if dat[i]==date.Name{
			fmt.Println(date.Name,"该用户名已存在")
			return false
		}
	}
	_,err1 := database.Exec("insert into " +
		"lol_user(name,user_password)"+
		"values(?,?)",date.Name,date.Password)
	if err1 != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}

func Switch_userpassword(user string,password string)(string,bool){
	database,err :=sql.Open("mysql","root:lch123456@tcp(127.0.0.1:3306)/mydate?charset=utf8")
	defer database.Close()
	if err != nil{
		fmt.Println("数据库打开失败")
		log.Fatal(err.Error())
		return "错误",false
	}
    er :=	 Create_table(database)
    if !er{
    	return "暂无任何用户",false
	}else{
		qur,err2:= database.Query("select *from lol_user ")
		if err2 != nil {
			log.Fatal(err.Error())
			return "错误",false
		}
		var dat []string
		var pas []string
		var id string
		str,str2 := "",""
		for {
			if qur.Next(){
				qur.Scan(&id,&str,&str2)
				dat = append(dat,str)
				pas = append(pas,str2)
			}else{
				break
			}
		}
		fmt.Println(dat,pas)
		for i:=0;i<len(dat);i++{
			if dat[i]==user{
				if pas[i]==password{
					return "密码正确",true
				}else {
					return "密码错误",false
				}
			}
		}
		return "用户不存在",false
	}
}