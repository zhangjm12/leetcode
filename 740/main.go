package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) (ans int) {
	sort.Ints(nums)
	sum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if val := nums[i]; val == nums[i-1] {
			sum[len(sum)-1] += val
		} else if val == nums[i-1]+1 {
			sum = append(sum, val)
		} else {
			ans += rob(sum)
			sum = []int{val}
		}
	}
	ans += rob(sum)
	return
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 4, 2}
	fmt.Println(deleteAndEarn((nums)))
	nums = []int{2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn((nums)))
}
