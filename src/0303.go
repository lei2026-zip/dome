package main

import "fmt"

func main(){
	a := 0
	println("作业一请输入1—7的数")
	fmt.Scanf("%d",&a)
	switch a {
	case 1:  fmt.Printf("\n星期一")
	case 2:  fmt.Printf("\n星期二")
	case 3:  fmt.Printf("\n星期三")
	case 4:  fmt.Printf("\n星期四")
	case 5:  fmt.Printf("\n星期五")
	case 6:  fmt.Printf("\n星期六")
	case 7:  fmt.Printf("\n星期天")
	default:
		println("您输入了错误的信息")
	}
	print("\n")
	month := 0
	println("作业二请输入某一月份1-12：")
	fmt.Scanf("%d",&month)
	switch month {
	case 1:  fmt.Printf("\n冬季")
	case 2:  fmt.Printf("\n冬季")
	case 3:  fmt.Printf("\n春季")
	case 4:  fmt.Printf("\n春季")
	case 5:  fmt.Printf("\n春季")
	case 6:  fmt.Printf("\n夏季")
	case 7:  fmt.Printf("\n夏季")
	case 8:  fmt.Printf("\n夏季")
	case 9:  fmt.Printf("\n秋季")
	case 10:  fmt.Printf("\n秋季")
	case 11:  fmt.Printf("\n秋季")
	case 12:  fmt.Printf("\n冬季")
	default:
		println("您输入了错误的信息")
	}
	print("\n")
	num1,num2,str := 0,0,0
	println("作业三请输入两个数中间以运输符代数衔接衔接")
	print("\n 0|1++2--  1|+ 2|- 3|* 4|/ 5% 例如输入4 2 1回车结果为4-1=3\n")
	fmt.Scanf("%d %d %d",&num1,&str,&num2)
	switch str {
	case 0: num2--;num1++;  fmt.Printf("\nnum1++%dnum2--%d",num1,num2)
	case 1:  fmt.Printf("\n %d+%d=%d",num1,num2,num1+num2)
	case 2:  fmt.Printf("\n %d-%d=%d",num1,num2,num1-num2)
	case 3:  fmt.Printf("\n %d*%d=%d",num1,num2,num1*num2)
	case 4:  fmt.Printf("\n %d/%d=%d",num1,num2,num1/num2)
	case 5:  fmt.Printf("\n %d余%d=%d",num1,num2,num1%num2)
	default:
		println("您输入了错误的信息")
	}

}
