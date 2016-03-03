package ComputeParity

func computeParity(num uint64) bool {
  bParity := uint64(0)
  for num != 0 {
    bParity ^= num & 1
    num = num >> 1
  }
  if bParity == 1 {
    return true
  }
  return false
}
