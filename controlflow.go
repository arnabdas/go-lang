package main

import (
	"fmt"
	"strconv"
)

func checkFor() {
	i := 1
	for i < 5 {
		status := " "
		if i%2 == 0 {
			status += "even"
		} else {
			status += "odd"
		}
		fmt.Println(strconv.Itoa(i) + status)
		i++
	}
}
