package StrIntConvert

import "fmt"

func AToI(num int64) string {
  if num == 0 {
    return "0"
  }
  var result []byte
  isNegative := false
  if num < 0 {
    isNegative = true
    num = 0 - num
  }
  for ; num > 0; num /= 10 {
    newbyte := byte(num%10) + '0'
    result = append([]byte{newbyte}, result...)
  }
  if isNegative {
    result = append([]byte("-"), result...)
  }
  return string(result)
}

func IToA(str string) (int64, error) {
  if len(str) == 0 {
    return 0, fmt.Errorf("Nothing to convert.")
  }
  var result int64
  startingIndex := 0
  isNegative := false
  // Check the first character since it's a special case;
  // if it's -, it means the result is negative,
  // iff the remaining str is a valid string.
  // if it's 0, it means the result is 0,
  // iff there is no other numbers after it. (ex: 01 is invalid.)
  if str[0] == '0' {
    if len(str) != 1 {
      return 0, fmt.Errorf("Invalid input %s", str)
    }
    return 0, nil
  }
  if str[0] == '-' {
    if len(str) == 1 {
      return 0, fmt.Errorf("Invalid input %s", str)
    }
    isNegative = true
    startingIndex = 1
  }
  for _, char := range str[startingIndex:] {
    // Check if it's a number.
    if char < '0' || char > '9' {
      return result, fmt.Errorf("Invalid input string %s",
        str)
    }
    result *= 10
    result += int64(char) - '0'
  }
  if isNegative {
    return 0 - result, nil
  }
  return result, nil
}
