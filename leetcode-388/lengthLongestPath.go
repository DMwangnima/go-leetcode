package leetcode_388

// 这种方法没有使用字符串API，一个字节一个字节解析，需要维护较多状态
func lengthLongestPath(input string) int {
	input = input + "\n"
	var start int
	var maxLength, curLength int
	var fileFlag bool
	curLevel := 1
	dirs := []string{""}
	for i := 0; i < len(input); i++ {
		if input[i] == '.' {
			fileFlag = true
		}
		if input[i] == '\n' {
			if fileFlag {
				curLength = 0
				for i := 0; i < curLevel; i++ {
					curLength += len(dirs[i])
				}
				curLength += len(input[start:i])
				curLength += curLevel - 1
				if curLength > maxLength {
					maxLength = curLength
				}
				start = i + 1
				curLevel = 1
				fileFlag = false
			} else {
				if len(dirs) < curLevel + 1 {
					dirs = append(dirs, input[start:i])
				} else {
					dirs[curLevel] = input[start:i]
				}
				curLevel = 1
				start = i + 1
			}
		}
		if input[i] == '\t' {
			start = i + 1
			curLevel += 1
		}
	}
	return maxLength
}
