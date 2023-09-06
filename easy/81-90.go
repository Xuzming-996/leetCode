package easy

// Merge 88.合并两个有序数组
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	var list []int
	p1, p2 := 0, 0
	for {
		if p1 == m {
			list = append(list, nums2[p2:]...)
			break
		}
		if p1 == n {
			list = append(list, nums1[p1:]...)
			break
		}

		if nums1[p1] <= nums2[p2] {
			list = append(list, nums1[p1])
			p1++
		} else {
			list = append(list, nums2[p2])
			p2++
		}
	}
	nums1 = list
	copy(nums1, list)
	return nums1
}
