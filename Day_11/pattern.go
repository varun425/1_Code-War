package main

import (
	"fmt"
)

func main() {

	pattern()
	patternTraingle()

}

func pattern() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v", "")
			fmt.Printf("%v", j)
		}
		fmt.Println()
	}
}

/**
        *
      * * *
    * * * * *
  * * * * * * *
* * * * * * * * *

*/

func patternTraingle() {
	var k int

	for i := 1; i <= 9; i++ {
		k = 0
		for s := 1; s <= 9-i; s++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")

	}
}
