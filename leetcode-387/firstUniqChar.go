package leetcode_387

func firstUniqChar(s string) int {
	m := map[byte]int
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
