package main

import (
	"fmt"
	"player"
)

func main(){
   fmt.Printf("\n作业一\n")
    ro:=rot{
		r:    10,
		s:    0,
		Long: 0,
	}
	ro.s = ro.S()
	ro.Long = ro.L()
	fmt.Printf("%v",ro)
	//==========================================
	fmt.Printf("\n作业二\n")
    sp := player.Sporter{
		Name:  "王二狗",
		Sex:   "未知",
		Sport: "男子蛙泳100米",
		Grade: 100,
	}
	sp.Run()
    //=============================================
   fmt.Printf("\n作业三\n")
   car := player.Car(player.ElectroCar{
	   Bran:"特斯拉",
	   Kin: "电动汽车",
   })
   car.Brand()
   car.Kind()
   car = player.Tricycle{
	   Bran: "时风",
	   Kin:  "三轮车",
   }
	car.Brand()
	car.Kind()
	car = player.Truck{
		Bran: "解放牌",
		Kin:  "卡车",
	}
	car.Brand()
	car.Kind()
}

type rot struct {
	r float64
	s float64
	Long float64
}

func (a rot)S() float64 {
	r := a.r
	s:= r*r*3.14
	return s
}

func (a rot)L() float64 {
	r:=a.r
	long:= 2*r*3.14
	return long
}
//=======================================
