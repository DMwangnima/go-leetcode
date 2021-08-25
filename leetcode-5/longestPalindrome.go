package leetcode_5

// brutal force，找出可能的所有回文字符串
func longestPalindrome(s string) string {
	var res string
	length := len(s)
	for i := range s {
		if len(res) >= length-i {
			break
		}
		for j := i; j < length; j++ {
			tmp := s[i : j+1]
			if isPalindrome(tmp) && j+1-i > len(res) {
				res = tmp
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head += 1
		tail -= 1
	}
	return true
}

// 开始试着使用递归的思路在做，但结果是错误的
func longestPalindrome1(s string) string {
	res, _ := helper(s)
	return res
}

func helper(s string) (string, int) {
	length := len(s)
	if length == 0 || length == 1 {
		return s, length
	}
	if s[0] == s[length-1] {
		_, count := helper(s[1 : length-1])
		if count == length-2 {
			return s, length
		}
	}
	leftStr, leftCount := helper(s[0 : length-1])
	rightStr, rightCount := helper(s[1:length])
	if leftCount > rightCount {
		return leftStr, leftCount
	}
	return rightStr, rightCount
}

// todo
// 动态规划
func longestPalindrome2(s string) string {

}

// todo
// 中心扩散
func longestPalindrome3(s string) string {

}