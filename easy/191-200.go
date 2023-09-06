package easy

//190二进制中1的个数
func HammingWeight(num uint32) int {
	//每一位的位运算
	count := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			count++
		}
	}

	count = 0
	//将最低一位为1的置为0，然后++
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}
