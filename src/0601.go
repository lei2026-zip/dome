package main

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	fmt.Println("欢迎进入豆瓣电影爬虫系统")
	//var page int
	//fmt.Println("请输入你要采集的页码数字（比如2代表采集前2页,请保持在10以内)：")
	//_,err :=fmt.Scan(&page)
	//if err != nil{
	//	log.Fatal(err.Error())
	//	return
	//}
	//if page > 10{
	//	fmt.Println("只能采集前10页数据，我尽力了")
	//	os.Exit(0)
	//	return
	//}

	//url: https://movie.douban.com/top250

	/**
	  第一步：确定目标url，进行网络请求，获取网页
	*/
	client := http.Client{}
	//英文单词：filter：过滤器
	//start表示的是：展示过了多少条数据
	//一共有250条数据，25条是一夜，一共10页
	//https://movie.douban.com/top250?start=100&filter=
	//例如：取100条数据  一共有：100 / 25 = 4页
	// 取110条数据 一共有：110 / 25 = 4 ... 10(1页)  共有5页
	//movieNum := 100 //100条
	//page := movieNum / 25
	//if movieNum % 25 !=0{
	//	page++ //页码加1
	//}


//	for a := 0; a < page; a++ {//页数
	//	fmt.Printf("采集第 %d页数据\n",a)
	//	offsetNum := strconv.Itoa(a * 25)
		request,err := http.NewRequest("GET","https://movie.douban.com/top250?start=0&filter=",nil)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		//添加请求头，模仿浏览器
		//可以接受什么格式的数据
		request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

		//可以接受的语言
		request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,vi;q=0.7")
		request.Header.Add("Cache-Control","no-cache")//缓存的设置：不适用缓存
		request.Header.Add("Connection","keep-alive")//连接：保持连接
		request.Header.Add("Host","movie.douban.com")
		request.Header.Add("Pragma","no-cache")
		request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")

		//请求数据
		response,err := client.Do(request)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		htmlBytes,err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		//通过网络请求把网页数据转换成html的string格式，然后输出
		html := string(htmlBytes)//网页源码

		/**
		  第二步：从网页源码中剥离数据，筛选符合规则的数据，进行爬取
		    2.1确定爬取的内容：电影编号，封面、电影名字、电影的评分、电影评价人数、一句话描述
				电影编号：<a href="https://movie.douban.com/subject/ + 编号 + /">
				电影名字：<img width="100" alt=" + 电影名称 +"
				电影封面图：src标签 src="封面图片地址url" class="">
				电影评分：<span class="rating_num" property="v:average"> + 评分 + </span>
		*/

		//编号
		movieIdReg := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/">`)
		idSlice := movieIdReg.FindAllStringSubmatch(html,-1)

		//(.*?)代表的是：任意的字符串，任意的长度都会被匹配到
		nameReg := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
		nameSlice := nameReg.FindAllStringSubmatch(html,-1)//返回值：2维切片，[][]string
		//fmt.Println(nameSlice)

		//评分
		ratReg := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
		ratSlice := ratReg.FindAllStringSubmatch(html,-1)//-1代表获取html数据中的所有符合规则的数据

		//评价人数：投票vote
		voteReg := regexp.MustCompile(`<span>(.*?)人评价</span>`)
		voteSlice := voteReg.FindAllStringSubmatch(html,-1)

		//一句描述：
		descReg := regexp.MustCompile(`<span class="inq">(.*?)</span>`)
		descSlice := descReg.FindAllStringSubmatch(html,-1)

		//封面
		coverImgReg := regexp.MustCompile(`src="(.*?) class="">`)
		imgSlice := coverImgReg.FindAllStringSubmatch(html,-1)

		fmt.Println("电影编号   电影名称  评分  评价人数  一句话总结  封面图")
		for i :=0; i < len(nameSlice); i++{//当前页面电影的个数
			fmt.Printf("%s  %s  %s  %s %s %s \n",
				idSlice[i][1],
				nameSlice[i][1],
				ratSlice[i][1],
				voteSlice[i][1],
				descSlice[i][1],
				imgSlice[i][1])
		}

		/**
		  第三部：把解析爬取到的数据存储到mysql数据库
			3.1 创建数据库，并创建数据表（电影信息表）
				注意事项：每一条电影应该是唯一的，通过电影 “编号” 区分电影记录
		    3.2 使用go语言，编写程序实现数据库的连接：douban_movie
		    3.3 执行插入语句：将获取到数据存入到数据库表当中
		*/
		database,err :=sql.Open("mysql","root:lch123456@tcp(127.0.0.1:3306)/mydate?charset=utf8")
		defer database.Close()
		if err != nil{
			log.Fatal(err.Error())
			return
		}
		//执行sql语句：插入数据
		for j :=0; j < len(idSlice); j++{
			_, err := database.Exec("insert into " +
				"movie(movie_id, movie_name, movie_rate, movie_vote, movie_desc, movie_url)" +
				"values(?,?,?,?,?,?)",
				idSlice[j][1],
				nameSlice[j][1],
				8.0,
				voteSlice[j][1],
				descSlice[j][1],
				imgSlice[j][1])
			if err != nil {
				log.Fatal(err.Error())
				break
			}
		}
		fmt.Println()
//	}

	fmt.Println("豆瓣电影信息采集完成！")
}
