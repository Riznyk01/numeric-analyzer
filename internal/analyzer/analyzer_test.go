package analyzer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestProcessNumericData(t *testing.T) {
	testCases := []struct {
		d              []int
		expectedMin    int
		expectedMax    int
		expectedMedian float32
		expectedAvg    float32
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 5, 3.0, 3},
		{[]int{6, 5, 2, 3, 1, 4}, 1, 6, 3.5, 3.5},
		{[]int{-1, 0, 1}, -1, 1, 0, 0},
	}
	for i, tc := range testCases {
		minimum, maximum, median, avg := ProcessNumericData(tc.d)
		testFailed := fmt.Sprintf("Test case %d failed", i+1)
		assert.Equal(t, tc.expectedMin, minimum, fmt.Sprintf("%s - minimum", testFailed))
		assert.Equal(t, tc.expectedMax, maximum, fmt.Sprintf("%s - maximum", testFailed))
		assert.Equal(t, tc.expectedMedian, median, fmt.Sprintf("%s - median", testFailed))
		assert.Equal(t, tc.expectedAvg, avg, fmt.Sprintf("%s - avg", testFailed))
	}
}

func TestFindSequences(t *testing.T) {
	testCases := []struct {
		d           []int
		seqType     bool
		expectedSeq [][]int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, true, [][]int{
			{1, 2, 3, 4, 5, 6},
		}},
		{[]int{1, 2, 3, 4, 5, 6, 0, 1, 2}, true, [][]int{
			{1, 2, 3, 4, 5, 6},
		}},
		{[]int{1, 2, 3, 4, 5, 6, 0, 1, 2, 1, 2, 3, 4, 5, 6, 7}, true, [][]int{
			{1, 2, 3, 4, 5, 6, 7},
		}},
		{[]int{-100, 1, 2, 3, 4, 5, 6, 0, 1, 2, 1, 2, 3, 4, 5, 6, 7}, true, [][]int{
			{-100, 1, 2, 3, 4, 5, 6}, {1, 2, 3, 4, 5, 6, 7},
		}},
		{[]int{6, 5, 4, 3, 2, 1}, false, [][]int{
			{6, 5, 4, 3, 2, 1},
		}},
		{[]int{6, 5, 4, 3, 2, 1, 3, 2, 1}, false, [][]int{
			{6, 5, 4, 3, 2, 1},
		}},
		{[]int{7, 6, 5, 4, 3, 2, 1, 3, 2, 1, 6, 5, 4, 3, 2, 1}, false, [][]int{
			{7, 6, 5, 4, 3, 2, 1},
		}},
		{[]int{7, 6, 5, 4, 3, 2, 1, 3, 2, 1, 6, 5, 4, 3, 2, 1, -100}, false, [][]int{
			{7, 6, 5, 4, 3, 2, 1}, {6, 5, 4, 3, 2, 1, -100},
		}},
	}
	for i, tc := range testCases {
		result := FindSequences(tc.d, tc.seqType)

		if !reflect.DeepEqual(result, tc.expectedSeq) {
			t.Errorf("Test case %d failed. Expected %v, got %v", i+1, tc.expectedSeq, result)
		}
	}
}
