package main

import (
	"dataSystem"
	"fmt"
	"time"
)

/**
用于检测SystemClock能否正常使用
 */


func main() {
	t := dataSystem.NewSystemClock()
	t.Relative = dataSystem.R1
	t.Walk()
	time.Sleep(10 * time.Second)
	fmt.Println(t.Now())
}
