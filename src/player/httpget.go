package player

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type UP_date struct {
	UID string
	Name string
	Sex string
	Face string
	Follower string
	Following string
}
//==================================================================================
//名字: GetUpDate(mid string)UP_date
//功能: 通过uid获取对应up主的相关信息
func GetUpDate(mid string)UP_date{
	var url,url2 string
	UP:= UP_date{}
	var accept,referer,user_agent string
	var date,date2 string
	url = "https://api.bilibili.com/x/space/acc/info?mid=" + mid + "&jsonp=jsonp"
	url2 = "https://api.bilibili.com/x/relation/stat?vmid=" + mid + "&jsonp=jsonp&callbac"
	//===================================================================
	accept =  "application/json, text/plain, */*"
	referer = "https://space.bilibili.com/" + mid + "?spm_id_from=333.788.b_765f7570696e666f.1"
	user_agent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	date = GetHttpPage(url,accept,referer,user_agent)
	//===================================================================
	accept =  "application/json, text/plain, */*"
	referer = "https://space.bilibili.com/"+mid+"/fans/follow"
	user_agent = "ozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	date2 =  GetHttpPage(url2,accept,referer,user_agent)
	//正则表达式，筛选出想要的数据
	name := regexp.MustCompile(`name":"(.*?)"`)
	name2 := name.FindAllStringSubmatch(date, 1)
	if len(name2)==0{
		return  UP
	}
	UP.UID = mid
    UP.Name = name2[0][1]
	sex := regexp.MustCompile(`"sex":"(.*?)"`)
	UP.Sex = sex.FindAllStringSubmatch(date, 1)[0][1]
	face := regexp.MustCompile(`"face":"(.*?)"`)
	UP.Face= face.FindAllStringSubmatch(date,1)[0][1]
	follower := regexp.MustCompile(`follower":([\d]+)`)
	follower2 := follower.FindAllStringSubmatch(date2, 1)
	following := regexp.MustCompile(`following":([\d]+)`)
	UP.Following = following.FindAllStringSubmatch(date2, 1)[0][1]
	if len(follower2)>4{
		UP.Follower = follower2[0][1][:len(follower2)-4]+"万"
	}else{
		UP.Follower = follower2[0][1]
	}
	return  UP

}
//==================================================================================
//名字: GetHttpPage(url string,Accept string,Referer string,UserAgent string) string
//功能: 请求并获取对于网页的数据并解析返回
//==================================================================================
func GetHttpPage(url string,Accept string,Referer string,UserAgent string) string {
	client := http.Client{}
	requst,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
		fmt.Println("网页请求错误")
		log.Fatal(err)
		return "网页请求失败"
	}
	requst.Header.Add("Accept", Accept)
	requst.Header.Add("Referer", Referer)
	requst.Header.Add("User-Agent", UserAgent)

	response,err := client.Do(requst)
	if err!=nil{
		fmt.Println("运行请求错误")
		log.Fatal(err)
		return "运行请求失败"
	}
	htmlbyte,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("解析错误")
		log.Fatal(err)
		return "解析失败"
	}
	defer response.Body.Close()
	return string(htmlbyte)
}


