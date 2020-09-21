package into

import "fmt"
type Mouse interface {
	name() string
}

type Staff struct {
	Name string
	Cert Technical
	Age int8
    Birth Birthday
}

type Technical struct {
	Driving bool
	Degree string
}

type Birthday struct {
	Year int
	Mouth int8
	Day int8
}

func (a Staff)ShowInformation(){
   fmt.Printf("名字%s 年龄%d \n出生年月日%d年%d月%d日\n 是否有机动车驾驶证:%t \n 学位:%s",a.Name,a.Age,a.Birth.Year,a.Birth.Mouth,a.Birth.Day,a.Cert.Driving,a.Cert.Degree)
}