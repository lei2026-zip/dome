package main

import "fmt"

func main(){
	fmt.Printf("\n作业一\n")
	b := []int{1,2,3,4,5}
	fmt.Printf("%d",b)
	b = chang(b)
	fmt.Printf("%d",b)
	fmt.Printf("\n作业二\n")
	c := []int{23,232,73,1293,321}
	fmt.Printf("%d",c)
	d := []*int{&c[0],&c[1],&c[2],&c[3],&c[4]}
	e :=chang2(d)
	fmt.Printf("%d",*e)
}

func chang(a []int) []int{
	l := len(a)
	for i:=0;i<l/2;i++{
        a[i],a[l-i-1] = a[l-i-1],a[i]
	}
	return a
}

func chang2(a []*int) *[]int{
	var l int = len(a)
	arr := make([]int,l,l)
   for i:=0;i<l;i++{
   	   for j:=i;j<l;j++ {
		   if *a[j] < *a[i] {
			   *a[i],*a[j] =*a[j],*a[i]
		   }
	   }
	   arr[i] = *a[i]
   }
   return &arr
}


