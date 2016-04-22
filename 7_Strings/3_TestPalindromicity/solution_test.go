package TestPalindromicity

import "testing"

func TestTestPalindromicity(tt *testing.T) {
  type testCase struct {
    input string
    expected bool
  }
  testCases := []testCase {
    testCase{"", true},
    testCase{"A man, a plan, a canal, Panama.", true},
    testCase{"Able was I, ere I saw Elba!", true},
    testCase{"Ray a Ray", false},
    testCase{"......................", true},
    testCase{"z", true},
    testCase{"AaA", true},
  }
  for _, tc := range testCases {
    actual := TestPalindromicity(tc.input)
    if actual != tc.expected {
      tt.Errorf("\"%s\" returns %t, but expecting otherwise.",
        tc.input, actual)
    }
  }
}