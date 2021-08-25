package leetcode_14

// 简单的递归，但因为栈空间的消耗，空间复杂度为O(N)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	headStr := strs[0]
	leftStr := longestCommonPrefix(strs[1:])
	headPtr, leftPtr := 0, 0
	for headPtr < len(headStr) && leftPtr < len(leftStr) {
		if headStr[headPtr] != leftStr[leftPtr] {
			break
		}
		headPtr += 1
		leftPtr += 1
	}
	if headPtr == 0 {
		return ""
	}
	return headStr[:headPtr]
}
