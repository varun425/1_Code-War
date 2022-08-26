package main

import "fmt"

func main() {
	fmt.Println(appendTolast())
	optimizeWayToshiftZeros()
}

func appendTolast() []int {
	arr := []int{1, 0, 2, 0, 3, 0, 4, 5, 6}
	arr2 := []int{}
	arr3 := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {

			// count := count + 1
			arr3 = append(arr3, 0)
		} else {
			// arr = arr[i:]
			arr2 = append(arr2, arr[i])
		}
	}
	arr2 = append(arr2, arr3...)
	return arr2
	// mt.Println(arr)
}

func optimizeWayToshiftZeros() {

	arr := []int{1, 0, 2, 0, 3, 0, 4, 5, 6}
	j := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	fmt.Println(arr)
}
