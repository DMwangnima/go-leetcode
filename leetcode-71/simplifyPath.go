package leetcode_71

import "strings"

// 利用栈维持当前的目录情况
func simplifyPath(path string) string {
	sta := []string{}
	i := 0
	for i < len(path) {
		if path[i] == '/' {
			for i < len(path) && path[i] == '/' {
				i += 1
			}
		} else {
			tmp := &strings.Builder{}
			for i < len(path) && path[i] != '/' {
				tmp.WriteByte(path[i])
				i += 1
			}
			tmpStr := tmp.String()
			switch tmpStr {
			case ".":
			case "..":
				if len(sta) > 0 {
					sta = sta[:len(sta)-1]
				}
			default:
				sta = append(sta, tmpStr)
			}
		}
	}
	if len(sta) == 0 {
		return "/"
	}
	res := &strings.Builder{}
	for i := 0; i < len(sta); i++ {
		res.WriteByte('/')
		res.WriteString(sta[i])
	}
	return res.String()
}

// 如果要更简洁一点，应该使用string的API对/进行分割字符串，之后思路相同
// 这种方式内存占用比上一种方式更多，需要研究一下Split的实现
func simplifyPath1(path string) string {
	sta := []string{}
	strs := strings.Split(path, "/")
	for i := 0; i < len(strs); i++ {
		switch strs[i] {
		case "":
		case ".":
		case "..":
			if len(sta) > 0 {
				sta = sta[:len(sta)-1]
			}
		default:
			sta = append(sta, strs[i])
		}
	}
	if len(sta) == 0 {
		return "/"
	}
	res := &strings.Builder{}
	for i := 0; i < len(sta); i++ {
		res.WriteByte('/')
		res.WriteString(sta[i])
	}
	return res.String()
}