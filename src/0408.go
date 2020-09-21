package main

import "fmt"

type company_staff struct{
	name string
	cord int
	area string
	salary int
	tallage float32
}

func main(){
	fmt.Printf("\n作业一\n")
	sb := new(company_staff)
	sb.name = "王二狗"
	sb.area = "未知"
	sb.salary = 5001
	fmt.Printf("%#v",sb)
	fmt.Printf("\n作业二\n")
    *sb = selec(*sb)
	fmt.Printf("应收税%f",sb.tallage)
}

func selec(a company_staff) company_staff{
	b:= a.salary
	if a.salary>5000 {
		a.tallage = float32(b-5000)*0.2
	}else{
		a.tallage = 0
	}
	return a
}