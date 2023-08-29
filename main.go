package main

import (
	"fmt"
	"strings"

	"math/rand"
	"time"
)

func main() {
	//task1
	printSegitiga(6)

	//task2
	genPass("abcdef", "Med")

	//task 3
	durationMovie([]int{2, 7, 6, 4, 8, 9}, 9)

}

// Task 1
func printSegitiga(rows int) {
	var left, right int
	for left = rows; left >= 1; left-- {
		for barisBintang := 1; barisBintang <= rows-left; barisBintang++ {
			fmt.Print("  ")
		}
		for right = left; right <= 2*left-1; right++ {
			fmt.Printf("* ")
		}
		for right = 0; right < left-1; right++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}
}

// task 2
var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	mediumPass     = upperCharSet + lowerCharSet + numberSet
	allCharSet     = lowerCharSet + upperCharSet + (specialCharSet) + numberSet + specialCharSet
)

func genPass(pass string, level string) {
	lengthPass := len(pass)
	rand.Seed(time.Now().UnixNano())

	numRand := numberSet[rand.Intn(len(numberSet))]
	numStr := string(numRand)

	if level == "Low" && lengthPass >= 6 {
		passLow := pass + numStr
		fmt.Print(passLow)
		fmt.Println(numRand)
	}

	if level == "Med" && lengthPass >= 6 {

		chars := []rune(mediumPass)
		var b strings.Builder
		for i := 0; i < lengthPass; i++ {
			b.WriteRune(chars[rand.Intn((len(chars)))])
		}
		str := b.String()
		passMed := pass + str
		fmt.Println(passMed)
	}

	if level == "Strong" && lengthPass >= 6 {

		chars := []rune(allCharSet)
		var b strings.Builder
		for i := 0; i < lengthPass; i++ {
			b.WriteRune(chars[rand.Intn((len(chars)))])
		}
		str := b.String()
		passStrong := pass + str
		fmt.Println(passStrong)

	} else if level == "" || lengthPass < 6 {
		fmt.Println("panjang password harus lebih dari 6 atau masukkan level")
	}
}

// task 3
func durationMovie(arr []int, n int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == n {
				fmt.Print(arr[i], " dan ", arr[j])
			}
		}
	}
	return
}
