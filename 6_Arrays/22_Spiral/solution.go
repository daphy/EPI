package Spiral

import "fmt"

func PrintSpiral(matrix [][]int, length int) {
  var printOneCircle func(offset int)
  printOneCircle = func(offset int) {
    // Print the top line left -> right.
    for i := offset; i < length-offset; i++ {
      fmt.Printf("%d, ", matrix[offset][i])
    }
    // Print the right line top -> down.
    // The top-most element is printed in the previous loop,
    // so here we'll start from offset + 1.
    for i := offset+1; i < length-offset; i++ {
      fmt.Printf("%d, ", matrix[i][length-offset-1])
    }
    // Print the down line right -> left.
    // The right-most element is printed in the previous loop already.
    for i := length-offset-1-1; i >=offset ; i-- {
      fmt.Printf("%d, ", matrix[length-offset-1][i])
    }
    // Print the left line down -> top.
    // The lowest element is printed already, and so is the top most one.
    for i := length - offset - 1 - 1; i >= offset + 1 ; i-- {
      fmt.Printf("%d, ", matrix[i][offset])
    }
  }

  for i := 0 ; i < length / 2 ; i++ {
    printOneCircle(i)
  }
  // If length is an odd number, the middle one would not have been visited.
  // Print it out now.
  if length % 2 == 1 {
    fmt.Printf("%d, ", matrix[length/2][length/2])
  }
  fmt.Printf("\nEnd\n")
}