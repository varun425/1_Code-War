package main

import "fmt"

func main() {
	fmt.Println(BouncingBall(3, 0.1, 1.5))
}

func BouncingBall(h, bounce, window float64) int {
	// your code

	var Counter int = 1
	if h == window {
		return -1
	}
	if h == 0 || bounce >= 1 || bounce <= 0 || window >= h {
		return -1
	} else {
		//	l := bounce * h // 16
		//	var l float64
		for {
			h = bounce * h
			if h < window {
				//fmt.Println(h)
				break
			} else {
				Counter = Counter + 2
			}

			//fmt.Println(h)

		}

	}
	return Counter
}
