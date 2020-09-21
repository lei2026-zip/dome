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
	for {
		fmt.Printf("请输入要获取的UP主UID:(例如我的uid =344485144)\n")
		fmt.Scanln(&mid)
		fmt.Printf("\n稍等片刻...")

		herf2 := "https://api.bilibili.com/x/space/acc/info?mid=" + mid + "&jsonp=jsonp"
		client2 := http.Client{}
		requst2, err := http.NewRequest("GET",herf2, nil)
		requst2.Header.Add("Accept", "application/json, text/plain, */*")
		requst2.Header.Add("Referer", "https://space.bilibili.com/" + mid + "?spm_id_from=333.788.b_765f7570696e666f.1")
		requst2.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

		if err != nil {
			fmt.Printf("请求错误,请检查网络")
			log.Fatal(err)
			return
		}
		response2,err := client2.Do(requst2)
		if err != nil {
			log.Fatal(err)
			return
		}
		html2,err := ioutil.ReadAll(response2.Body)
		defer response2.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		htmlbyte2 :=string(html2)
		str2 := regexp.MustCompile(`name":"(.*?)"`)
		name := str2.FindAllStringSubmatch(htmlbyte2, 1)
		str3 := regexp.MustCompile(`"sex":"(.*?)"`)
		sex := str3.FindAllStringSubmatch(htmlbyte2, 1)
		str4 := regexp.MustCompile(`"face":"(.*?)"`)
		picture := str4.FindAllStringSubmatch(htmlbyte2, 1)
		herf := "https://api.bilibili.com/x/relation/stat?vmid=" + mid + "&jsonp=jsonp&callbac"
		client := http.Client{}
		requst, err := http.NewRequest("GET", herf, nil)
		if err != nil {
			fmt.Printf("请求错误,请检查网络")
			log.Fatal(err)
			return
		}
		requst.Header.Add("Accept", "application/json, text/plain, */*")
		requst.Header.Add("Referer", "https://space.bilibili.com/"+mid+"/fans/follow")
		requst.Header.Add("User-Agent", "ozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
		response, err := client.Do(requst)
		if err != nil {
			fmt.Printf("回应错误")
			log.Fatal(err)
			return
		}
		htmlbyte, err := ioutil.ReadAll(response.Body)
		defer  response.Body.Close()
		if err != nil {
			fmt.Printf("字符解析错误")
			log.Fatal(err)
			return
		}
		html := string(htmlbyte)
		str := regexp.MustCompile(`follower":([\d]+)`)
		date := str.FindAllStringSubmatch(html, 1)
		str = regexp.MustCompile(`following":([\d]+)`)
		date2 := str.FindAllStringSubmatch(html, 1)
		if date2==nil||sex==nil{
			fmt.Println()
			fmt.Printf(html,htmlbyte2)
			fmt.Printf("\n请求过程中发生意料之外的事...稍后请重试")
			fmt.Printf("\n获取失败,true回车->再来一次或输入false回车结束程序\n")
			fmt.Scanln(&i)
			if i==false{
				break
			}
			continue
		}
		sex2 := sex[0][1]
		name2 := name[0][1]
		picture2 := picture[0][1]
		fmt.Printf("\n name:%s sex:%s 头像:%s",name2,sex2,picture2)
		numb := date[0][1]
		numb2 := date2[0][1]

		if len(numb) > 4 {
			fmt.Printf("  关注数为: %s个 粉丝数为 %s万", numb2, numb[:len(numb)-4])
		} else {
			fmt.Printf("  关注数为: %s个 粉丝数为 %s", numb2, numb)
		}
		fmt.Printf("\n获取成功,true回车->再来一次或输入false回车结束程序\n")
        fmt.Scanln(&i)
		if i==false{
			break
		}
	}
}

