package main

import "LeetcodeProject/collection/algorithm"

func main() {
	var nums = []int{3, 2, 4}
	target := 6
	nums = algorithm.SumOfTwoNum(nums, target)
	for _,v := range nums {
		println("main nums:", v)
	}
}
