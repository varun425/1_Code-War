package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution("abc"))
	fmt.Println(Solution1("abc"))

}

func Solution(str string) []string {

	s := []string{}

	if (len(str) % 2) != 0 {
		str = fmt.Sprintf("%v%v", str, "_")
		fmt.Println(str)

	}
	a := strings.Split(str, "")
	for i := 0; i < len(a); i++ {
		t := fmt.Sprintf("%v%v", a[i], a[i+1])
		s = append(s, t)
		i++
	}
	return s
}

//-------------------------codewar best practice solutions ----------------------

func Solution1(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 { //abc_
		res = append(res, str[i:i+2])
		// [ab] i = 0
		// [c_] i = 2
	}
	return res
}
