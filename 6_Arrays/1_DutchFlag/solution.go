package DutchFlag

import "fmt"

// In go, slices are "descriptors of arrays"; thus if we modify any element in
// the slice numbers, since the pointer is the same, the caller's slice content
// will change too. This automatically satisfy the requirement of O(1) space
// complexity.
func dutchFlag(numbers []int, index int) error {
  if index >= len(numbers) || index < 0 {
    return fmt.Errorf("index %d is outside of the range of numbers.", index)
  }
  pivot := numbers[index]
  small := 0
  large := len(numbers) - 1 // The last element in the slice.
  equal := 0
  var swap = func(i, j int) {
    temp := numbers[i]
    numbers[i] = numbers[j]
    numbers[j] = temp
  }
  for equal <= large {
    if numbers[equal] < pivot {
      swap(equal, small)
      equal += 1
      small += 1
    } else if numbers[equal] > pivot {
      swap(equal, large)
      large -= 1
    } else { // numbers[equal] == pivot
      equal += 1
    }
  }
  return nil
}
