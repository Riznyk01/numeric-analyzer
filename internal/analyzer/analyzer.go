package analyzer

import (
	"sort"
)

func ProcessNumericData(nums []int) (min, max int, median, avg float32) {
	sortedNums := make([]int, len(nums))
	sum := 0

	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	min = sortedNums[0]
	max = sortedNums[len(sortedNums)-1]

	for _, n := range nums {
		sum += n
	}
	avg = float32(sum) / float32(len(nums))

	n := len(sortedNums)
	if n%2 == 0 {
		median = float32(sortedNums[n/2]+sortedNums[n/2-1]) / 2
	} else {
		median = float32(sortedNums[n/2])
	}

	return min, max, median, avg
}

func FindSequence(nums []int, incr bool) []int {
	result := make([]int, 0)
	tmp := make([]int, 1)
	tmp[0] = nums[0]

	fc := func() {
		if len(tmp) > len(result) {
			result = make([]int, len(tmp))
			copy(result, tmp)
		}
	}

	for i := 1; i < len(nums); i++ {
		if (incr && nums[i] > nums[i-1]) || (!incr && nums[i] < nums[i-1]) {
			tmp = append(tmp, nums[i])
			fc()
		} else {
			fc()
			tmp = []int{nums[i]}
		}
	}

	if len(result) < 2 {
		return nil
	} else {
		return result
	}
}
