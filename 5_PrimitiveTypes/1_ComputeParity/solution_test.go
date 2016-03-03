package ComputeParity

import (
  "testing"
)

func TestComputeParityEven(tt *testing.T) {
  if computeParity(123) { // 1111011 => false
    tt.Fatal("123 should return false.")
  }
}

func TestComputeParityOdd(tt *testing.T) {
  if !computeParity(122) { // 1111010 => true
    tt.Fatal("122 should return true.")
  }
}

func TestComputeParityZero(tt *testing.T) {
  if computeParity(0) { // 0 => false
    tt.Fatal("0 should return false.")
  }
}
