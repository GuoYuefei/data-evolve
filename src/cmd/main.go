package main

import (
	"dataTime"
	"fmt"

)

const (
	R0 int64 = 1
	R1 = R0 << iota				//可见iota不是出现一次就加1的，而是每在const的括号里定义一个变量就加一，所以这里的iota=1
	R2
	R3
	R4
	R5 = iota+iota
)

func main() {
	fmt.Println(R1,R2,R3,R4,R5)
	//var j int = 0
	//var start,end time.Time
	//start = time.Now()
	//for i := 0; i < 10000; i++ {
	//	if random.PercentTrue(0.1) {
	//		j++
	//		if j % 10 == 0 {
	//			fmt.Println(j)
	//		}
	//	}
	//}
	//end = time.Now()
	//fmt.Println("ns : ", end.Sub(start))
	//fmt.Println(j)
	//fmt.Println(end.Unix(),end.UnixNano(),end.Nanosecond(),end.String(),end)

	t := dataTime.NewTime(12*dataTime.MY + 2018*dataTime.Year + 10 * dataTime.Month + 22 * dataTime.Day +
		16 * dataTime.Hour + 7 * dataTime.Minute + 41 * dataTime.Second + 2 * dataTime.TenSecond)
	fmt.Println(t.TenSec)
	fmt.Println(t.Second())
	fmt.Println(t)
	fmt.Println(t.Year())
}

//31104000000 1y
