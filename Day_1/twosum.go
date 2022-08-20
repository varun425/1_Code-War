package main

func main() {
	//nums := [...]int{1, 2, 4, 6}
	nums := []int{1000, 2, 3, 17, 50}
	//twoSum(nums, 67)
	// println(len(nums))
	println(twoSum(nums, 67))
}

func twoSum(nums []int, target int) []int {
	// var tempArray = []int{}
	// return (append(tempArray, 3, 4))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return ([]int{})

}
