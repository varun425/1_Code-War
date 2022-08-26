package main

import "fmt"

func main() {

	fmt.Println(substring())
	fmt.Println(substringFrequency())

}

func substring() bool {
	var str string = "varurrn"
	var str2 string = "rr"

	count := 0
	count2 := 0
	for {

		if count2 >= len(str2) {

			return true

		}
		if count >= len(str) {
			return false
			//break
		}

		if rune(str2[count2]) == rune(str[count]) {
			count2 = count2 + 1 //  now at 1
			count = count + 1

		} else {
			count = count + 1
		}

	}

}

func substringFrequency() int {
	var str string = "varunrururru"
	var str2 string = "r"

	count := 0
	count2 := 0
	f := 0
	for {

		if count2 == len(str2) {
			count2 = 0
			f = f + 1

		}
		if count >= len(str) {
			return f

		}

		if rune(str2[count2]) == rune(str[count]) {
			count2 = count2 + 1 //  now at 1
			count = count + 1

		} else if count2 < len(str2) {
			count = count + 1
		}

	}

}
