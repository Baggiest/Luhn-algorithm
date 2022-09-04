package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

var finalSum int

func main() {

	cc := "4539479425649801"
	_ = cc
	ccArray := []string{"4539479422649801", "4539479422649801", "4539479425649801", "4539479425649801", "4539479425649801", "4539479425649801","4539479425649801","45394794d25649801","4539479425649801", "4539479425649801"}
	start := time.Now()

	for i := 0; i < len(ccArray); i++ {
		evalute(ccArray[i])
	}

	log.Printf("Time to evaluate 10 %s\n", time.Since(start))
}

func evalute(ccNumString string) bool {

	stringNumArray := []byte(ccNumString)

	var numArray []int

	for i := 0; i < len(ccNumString); i++ {

		num, error := strconv.Atoi(string(stringNumArray[i]))

		if error != nil {
			fmt.Println("Its a credit card number not a credit card string my brother ⁉️")
			return false
		}

		//fmt.Println(num)
		numArray = append(numArray, num) // changed my mind, this language is fucking awesome
	}

	// fmt.Println(len(numArray))
	// fmt.Println(numArray)

	for i := 0; i < len(numArray); i++ {
		if i%2 == 0 {
			evenIndex(numArray[i])
		} else {
			oddIndex(numArray[i])
		}
	}

	if finalSum%10 == 0 {
		fmt.Println("This is a valid credit card number ✅")
		finalSum = 0
		return true

	} else {
		fmt.Println("This is not a valid credit card number ❌")
		finalSum = 0
		return false
	}
}

func evenIndex(num int) {
	bothDigits := num * 2
	firstDigit := bothDigits % 10

	temp := float64(bothDigits)
	secondDigit := int(math.Floor(temp / 10))

	finalSum += secondDigit + firstDigit
}

func oddIndex(num int) {
	finalSum += num
}
