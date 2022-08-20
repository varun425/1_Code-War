package main

import "fmt"

func main() {

	fmt.Println(ValidBraces("(({{[[]]}}))")) // paased
	//fmt.Println(test())

}

/**
"(){}[]"   =>  True	[40 41 123 125 91 93]
"([{}])"   =>  True	[40 91 123 125 93 41]
"(}"       =>  False
"[(])"     =>  False	[91 40 93 41]
"[({})](]" =>  False  [91 40 123 125 41 93 40 93]
"([]){}" => true

*/

func ValidBraces(str string) bool {
	//strOfBraces := "(({{[[]]}}))"
	a := []rune(str) // 40 41
	fmt.Println(a)
	var Stack = []int{}
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case 40:
			Stack = append(Stack, int(a[i]))
		case 91:
			//fmt.Println("aaa", a[i])
			Stack = append(Stack, int(a[i])) //pass
		case 123:
			//fmt.Println("aaa", a[i])
			Stack = append(Stack, int(a[i]))
		case 125:
			Stack = append(Stack, int(a[i]))
			temp := len(Stack)
			// if temp == 1 && a[i] == 93 && a[i] == 41 && a[i] == 125 {
			// 	return false
			// }
			if temp < 2 {
				return false
			}
			if Stack[temp-1] == Stack[temp-2]+2 {
				Stack = Stack[:temp-2]
			} else {
				return false
			}
		case 41:
			Stack = append(Stack, int(a[i]))
			temp := len(Stack)
			if temp < 2 {
				return false
			}
			if Stack[temp-2]+1 == Stack[temp-1] {
				fmt.Println(temp)
				Stack = Stack[:temp-2]
			} else {
				return false
			}
		case 93:
			Stack = append(Stack, int(a[i]))
			temp := len(Stack)
			if temp < 2 {
				return false
			}
			if Stack[temp-1] == Stack[temp-2]+2 {
				Stack = Stack[:temp-2]
			} else {
				return false
			}
		}

	}
	if len(Stack) == 0 {
		//fmt.Println(len(Stack))
		return true
	}
	return false

}

// func test() []int {
// 	a := []int{}
// 	for i := 0; i < 2; i++ {
// 		a = append(a, i)
// 	}
// 	return a[:0]
// }
