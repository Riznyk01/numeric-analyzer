package analyzer

import (
	"errors"
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
		expectedError  error
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 5, 3.0, 3, nil},
		{[]int{6, 5, 2, 3, 1, 4}, 1, 6, 3.5, 3.5, nil},
		{[]int{-1, 0, 1}, -1, 1, 0.0, 0.0, nil},
		{[]int{}, 0, 0, 0.0, 0.0, ErrDataIsEmpty},
	}
	for i, tc := range testCases {
		minimum, maximum, median, avg, err := ProcessNumericData(tc.d)

		if !errors.Is(err, tc.expectedError) {
			t.Errorf("Test case %d failed. Expected error %v, got %v", i+1, tc.expectedError, err)
		}

		testFailed := fmt.Sprintf("Test case %d failed", i+1)
		assert.Equal(t, tc.expectedMin, minimum, fmt.Sprintf("%s - minimum", testFailed))
		assert.Equal(t, tc.expectedMax, maximum, fmt.Sprintf("%s - maximum", testFailed))
		assert.Equal(t, tc.expectedMedian, median, fmt.Sprintf("%s - median", testFailed))
		assert.Equal(t, tc.expectedAvg, avg, fmt.Sprintf("%s - avg", testFailed))
	}
}

func TestFindSequences(t *testing.T) {
	testCases := []struct {
		d             []int
		seqType       bool
		expectedSeq   [][]int
		expectedError error
	}{
		{[]int{1, 2, 3, 4, 5, 6}, true, [][]int{
			{1, 2, 3, 4, 5, 6}}, nil},
		{[]int{1, 2, 3, 4, 5, 6, 0, 1, 2}, true, [][]int{
			{1, 2, 3, 4, 5, 6}}, nil},
		{[]int{1, 2, 3, 4, 5, 6, 0, 1, 2, 1, 2, 3, 4, 5, 6, 7}, true, [][]int{
			{1, 2, 3, 4, 5, 6, 7}}, nil},
		{[]int{-100, 1, 2, 3, 4, 5, 6, 0, 1, 2, 1, 2, 3, 4, 5, 6, 7}, true, [][]int{
			{-100, 1, 2, 3, 4, 5, 6}, {1, 2, 3, 4, 5, 6, 7}}, nil},
		{[]int{6, 5, 4, 3, 2, 1}, false, [][]int{
			{6, 5, 4, 3, 2, 1}}, nil},
		{[]int{6, 5, 4, 3, 2, 1, 3, 2, 1}, false, [][]int{
			{6, 5, 4, 3, 2, 1}}, nil},
		{[]int{7, 6, 5, 4, 3, 2, 1, 3, 2, 1, 6, 5, 4, 3, 2, 1}, false, [][]int{
			{7, 6, 5, 4, 3, 2, 1}}, nil},
		{[]int{7, 6, 5, 4, 3, 2, 1, 3, 2, 1, 6, 5, 4, 3, 2, 1, -100}, false, [][]int{
			{7, 6, 5, 4, 3, 2, 1}, {6, 5, 4, 3, 2, 1, -100}}, nil},
		{[]int{}, false, nil, ErrDataIsEmpty},
		{[]int{}, true, nil, ErrDataIsEmpty},
	}
	for i, tc := range testCases {
		result, err := FindSequences(tc.d, tc.seqType)

		if !errors.Is(err, tc.expectedError) {
			t.Errorf("Test case %d failed. Expected error %v, got %v", i+1, tc.expectedError, err)
		}

		if !reflect.DeepEqual(result, tc.expectedSeq) {
			t.Errorf("Test case %d failed. Expected %v, got %v", i+1, tc.expectedSeq, result)
		}
	}
}
