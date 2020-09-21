package main

import "fmt"

func main(){
	//var a int =0
	//rand.Seed(time.Now().UnixNano())
	// for i:=0;i<10;i++{
	// 	a = rand.Intn(10)
    //     println(a)
	// }
	fmt.Printf("作业一:\n")
	a := []int{1,2,3,4,5}
	for _,value := range a{
		fmt.Printf("%d",value)
	}
	fmt.Printf("作业二:\n")
	str := []string{"哈哈","嘿嘿","嘤嘤"}
	for _,value := range str{
		fmt.Printf("%s",value)
	}
	fmt.Printf("\n作业三:\n")
	flo := []float32{1.11,2.22,3.33,4.44}
	 var F float32=0
	for _,value := range flo{
		F +=value
	}
	fmt.Printf("%0.2f 总和%.2f",flo,F)
	fmt.Printf("\n作业四:\n")
	u8 := [10]int{}
	for i:=0;i<10;i++{
		fmt.Scanf("%d",&u8[i])
	}
	fmt.Printf("%d",u8)
}


