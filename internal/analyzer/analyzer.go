package analyzer

import (
	"errors"
	"sort"
)

var ErrDataIsEmpty = errors.New("input data is empty, nothing to process")

func CalculateMaxMinMedian(nums []int) (maximumValue, minimumValue int, medianValue float32, err error) {
	sortedNums := make([]int, len(nums))
	sum := 0

	if len(nums) == 0 {
		return 0, 0, 0.0, ErrDataIsEmpty
	}

	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	maximumValue, minimumValue = sortedNums[len(sortedNums)-1], sortedNums[0]

	for _, n := range nums {
		sum += n
	}

	n := len(sortedNums)
	if n%2 == 0 {
		medianValue = float32(sortedNums[n/2]+sortedNums[n/2-1]) / 2
	} else {
		medianValue = float32(sortedNums[n/2])
	}

	return maximumValue, minimumValue, medianValue, nil
}

func CalculateAvg(nums []int) (avgValue float32, err error) {
	sum := 0

	if len(nums) == 0 {
		return 0.0, ErrDataIsEmpty
	}

	for _, n := range nums {
		sum += n
	}

	return float32(sum) / float32(len(nums)), nil
}

func FindSequences(nums []int, incr bool) ([][]int, error) {

	if len(nums) == 0 {
		return nil, ErrDataIsEmpty
	}

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
		return result, nil
	}
	return nil, nil
}
