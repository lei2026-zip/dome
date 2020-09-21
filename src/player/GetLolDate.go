package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
var heroLi []HeroList
//存储后续获取的详细英雄信息
type HeroList struct {
	Hero Herodata
	Skins []Skin   //皮肤系列
	Spells []Spell //全部技能
}
//获取最初用来检索每个英雄的信息
type heroName struct {
	Hero []Herodata
	Version string
	FileTime string
}

//英雄信息集合
type Herodata struct {
	Heroid string
	Name string
	Alias string
	Title string
	Roles []string
	IsWeekFree string
	Face string
//	Date string
	ShortBio string //英雄人物志
	Allytips []string  //所以英雄技巧
    Enemytips []string //攻击技巧
    SelectAudio string  //声音
}
//皮肤集合
type a struct {
 a string

}
type Skin struct {
    SkinId string
    HeroId string
    Name string
    Chromas string
    Description string
    IconImg  string
    LoadingImg string
    SourceImg string
    VideoImg string
}
//技能集合
type Spell struct {
	Name string
	SpellKey string
	AbilityIconPath string
	Description string
	DynamicDescription string
}

//回调函数
func visit_herolist(vis []HeroList,f func(HeroList)){
	for _,value:= range vis{
		f(value)
	}
}


func GetLoldate() ([]HeroList,bool){
	fmt.Println("welcome to the league of legends")
	url:="https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
	herodata := Geturldate(url)
	if herodata == nil{
		fmt.Println("连接失败,请检查网络")
		return  heroLi,false
	}
	//解析传输过来的字符串并贮存到相应的结构体中
	var hero heroName
	err4:=json.Unmarshal(herodata,&hero)
	if err4 != nil {
		fmt.Println(err4.Error())
		fmt.Println("解析错误")
		return  heroLi,false
	}
	for i:=0;i<len(hero.Hero);i++{
		var herolist HeroList
		herolist.Hero = hero.Hero[i]
		herolist.Hero.Face = "http://game.gtimg.cn/images/lol/act/img/champion/"+hero.Hero[i].Alias+".png"
		ur2 := "https://game.gtimg.cn/images/lol/act/img/js/hero/"+hero.Hero[i].Heroid+".js"

		herodata = Geturldate(ur2)
		if herodata == nil{
			fmt.Println("herodata2 为空")
			return  heroLi,false
		}
		err :=json.Unmarshal(herodata,&herolist)
		if err != nil {
			fmt.Println("错误1"+err.Error())
			return  heroLi,false
		}
		heroLi = append(heroLi, herolist)
	}
	//因为获取的第一个英雄安妮的技能表里有一个多余的空的被动技能结构体所以在这里把它去的掉,
	//不知道其他英雄技能信息有没有这种情况 目前只发现这一个个
	//==============================================================================================
	visit_herolist(heroLi, func(list HeroList) {
		fmt.Println(list)
	})
	for i:=0;i<len(heroLi);i++{
		heroLi[i].Spells = Clear_emply_spell(heroLi[i].Spells)
		heroLi[i] = Change_Hero_order(heroLi[i])
		heroLi[i].Hero = translate_roles(heroLi[i].Hero)
		heroLi[i].Skins = clear_emply_skins(heroLi[i].Skins)
	}
	//==============================================================================================
	return  heroLi,true
}
func Clear_emply_spell(Hero []Spell) []Spell{
    for index,value:=range Hero{
    	if value.Name ==""{
    		if index<len(Hero)-1{
				Hero = append(Hero[:index],Hero[index+1:]...)
			}else {
				Hero = append(Hero[:index])
			}
		}
	}
	return Hero
}
//==================================================================================
//名字: Change_Hero_order(Hero HeroList) HeroList
//功能: 把原来获取的的技能顺序 e p q r w 转换成 p q w e r 的顺序
//==================================================================================
func Change_Hero_order(Hero HeroList) HeroList{
	if len(Hero.Spells)<5{
         fmt.Println("长度不够")
         fmt.Println("======================%s",Hero.Hero.Name)
         return Hero
	}
	Hero.Spells[1],Hero.Spells[2] = Hero.Spells[2],Hero.Spells[1]
	Hero.Spells[3],Hero.Spells[4] = Hero.Spells[4],Hero.Spells[3]
	Hero.Spells[2],Hero.Spells[3] = Hero.Spells[3],Hero.Spells[2]
	Hero.Spells[0],Hero.Spells[3] = Hero.Spells[3],Hero.Spells[0]
	return  Hero
}
//==================================================================================
//名字: Geturldate(url string) []byte
//功能: 请求并返回对应网址的数据
//==================================================================================
func Geturldate(url string) []byte{    // 返回为 ascll码集 必要时需用string()转化
	client:=http.Client{}
	request,err1:=http.NewRequest("GET",url,nil)
	if err1 != nil {
		fmt.Println("错误2",request)
		return nil
	}
	response,err2:=client.Do(request)
	if err2 != nil {
		fmt.Println("错误3",request)
		return nil
	}
	heroData,err3:=ioutil.ReadAll(response.Body)
	if err3 != nil {
		fmt.Println("错误4",request)
		return nil
	}
	return heroData
}
//======================================================
//名字：translate_roles(data Herodata)Herodata
//描述：把获取的英雄定位 信息 英文转中文
//====================================================
func translate_roles(data Herodata)Herodata{
   for i:=0;i<len(data.Roles);i++{
	   switch data.Roles[i] {
	   case "support" : data.Roles[i]="辅助";break;
	   case "fighter": data.Roles[i]="战士";break;
	   case "tank": data.Roles[i]="坦克";break;
	   case "mage": data.Roles[i]="法师";break;
	   case "marksman":data.Roles[i]="射手";break
	   case "assassin": data.Roles[i]="刺客";break
	   default:
		   break
	   }
   }
   return data
}
//======================================================
//名字：clear_emply_skins(skin []Skin)[]Skin
//描述: 清除 多余的空白信息的皮肤
//====================================================
func clear_emply_skins(list []Skin)[]Skin{
	lenth:= len(list)
	for i:=0;i<lenth;i++{
		for index,value:= range list {
			if value.Chromas == "1" {
				if index < len(list)-1 {
					list = append(list[:index], list[index+1:]...)
					//fmt.Println(list[index].Name)
					break
				} else {
					list = append(list[:index])
				    break
				}
			}
		}
	}
  return list
}