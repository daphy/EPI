package StrIntConvert

import (
  "testing"
)

type testcase struct {
  num int64
  str string
}

var validTestcases = []testcase{
  {12345, "12345"},
  {0, "0"},
  {-12345, "-12345"},
  {20, "20"},
}

var invalidStrings = []string{
  "123abc", "102.3", "", " ", "abc", "--", "-", "0002",
}

func TestAToINoError(tt *testing.T) {
  for _, tc := range validTestcases {
    result := AToI(tc.num)
    if result != tc.str {
      tt.Errorf("AToI(%d) returns \"%s\", expected \"%s\"",
        tc.num, result, tc.str)
    }
  }
}

func TestIToANoError(tt *testing.T) {
  for _, tc := range validTestcases {
    result, err := IToA(tc.str)
    if err != nil {
      tt.Errorf("IToA(\"%s\") should be valid but returns error %s.",
        tc.str, err)
      continue
    }
    if result != tc.num {
      tt.Errorf("IToA(\"%s\") returns %d, expected %d.", tc.str,
        result, tc.num)
      continue
    }
  }
}

func TestIToASpecial(tt *testing.T) {
  tc := "-0"
  exp := int64(0)
  result, err := IToA(tc)
  if err != nil {
    tt.Errorf("IToA(\"%s\") should be valid but returns error %s.",
      tc, err)
  }
  if result != exp {
    tt.Errorf("IToA(\"%s\") returns %d, expected %d.", tc, result, exp)

  }
}

func TestIToAError(tt *testing.T) {
  for _, tc := range invalidStrings {
    result, err := IToA(tc)
    if err == nil {
      tt.Errorf("IToA(\"%s\") should return error; got %d instead.", tc, result)
      continue
    }
  }
}
