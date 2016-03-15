package StackMax

import "fmt"

type Stack2 struct {
  elements      []int32
  size          int32
  maxValueSoFar [][2]int32 // []int32{max value, num of times count}
  maxValueSize  int32
}

func (stack *Stack2) Push(num int32) {
  stack.elements = append(stack.elements, num)
  if stack.size == 0 || stack.maxValueSize == 0 ||
    num > stack.maxValueSoFar[stack.maxValueSize-1][0] {
    stack.maxValueSoFar = append(stack.maxValueSoFar, [2]int32{num, 1})
    stack.maxValueSize += 1
  } else {
    stack.maxValueSoFar[stack.maxValueSize-1][1] += 1
  }
  stack.size += 1
}

func (stack *Stack2) Pop() (int32, error) {
  if stack.size == 0 {
    return 0, fmt.Errorf("Stack is empty.")
  }
  stack.size -= 1
  popingNum := stack.elements[stack.size]
  if stack.size == 0 {
    stack.elements = []int32{}
    stack.maxValueSoFar = [][2]int32{}
  } else {
    stack.elements = stack.elements[:stack.size]
    // Decrease the last element's count.
    stack.maxValueSoFar[stack.maxValueSize-1][1] -= 1
    if stack.maxValueSoFar[stack.maxValueSize-1][1] <= 0 {
      stack.maxValueSize -= 1
      stack.maxValueSoFar = stack.maxValueSoFar[:stack.maxValueSize]
    }
  }
  return popingNum, nil
}

func (stack *Stack2) Max() (int32, error) {
  if stack.size == 0 || stack.maxValueSize == 0 {
    return 0, fmt.Errorf("Stack is empty.")
  }
  return stack.maxValueSoFar[stack.maxValueSize-1][0], nil
}
