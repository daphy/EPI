package FirstOccurance

import "testing"

type testcase struct {
  array         []int
  searchNum     int
  expectedIndex int
}

func TestFindFirst(tt *testing.T) {
  testcases := []testcase{
    {[]int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}, 108, 3},
    {[]int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}, 285, 6},
    {[]int{-1}, -1, 0},
    {[]int{0, 0, 0, 0, 0, 0}, 0, 0},
    {[]int{0, 0, 0, 0, 0, 0}, 100, -1},
    {[]int{0, 0, 0, 0, 0, 0}, -100, -1},
    {[]int{0}, 1, -1},
    {[]int{}, 1, -1},
  }
  for _, tc := range testcases {
    if actual := FindFirstOccurance(tc.array, tc.searchNum); actual != tc.expectedIndex {
      tt.Errorf("Finding %d in %v;Expected: %d, actual: %d",
        tc.searchNum, tc.array, tc.expectedIndex, actual)
    }
  }
}
