package main

import (
	"fmt"
)

func rub(nums []int) int {
	n := len(nums)
	r := make([]int, n)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	r[0] = nums[0]
	r[1] = Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		r[i] = Max(nums[i]+r[i-2], r[i-1])
	}
	return r[n-1]
}

func rub1(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return Max(nums[0], nums[1])
	}

	return Max(RobRange(nums, 0, n-2), RobRange(nums, 1, n-1))
}

func RobRange(nums []int, start int, end int) int {
	r := make([]int, end-start+1)

	r[0] = nums[start]
	r[1] = Max(nums[start], nums[start+1])

	for i := 2; i < end-start+1; i++ {
		r[i] = Max(nums[start+i]+r[i-2], r[i-1])
	}
	return r[end-start]
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	test := []int{3, 1, 2, 5, 3, 1, 4, 11, 5, 5}
	fmt.Println(rub(test))
	fmt.Println(rub1(test))
}
