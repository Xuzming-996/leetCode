package _1_50

//42连续子数组最大的和
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//50第一个只出现一次的字符
func firstUniqChar(s string) byte {
	list := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		list[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		count, ok := list[s[i]]
		if ok && count == 1 {
			return s[i]
		}
	}
	return ' '
}
