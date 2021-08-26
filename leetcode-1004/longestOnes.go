package leetcode_1004

// 使用滑动窗口
func longestOnes(nums []int, k int) int {
	// [start, end) => 0用1填充后全1子数组
	var start, end, res, curLength int
	var queue []int
	for end < len(nums) {
		if nums[end] == 0 {
			queue = append(queue, end)
			if k == 0 {
				curLength = end - start
				res = max(res, curLength)
				start = queue[0] + 1
				queue = queue[1:]
			} else {
				k -= 1
			}
		}
		end += 1
	}
	res = max(res, end - start)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
