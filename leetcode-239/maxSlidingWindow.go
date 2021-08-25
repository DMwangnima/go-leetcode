package leetcode_239

import (
	"container/heap"
	"math"
)

// 暴力法，每次滑动窗口，就计算一次最大值
func maxSlidingWindow(nums []int, k int) []int {
    var res []int
    for i := k-1; i < len(nums); i++ {
    	res = append(res, findMax(nums[i+1-k:i+1]))
	}
	return res
}

func findMax(arr []int) int {
	max := math.MinInt
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}

// 对于这种数组滑动求最大值的问题，关键在于求得子数组的最大值
func maxSlidingWindow1(nums []int, k int) []int {
	heap.Push()
}
