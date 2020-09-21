package postdate

import (
	"DateStuct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*//http://v.juhe.cn/toutiao/index?type=all&key=bb5373e8ead75c7ef2a932d401fe3e4b*/

func Get_Headline(class string)([]DateStuct.List,bool){
	var DATA DateStuct.News
	DATA.Result.Class = class
	var key string = "bb5373e8ead75c7ef2a932d401fe3e4b"
	var url string = "http://v.juhe.cn/toutiao/index?"
	herf :=  url+"type="+class+"&key="+key
	byte,err:= http.Get(herf)
	if err!=nil{
		fmt.Println(err)
		return DATA.Result.Data,false
	}
	date,err:=ioutil.ReadAll(byte.Body)
	if err!=nil{
		fmt.Println(err)
		return DATA.Result.Data,false
	}
	err=json.Unmarshal(date,&DATA)
	if err!=nil{
		fmt.Println(err)
		return DATA.Result.Data, false
	}
	fmt.Printf("%T",DATA)
	if DATA.Error_code!=0{
		fmt.Println("Error_code = ")
		fmt.Println( DATA.Error_code)
		return DATA.Result.Data, false
	}
	return DATA.Result.Data,true
}

func Get_All_Headline(str []string)([]DateStuct.News,bool){
	var DATA []DateStuct.News
	var new DateStuct.News
	var err bool
	lenth := len(str)
	for i:=0;i<lenth;i++ {
		new.Result.Class = str[i]
		new.Result.Data,err= Get_Headline(str[i])
		if !err{
			fmt.Println("获取失败序列")
			fmt.Println(i,str[i])
			continue
		}
		DATA  =	append(DATA,new)
	}
	if len(DATA)==0{
		return DATA,false
	}
	return DATA,true
}