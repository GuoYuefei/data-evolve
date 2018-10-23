package main

import (
	"dataSystem"
	"fmt"
	"time"
)

/**
用于检测SystemClock能否正常使用
 */

 var SystemTimeNow *dataSystem.SystemClock

func main() {
	SystemTimeNow = dataSystem.NewSystemClock()
	SystemTimeNow.Relative = dataSystem.R50
	SystemTimeNow.Walk()

	go func() {
		for i := 0; i < 130; i++ {

			time.Sleep(time.Second * 10)
			fmt.Println("10 \t",i,":",SystemTimeNow.Now())
		}
	}()
	go func() {
		for i := 0; i < 150; i++ {
			time.Sleep(time.Second * 3)
			fmt.Println("3 \t",i,":",SystemTimeNow.Now())
		}
	}()
	go func() {

		for i := 0; i < 10000000; i++ {
			time.Sleep(time.Second * 17)
			fmt.Println("17 \t",i,":",SystemTimeNow.Now())
		}
	}()

	time.Sleep(10*time.Hour)
	//fmt.Println(SystemTimeNow.Now())
	//fmt.Println(SystemTimeNow.Now().Second(), SystemTimeNow.Now().TenSec, SystemTimeNow.Now().Minute())

}
