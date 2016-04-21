package Spiral

import "testing"


func TestPrintSpiralOdd(tt *testing.T) {
  inputs := [][]int{
    {11, 12, 13, 14, 15},
    {21, 22, 23, 24, 25},
    {31, 32, 33, 34, 35},
    {41, 42, 43, 44, 45},
    {51, 52, 53, 54, 55}}
  PrintSpiral(inputs, 5)
}

func TestPrintSpiralEven(tt *testing.T) {
  inputs := [][]int{
    {11, 12, 13, 14},
    {21, 22, 23, 24},
    {31, 32, 33, 34},
    {41, 42, 43, 44}}
  PrintSpiral(inputs, 4)
}

func TestPrintSpiralEmpty(tt *testing.T) {
  inputs := [][]int{}
  PrintSpiral(inputs, 0)
}