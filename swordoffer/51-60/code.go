package _1_60

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//52 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headAmap := map[int]struct{}{}
	for p := headA; p.Next != nil; p = p.Next {
		headAmap[p.Val] = struct{}{}
	}

	for p := headB; p.Next != nil; p = p.Next {
		_, ok := headAmap[p.Val]
		if ok {
			return p
		}
	}
	return nil
}

//53 在排序数组中查找数字(使用二分法查找)
func Search(nums []int, target int) int {
	nums = []int{5, 7, 7, 8, 8, 10}
	left := searchInt(nums, target)
	if left == len(nums) || nums[left] != target {
		return 0
	}
	right := searchInt(nums, target+1) - 1
	return right - left + 1
}

func searchInt(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		h := (i + j) / 2
		if nums[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

//57 和为s的两个数字（双指针最优解）
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if target-nums[left] < nums[right] {
			right--
			continue
		} else if target-nums[left] > nums[right] {
			left++
			continue
		} else {
			return []int{nums[left], nums[right]}
		}
	}

	return []int{}
}
