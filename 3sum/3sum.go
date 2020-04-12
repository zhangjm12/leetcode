package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret [][]int

	length := len(nums)

	sort.Ints(nums)
	for first := 0; first < length-2; first++ {
		firstNum := nums[first]
		if firstNum > 0 {
			break
		}
		if first > 0 && firstNum == nums[first-1] {
			continue
		}

		second, last := first+1, length-1

		for second < last {
			secondNum, lastNum := nums[second], nums[last]
			sum := firstNum + secondNum + lastNum

			if sum == 0 {
				ret = append(ret, []int{firstNum, secondNum, lastNum})
				for second < last && secondNum == nums[second+1] {
					second++
				}
				for second < last && lastNum == nums[last-1] {
					last--
				}
				second++
				last--
			} else if sum < 0 {
				second++
			} else {
				last--
			}
		}
	}
	return ret
}

func main() {
	nums := []int{-4, 3, 0, -1, 5, -1, 0, 1, -3, 2, -5}
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("3sum:%v\n", threeSum(nums))
}
