package main

import (
	"fmt"
)

// longestCommonPrefix 查找字符串切片中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 边界条件：切片为空，直接返回空字符串
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串作为基准前缀
	prefix := strs[0]

	// 遍历基准前缀的每个字符（带索引）
	for i := 0; i < len(prefix); i++ {
		// 取出当前要比对的字符
		currentChar := prefix[i]
		// 遍历切片中剩余的所有字符串
		for j := 1; j < len(strs); j++ {
			// 终止条件：
			// 1. 当前字符串长度不足i（没有第i个字符）
			// 2. 当前字符串的第i个字符与基准字符不匹配
			if i >= len(strs[j]) || strs[j][i] != currentChar {
				// 返回前i个字符作为最长公共前缀
				return prefix[:i]
			}
		}
	}

	// 所有字符都匹配，基准字符串就是最长公共前缀
	return prefix
}

func main() {
	// 测试用例
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "常规情况（部分匹配）",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "无公共前缀",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "单个字符串",
			input:    []string{"apple"},
			expected: "apple",
		},
		{
			name:     "所有字符串完全相同",
			input:    []string{"go", "go", "go"},
			expected: "go",
		},
	}

	fmt.Print("打印测试用例：", testCases)

	// 执行测试并输出结果
	for _, tc := range testCases {
		result := longestCommonPrefix(tc.input)
		fmt.Printf("测试用例：%s\n输入：%v\n预期：%s\n实际：%s\n是否一致：%t\n\n",
			tc.name, tc.input, tc.expected, result, result == tc.expected)
	}
}
