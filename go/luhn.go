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
	length := len(cc)
	start := time.Now()

	evalute(cc, length)

	log.Printf("main, execution time %s\n", time.Since(start))
}

func evalute(ccNumString string, length int) bool {

	stringNumArray := []byte(ccNumString)

	var numArray []int

	for i := 0; i < len(ccNumString); i++ {

		num, error := strconv.Atoi(string(stringNumArray[i]))
		_ = error

		//fmt.Println(num)
		numArray = append(numArray, num) // changed my mind, this language is fucking awesome
	}

	fmt.Println(len(numArray))
	fmt.Println(numArray)

	for i := 0; i < len(numArray); i++ {
		if i%2 == 0 {
			evenIndex(numArray[i])
		} else {
			oddIndex(numArray[i])
		}
	}

	if finalSum%10 == 0 {
		fmt.Println("This is a valid credit card number")
		return true

	} else {
		fmt.Println("This some bullshit")
		fmt.Println(finalSum)
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
