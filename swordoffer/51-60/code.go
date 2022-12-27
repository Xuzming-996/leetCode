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

//53-2 //0~n-1中缺失的数字（二分法，先从中间起判断是否已经不相等，如果不相等，就往前二分，否则往后）
func missingNumber(nums []int) int {
	l, r := -1, len(nums)
	for l+1 != r {
		m := (l + r) >> 1
		if m != nums[m] {
			r = m
		} else {
			l = m
		}
	}
	return l + 1
}

//54. 二叉搜索树的第K大节点
func kthLargest(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	var res = -1
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Left)
	}
	dfs(root)
	return res
}

//55 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max1(height(root.Left), height(root.Right)) + 1
}

func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
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

// 57-2 和为S的连续正数序列
func FindContinuousSequence(target int) [][]int {
	i, j, sum := 1, 1, 0
	ret := make([][]int, 0)

	target1 := target >> 1
	for i <= target1 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			tmp := make([]int, 0, j-i)
			for m := i; m < j; m++ {
				tmp = append(tmp, m)
			}
			ret = append(ret, tmp)
			sum -= i
			i++
		}
	}
	return ret
}
