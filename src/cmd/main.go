package main

import (
	"dataTime"
	"fmt"
	"random"
	"time"
)

func main() {
	var j int = 0
	var start,end time.Time
	start = time.Now()
	for i := 0; i < 10000; i++ {
		if random.PercentTrue(0.1) {
			j++
			if j % 10 == 0 {
				fmt.Println(j)
			}
		}
	}
	end = time.Now()
	fmt.Println("ns : ", end.Sub(start))
	fmt.Println(j)
	fmt.Println(end.Unix(),end.UnixNano(),end.Nanosecond(),end.String(),end)

	t := dataTime.NewTime(12*dataTime.MY + 2018*dataTime.Year + 10 * dataTime.Month + 22 * dataTime.Day +
		16 * dataTime.Hour + 7 * dataTime.Minute + 27 * dataTime.Second + 596 * dataTime.Millisecond )
	fmt.Println(t.Millisecond)
	fmt.Println(t.Sec())
	fmt.Println(t.String())
	fmt.Println(t)
}

//31104000000 1y
