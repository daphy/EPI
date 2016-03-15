package StackMax

type Stack interface {
  Push(num int32)
  Pop() (int32, error)
  Max() (int32, error)
}

func GetStack1() Stack {
  return &Stack1{}
}

func GetStack2() Stack {
  return &Stack2{}
}

// Force check if Stack1 and Stack2 implement Stack by
// declaring a dont-care _ variable as Stack, and assigning
// a type Stack1 pointer nil to it.
var _ Stack = (*Stack1)(nil) // &Stack1{}
var _ Stack = (*Stack2)(nil)
