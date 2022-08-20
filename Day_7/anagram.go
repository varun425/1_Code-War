package main

import "fmt"

func Anagrams(word string, words []string) []string {
	index := 0
	res := []string{}
	// aaa := len(words)
	// fmt.Println(aaa)
	var sum1 int = 0
	var sum2 int = 0
	i := 0
	//abc
	// fmt.Println([]rune(words[1]))

	if len(words) == 0 {
		return nil
	} else {

		for {
			mainArr := []rune(words[index])
			if (i >= len(mainArr)) && (len(words) == index) {
				break
			}
			if i >= len(mainArr) { // 0>3
				if sum1 == sum2 {
					res = append(res, words[index])
				}
				index = index + 1
				i = 0
				sum1 = 0
				sum2 = 0
			}
			if len(mainArr) != len([]rune(word)) && (index < len(words)-1) { // 3!=3
				index = index + 1
				mainArr = []rune(words[index])
			} else if len(mainArr) == len([]rune(word)) && (index <= len(words)-1) { //3==3 && 0 < 4
				mainArr = []rune(words[index])
				a1 := []rune(word) // 97,98,99
				//fmt.Println((a1))
				sum1 = sum1 + int(a1[i])      // 97
				sum2 = sum2 + int(mainArr[i]) //
				i = i + 1

			} else {
				// return res
				fmt.Println("-------")
				if len(res) == 0 {
					return nil

				} else {
					return res
				}
			}
		}
	}

	return res
}

func main() {
	a := []string{"lazing", "lazy", "lacer"}
	fmt.Println(Anagrams("carer", a))

}
