package analyzer

import (
	"sort"
)

func Min(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func Max(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func Avg(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}
func Median(nums []int) int {
	n := len(nums) / 2
	if len(nums)%2 == 0 {
		return (nums[n] + nums[n-1]) / 2
	}
	return nums[n]
}
