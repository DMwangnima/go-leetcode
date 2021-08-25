package leetcode_316

import "strings"

// 错误解法，其实已经接近答案了，但是根据用例错误地选择了基准点，没有发现栈的模型
func removeDuplicateLetters(s string) string {
	var res string
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		ind := s[i] - 'a'
		switch arr[ind] {
		case 0:
			arr[ind] = -1
		default:
			arr[ind] = i
		}

	}
	start := 0
	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == -1 {
			for j := start; j < i; j++ {
				if arr[s[j]-'a'] == -2 {
					continue
				}
				if s[j] < s[i] || arr[s[j]-'a'] == j {
					res += string(s[j])
					arr[s[j]-'a'] = -2
				}
			}
			res += string(s[i])
			start = i + 1
		}
	}
	for start < len(s) {
		if arr[s[start]-'a'] >= 0 {
			res += string(s[start])
			arr[s[start]-'a'] = -2
		}
		start += 1
	}
	return res
}

func removeDuplicateLetters1(s string) string {
	lastIndexs := make([]int, 26)
	for i := range s {
		lastIndexs[s[i]-'a'] = i
	}
	exists := make([]bool, 26)
	var stack []byte
	for i := range s {
		if exists[s[i]-'a'] {
			continue
		}
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			ind := top - 'a'
			if top > s[i] && lastIndexs[ind] > i {
				stack = stack[:len(stack)-1]
                exists[ind] = false
                continue
			}
			break
		}
		stack = append(stack, s[i])
		exists[s[i]-'a'] = true
	}
	builder := &strings.Builder{}
	for i := 0; i < len(stack); i++ {
		builder.WriteByte(stack[i])
	}
	return builder.String()
}