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

// In division, the dividend is divided by the divisor to get a quotient and there may be some remainder
func numberToWords(input int) string {
	rem := input % 10
	qtn := input / 10

	words := ""
	if qtn >= 10 {
		words += numberToWords(qtn)
	} else {
		words += digitToWord(qtn)
	}

	return words + digitToWord(rem)
}

func digitToWord(digit int) string {
	words := ""
	switch digit {
	case 0:
		words += "zero "
	case 1:
		words += "one "
	case 2:
		words += "two "
	case 3:
		words += "three "
	case 4:
		words += "four "
	case 5:
		words += "five "
	case 6:
		words += "six "
	case 7:
		words += "seven "
	case 8:
		words += "eight "
	case 9:
		words += "nine "
	default:
		words += "unknown "
	}

	return words
}

func fizzBizz() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBizz")
		} else if i%5 == 0 {
			fmt.Println("Bizz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
