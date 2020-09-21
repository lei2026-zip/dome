package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Printf("\n计算器-2020\n")
	var rec string
	var ind bool
	var index_H,index_L int =0,0
	fmt.Scanf("%s",&rec)
	a:
		ind = strings.Contains(rec,"(")
		fmt.Printf("%t",ind)
		if ind {
			index_L =strings.Index(rec,"(")
			index_H = strings.Index(rec,")")+1
			fmt.Printf("%d %d",index_H,index_L)
			str := count(rec[index_L:index_H])
			rec = strings.Replace(rec,rec[index_L:index_H],str,1)
			fmt.Printf("%s",rec)
			ind = strings.Contains(rec,"(")
			if ind{
				goto a
			}
		}
}

func count(H string)string{
  return  "qwew"
}
