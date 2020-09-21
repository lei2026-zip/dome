package main

import "fmt"

func main(){
	fmt.Printf("\n作业一\n")
	fmt.Printf("矩形长%d 宽%d 周长%d",24,12,long(24,12))
	fmt.Printf("\n作业二\n")
    st := "iwfnjdnisdnoqwidjqklwojfjoqwijdw"
	switc(st)
	fmt.Printf("\n作业三\n")
	fmt.Printf("\n半径为%0.0f的圆面积为%0.2f",4.0,rotar(4.0))
}

func long(a int,b int)int{
	return a*2+b*2
}

func switc(str string){
	fmt.Printf("%s\n",str)
	arr :=make(map[string]int)
	var a string = ""
	for i:=0;i<len(str);i++{
		a = str[i:i+1]
		if arr[a]!=0 {
			continue
		}
		for j:=0;j<len(str);j++{
			if a == str[j:j+1]{
				arr[a]++
			}
		}
	}
	fmt.Println(arr)
}

func rotar(r float32) float32{
    rot := r*r*3.14
    return rot
}