package StackMax

import "fmt"

type Stack1 struct {
  elements      []int32
  maxValueSoFar []int32
  size          int32
}

func (stack *Stack1) Push(num int32) {
  stack.elements = append(stack.elements, num)
  if stack.size == 0 || num > stack.maxValueSoFar[stack.size-1] {
    stack.maxValueSoFar = append(stack.maxValueSoFar, num)
  } else {
    stack.maxValueSoFar = append(stack.maxValueSoFar,
      stack.maxValueSoFar[stack.size-1])
  }
  stack.size += 1
}

func (stack *Stack1) Pop() (int32, error) {
  if stack.size == 0 {
    return 0, fmt.Errorf("Stack is empty.")
  }
  stack.size -= 1
  popingNum := stack.elements[stack.size]
  if stack.size == 0 {
    stack.elements = []int32{}
    stack.maxValueSoFar = []int32{}
  } else {
    stack.elements = stack.elements[:stack.size]
    stack.maxValueSoFar = stack.maxValueSoFar[:stack.size]
  }
  return popingNum, nil
}

func (stack *Stack1) Max() (int32, error) {
  if stack.size == 0 {
    return 0, fmt.Errorf("Stack is empty.")
  }
  return stack.maxValueSoFar[stack.size-1], nil
}
