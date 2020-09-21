package main

import "fmt"

func main() {
	var name string
	var password, weight, i int
	var height float32
	fmt.Printf("请输入用户名:")
	fmt.Scanf("%s", &name)
	fmt.Printf("请输入密码:")
	fmt.Scanf("%d", &password)
	fmt.Printf("你的身高是/m：")
	fmt.Scanf("%f", &height)
	fmt.Printf("你的体重是/kg：")
	fmt.Scanf("%d", &weight)
	fmt.Printf("请回车确认你的信息\n")
	fmt.Printf("用户名%s 密码%d 身高%.3fm 体重%dkg",name,password,height,weight)
	fmt.Printf("\n输入指定数字:")
	fmt.Scanln(&i)
	if i == 1 {
		fmt.Printf("\n星期一")
	} else {
		if i == 2 {
			fmt.Printf("\n星期二")
		} else {
			if i == 3 {
				fmt.Printf("\n星期三")
			} else {
					if i == 4 {
						fmt.Printf("\n星期四")
					} else {
						if i == 5 {
							fmt.Printf("\n星期五")
						} else {
							if i == 6 {
								fmt.Printf("\n星期六")
							} else {
								if  i == 7 {
									fmt.Printf("\n星期日")
								} else {
									fmt.Printf("\n错误信息")
								}
							}
						}
					}
				}

		}
	}
}