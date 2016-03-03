package DutchFlag

import (
  "testing"
)

func TestDutchFlagFunc(tt *testing.T) {
  inputs := [][]int{
    {1, 4, 3, 2, 5},
    {1, 1, 1, 1, 1},
    {1},
    {5, 4, 3, 2, 1},
    {1, 2, 3, 4, 5},
    {3, 1, 5, 4, 2},
  }
  index := 0
  for _, input := range inputs {
    pivot := input[index]
    mutableInput := make([]int, len(input))
    copy(mutableInput, input)
    dutchFlag(mutableInput, index)
    if !verify(mutableInput, pivot) {
      tt.Errorf("Input: %v\npivot: %d\nResult: %v", input, pivot, mutableInput)
    }
    tt.Logf("%v returns: %v", input, mutableInput)
  }
}

// verify returns true if every element before the pivot is smaller than the
// pivot and every element after the pivot is greater than the pivot; returns
// false otherwise.
func verify(numbers []int, pivot int) bool {
  bSeenPivot := false
  for i := 0; i < len(numbers); i++ {
    if numbers[i] < pivot && bSeenPivot {
      return false
    }
    if numbers[i] > pivot && !bSeenPivot {
      return false
    }
    if numbers[i] == pivot {
      bSeenPivot = true
    }
  }
  return true
}
