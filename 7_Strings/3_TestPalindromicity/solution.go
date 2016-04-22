package TestPalindromicity

func TestPalindromicity(input string) bool {

  // Define a function which returns the index of next English letter
  // from the startingIndex, as wella as the lowercase of the letter;
  // if the reversedOrder is set to true, it'll traversed backwards.
  // Returns -1 if it's out of range.
  getNextLetterIndex := func(startingIndex int, reversedOrder bool) (
    int, byte) {
    next := 1
    maxIndex := len(input) - 1
    if reversedOrder {
      next = -1
      maxIndex = 0
    }
    for index := startingIndex; index != maxIndex; index += next {
      if (input[index] >= 'a' && input[index] <= 'z') {
        return index, input[index]
      }
      if input[index] >= 'A' && input[index] <= 'Z' {
        return index, input[index] - 'A' + 'a'
      }
    }
    return -1, 0
  }

  // Start testing each character.
  headIndex := 0
  headLetter := byte(0)
  tailIndex := len(input) - 1
  tailLetter := byte(0)
  for ; headIndex <= tailIndex; {
    headIndex, headLetter = getNextLetterIndex(headIndex, false)
    tailIndex, tailLetter = getNextLetterIndex(tailIndex, true)
    if headIndex < 0 && tailIndex < 0 { // Nothing to test anymore...
      return true
    }
    if headIndex == tailIndex { // Both point at the same letter; done testing.
      return true
    }
    // If only one index is out of range, that letter would be 0,
    // but the other letter would NOT be 0; in this case we return false
    // as well.
    if headLetter != tailLetter {
      return false
    }
    headIndex += 1
    tailIndex -= 1
  }
  return true
}

