package main

import (
	"fmt"
	base "testProject/swordoffer/31-40"
)

func main() {
	ret := []int{3, 5, 7, 1, 578, 54, 342, 3434}
	fmt.Println(base.QuickSort(ret, 0, len(ret)-1))
}
