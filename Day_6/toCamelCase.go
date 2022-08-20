package main

import (
	"fmt"
	"unicode"
)

func ToCamelCase(s string) string {
	// your code

	//case 1 if str is empty
	/***
	The-steAlth-warrior ====> TheStealthWarrior
	the-stealth-warrior ===> theStealthWarrior
	the-Stealth-warrior  ===> theStealthWarrior
	the-stEalth-warrior  ===> TheStealthWarrior
	the_stEalth_warrior  ===> TheStealthWarrior

	**/
	returnStr := []rune{}
	a := []rune(s)

	for i := 0; i < len([]rune(s)); i++ {
		if unicode.IsUpper(a[i]) && i == 0 {
			returnStr = append(returnStr, a[i])
		} else if unicode.IsLower(a[i]) && i == 0 {
			returnStr = append(returnStr, (a[i]))
		} else if a[i] == '-' || a[i] == '_' {
			if unicode.IsUpper(a[i+1]) {
				returnStr = append(returnStr, (a[i+1]))
			} else {
				returnStr = append(returnStr, unicode.ToUpper(a[i+1]))
			}
			//fmt.Println(string(returnStr))
			//fmt.Println(len(returnStr))
		} else if returnStr[i-1] == unicode.ToUpper(a[i]) && (a[i-1] == '_' || a[i-1] == '-') {
			//fmt.Println("---")
			//fmt.Println("---", i)
			returnStr = append(returnStr, 0)
		} else {
			//fmt.Println((a[i]))
			if unicode.IsUpper(a[i]) == true {
				returnStr = append(returnStr, unicode.ToLower(a[i]))
			} else {
				returnStr = append(returnStr, a[i])

			}

			//fmt.Println((i))
		}
		//i = i + 2
		//	nlz9hr56
	}

	// if s == "" {
	// 	return ""
	// } else {
	// 	//for camel case
	// 	//var check bool

	// 	for count, v := range s {
	// 		if unicode.IsUpper(v) && count == 0 {
	// 			returnStr = append(returnStr, v)
	// 		} else if unicode.IsLower(v) && count == 0 {
	// 			returnStr = append(returnStr, unicode.ToUpper(v))
	// 		} else if v == '-' || v == '_' {
	// 			returnStr = append(returnStr, v+1)
	// 		} else {
	// 			returnStr = append(returnStr, v)
	// 		}

	// 	}
	// }
	return string(returnStr)
}

func main() {
	fmt.Println(ToCamelCase("The-steAlth-warrior"))
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("the-Stealth-warrior"))
	fmt.Println(ToCamelCase("the_stEalth_warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
	fmt.Println(ToCamelCase("TheStealthWarrior"))
	fmt.Println(ToCamelCase("The Stealth Warrior"))
	fmt.Println(ToCamelCase("The Stealth Warrior"))

	//Street_Blue_down_Black_Yellow

	fmt.Println(ToCamelCase(""))

	//test()
}

func test() {
	//a := []rune("tHE")
	//fmt.Println(a)
	a1 := "the"
	for count, v := range a1 {
		fmt.Println(count)
		if unicode.IsUpper(v) {
			fmt.Println(v)
		} else {
			fmt.Println("no")
		}

	}
	//fmt.Println(unicode.IsUpper())

}
