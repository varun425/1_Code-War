package main

import (
	"fmt"
)

func main() {
	fmt.Println(("varun"))
	frequencyOfeachChar()

}

func Charfrequency(a string, s string) int {
	var count int
	for _, value := range s {
		//fmt.Println(key, value)
		if a == string(value) {
			count = count + 1
		}
	}
	return count

}

func frequencyOfeachChar() {
	var char string = "vvaarrruuuunnnnnk"
	// var mapping map[string]int
	m := make(map[string]int)
	for _, val := range char {
		m[string(val)] = m[string(val)] + 1
	}
	fmt.Println(m)
}
