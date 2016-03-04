package MaxDifference

import "testing"

type testCase struct {
  input    []int64
  expected int64
}

func TestFindMaxDifference(tt *testing.T) {
  testCases := []testCase{
    testCase{[]int64{1, 2, 4, 3, 9, 0, 2}, 8},
    testCase{[]int64{2, 0}, 0},
    testCase{[]int64{1}, 0},
    testCase{[]int64{}, 0},
    testCase{[]int64{1, 2, 3}, 2},
    testCase{[]int64{1, 1, 1, 1, 1, 1, 1, 1}, 0},
  }
  for _, test := range testCases {
    if actual := findMaxDifference(test.input); actual != test.expected {
      tt.Errorf("Max difference of %v should be %d. Actual: %d", test.input,
        test.expected, actual)
    }
  }

}
