package MaxDifference

import "math"

func findMaxDifference(numbers []int64) int64 {
  currentMax := int64(0)
  currentLow := int64(math.MaxInt64)
  for _, num := range numbers {
    if num < currentLow {
      // There is no chance that this would be larger.
      currentLow = num
      continue
    }
    tempDistance := num - currentLow
    if tempDistance > currentMax {
      currentMax = tempDistance
    }
  }
  return currentMax
}
