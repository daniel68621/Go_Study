package main

import (
	"fmt"
)

// isValid 判断括号字符串是否有效
func isValid(s string) bool {
	// 定义右括号到左括号的映射，方便快速匹配
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 用切片模拟栈，存储左括号
	stack := []rune{}

	// 遍历字符串的每个字符
	for _, char := range s {
		// 如果是左括号（映射中不存在该key），压入栈
		if _, isRightBracket := bracketMap[char]; !isRightBracket {
			stack = append(stack, char)
		} else {
			// 是右括号，检查栈是否为空（无匹配的左括号）
			if len(stack) == 0 {
				return false
			}
			// 弹出栈顶元素，判断是否匹配
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != bracketMap[char] {
				return false
			}
		}
	}

	// 遍历结束后，栈必须为空（无未闭合的左括号）
	return len(stack) == 0
}

func main() {
	// 测试用例
	testCases := []struct {
		input  string
		expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},   // 空字符串视为有效
		{"(", false}, // 仅左括号
		{")", false}, // 仅右括号
	}

	// 执行测试
	for _, tc := range testCases {
		result := isValid(tc.input)
		fmt.Printf("输入: %q | 预期: %t | 实际: %t | %s\n",
			tc.input, tc.expect, result,
			map[bool]string{true: "✅ 正确", false: "❌错误"}[result == tc.expect])
	}
}
