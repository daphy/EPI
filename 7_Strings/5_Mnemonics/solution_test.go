package Mnemonics

import "testing"

func TestMnemonicsFunc(tt *testing.T) {
  tt.Log(Mnemonics("123"))
  tt.Log(Mnemonics("111111"))
  tt.Log(Mnemonics("8888"))
  tt.Log(Mnemonics(""))
  tt.Log(Mnemonics("4"))
  tt.Log(Mnemonics("0"))
}