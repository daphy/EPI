package FirstOccurance

// import "fmt"

func FindFirstOccurance(array []int, searchNum int) int {

  min := 0
  max := len(array) - 1
  foundIndex := -1
  for min <= max {
    if min == max {
      if array[min] == searchNum {
        foundIndex = min
        break
      }
    }
    middle := (min + max) / 2
    if array[middle] == searchNum {
      foundIndex = middle
      max = middle - 1
    } else if array[middle] > searchNum {
      max = middle - 1
    } else { // array[middle] < searchNum
      min = middle + 1
    }
  }
  return foundIndex
}
