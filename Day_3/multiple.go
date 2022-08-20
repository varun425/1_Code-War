package main

import "fmt"

func main() {
	fmt.Println(Multiple(1000))
}

func Multiple(n int) int {
	var sum int = 0
	for i := 0; i < n; i++ {
		if i%3 == 0 {
			sum = sum + i
		} else if i%5 == 0 {
			sum = sum + i
		} else {

		}
	}

	return sum
}
