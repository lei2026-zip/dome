package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main(){
	//我的uid = 344485144
	fmt.Printf("\n\n欢迎来到粉丝数获取界面 \n\n")
	var mid string
	var i bool
	var url,url2 string
	var accept,referer,user_agent string
	var date,date2 string
	for {
		fmt.Printf("请输入要获取的UP主UID:(例如我的uid =344485144)\n")
		fmt.Scanln(&mid)
		fmt.Printf("\n稍等片刻...")
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
		sex := regexp.MustCompile(`"sex":"(.*?)"`)
		sex2 := sex.FindAllStringSubmatch(date, 1)
		face := regexp.MustCompile(`"face":"(.*?)"`)
		face2 := face.FindAllStringSubmatch(date,1)
		follower := regexp.MustCompile(`follower":([\d]+)`)
		follower2 := follower.FindAllStringSubmatch(date2, 1)
		following := regexp.MustCompile(`following":([\d]+)`)
		following2 := following.FindAllStringSubmatch(date2, 1)

		if name2==nil{
			fmt.Println(date,date2)
			fmt.Printf("\n数据筛选失败.\n")
			continue
		}
		if len(following2)>4{
			fmt.Printf("\n\n name:%s sex:%s 粉丝follower:%s 关注following:%s万 face:%s",name2[0][1],sex2[0][1],follower2[0][1],following2[0][1][:len(following2)-4],face2[0][1])
		}else{
			fmt.Printf("\n\n name:%s sex:%s 粉丝follower:%s 关注following:%s face:%s",name2[0][1],sex2[0][1],follower2[0][1],following2[0][1],face2[0][1])
		}

		fmt.Printf("\n获取成功,true回车->再来一次或输入false回车结束程序\n")
		fmt.Scanln(&i)
		if i==false{
			break
		}
	}
}

//==================================================================================
//名字: GetHttpPage(url string,Accept string,Referer string,UserAgent string) string
//功能: 请求并获取对于网页的数据并解析返回
//==================================================================================
func GetHttpPage(url string,Accept string,Referer string,UserAgent string) string{
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