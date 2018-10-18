package main

import (
	"fmt"
	"random"
)

func main() {
	var j int = 0
	for i := 0; i < 10000000; i++ {
		if random.MillionTrue(10) {
			j++
			if j % 10 == 0 {
				fmt.Println(j)
			}
		}
	}
	fmt.Println(j)
}
