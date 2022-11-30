package _0_20

import (
	"math"
	"strconv"
)

// PrintNumbers 打印从1到最大的n位数
func PrintNumbers(n int) []int {
	var (
		nums, buf = make([]byte, 10), make([]byte, n)
		res       = make([]int, 0, int(math.Pow10(n)))
		dfs       func(x, end int)
	)

	for i := 0; i < 10; i++ {
		nums[i] = '0' + byte(i)
	}

	dfs = func(x, end int) {
		if x == end {
			num, _ := strconv.Atoi(string(buf[:end]))
			res = append(res, num)
			return
		}

		var i int
		if x == 0 { // 个位 i从1到9
			i = 1
		}

		for ; i < 10; i++ {
			buf[x] = nums[i]
			dfs(x+1, end)
		}
	}

	for i := 1; i <= n; i++ {
		dfs(0, i)
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//18.删除链表的节点
func DeleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	p := head
	cur := head.Next
	for cur != nil && cur.Val != val {
		p = cur
		cur = cur.Next
	}
	if cur != nil {
		p.Next = cur.Next
	}
	return head
}
