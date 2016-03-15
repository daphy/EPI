package StackMax

import "testing"

type testcase struct {
  pushSeq   []int32
  numOfPops int
  max       int32
}

var testcases = []testcase{
  {[]int32{1}, 0, 1},
  {[]int32{1, 6, 2, 3, 5, 10, 2}, 1, 10},
  {[]int32{1, 6, 2, 3, 5, 10, 2}, 2, 6},
}

func TestGetMax(tt *testing.T) {
  // getStack is a function which will return a Stack-type (interface) object.
  var executeTest = func(getStack func() Stack) {
    for _, tc := range testcases {
      stack := getStack()
      for _, pushingNum := range tc.pushSeq {
        stack.Push(pushingNum)
      }
      for ii := 0; ii < tc.numOfPops; ii += 1 {
        stack.Pop()
      }
      actualMax, err := stack.Max()
      if err != nil {
        tt.Errorf("Pushing %v to stack and pop %d times should have "+
          "max of %d; got error %s instead.", tc.pushSeq, tc.numOfPops,
          tc.max, err)
        continue
      }
      if actualMax != tc.max {
        tt.Errorf("Pushing %v to stack and pop %d times should have "+
          "max of %d; got %d instead.", tc.pushSeq, tc.numOfPops,
          tc.max, actualMax)
        continue
      }
    }
  } // End of executeTest definition.
  executeTest(GetStack1)
  executeTest(GetStack2)
}
