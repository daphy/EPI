package Mnemonics

var letters [10][]string = [10][]string {
  {"0"},
  {"1"},
  {"a", "b", "c"},
  {"d", "e", "f"},
  {"g", "h", "i"},
  {"j", "k", "l"},
  {"m", "n", "o"},
  {"p", "q", "r", "s"},
  {"t", "u", "v"},
  {"w", "x", "y", "z"},
}

func Mnemonics(phNum string) []string {
  if phNum == "" {
    return []string{}
  }

  var GetSubStrings func(int) []string
  GetSubStrings = func(index int) (subStrings []string) {
    if index == len(phNum) - 1 { // Last number in the phNum.
      return letters[phNum[index]-'0']
    }
    tailSubStrings := GetSubStrings(index + 1)
    for _, letter := range letters[phNum[index]-'0'] {
      for _, tailSubStr := range tailSubStrings {
        subStrings = append(subStrings, letter + tailSubStr)
      }
    }
    return
  }
  return GetSubStrings(0)
  // for _, number := range phNum { // outer most loop.

  // }
  // return []string{}
}