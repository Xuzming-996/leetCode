package _1_40

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//32 从上到下打印二叉树
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

//39 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	recordMap := make(map[int]int)
	ret := 0
	for _, v := range nums {
		recordMap[v]++
		if recordMap[v]*2 > len(nums) {
			return v
		}
	}
	return ret
}

func QuickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	key := partSort(arr, left, right)
	QuickSort(arr, left, key-1)
	QuickSort(arr, key+1, right)

	return arr
}

func partSort(arr []int, left, right int) int {
	key := left
	for left < right {
		for left < right && arr[right] >= arr[key] {
			right--
		}
		for left < right && arr[left] <= arr[key] {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}

	arr[key], arr[left] = arr[left], arr[key]
	return left
}

//40 最小的K个数
func getLeastNumbers(arr []int, k int) []int {
	a := QuickSort(arr, 0, len(arr)-1)
	a = a[:k]
	return a
}
