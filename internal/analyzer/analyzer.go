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

func FindSequences(nums []int, incr bool) [][]int {
	result := [][]int{{}}
	currentSequence := make([]int, 0)
	appendNext, currentSequenceProcessed := false, false

	// Check length of currentSequence.
	// If greater than any sequences in the result, set result to nil and append currentSequence.
	// If a sequence of this length already exists in the result, add current sequence.
	fc := func() {
		if len(currentSequence) > len(result[0]) {
			result = nil
			result = append(result, currentSequence)
		} else if len(currentSequence) == len(result[0]) && len(currentSequence) > 0 {
			result = append(result, currentSequence)
		}
		currentSequenceProcessed = true
		currentSequence = nil
	}

	for i := 0; i < len(nums)-1; i++ {
		if (incr && nums[i] < nums[i+1]) || (!incr && nums[i] > nums[i+1]) {
			currentSequence = append(currentSequence, nums[i])
			appendNext = true
			currentSequenceProcessed = false
		} else {
			if appendNext {
				currentSequence = append(currentSequence, nums[i])
				appendNext = false
			}
			fc()
		}
		if i == len(nums)-2 && !currentSequenceProcessed {
			if appendNext {
				currentSequence = append(currentSequence, nums[i+1])
				appendNext = false
			}
			fc()
		}
	}
	if len(result) > 0 && len(result[0]) > 0 {
		return result
	}
	return nil
}
