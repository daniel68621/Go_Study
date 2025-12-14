package main

import (
	"fmt"
	"sort"
	"strings"
)

// 只出现一次的数字
func singleNumber(nums []int) int {

	map1 := map[int]int{}

	for _, v := range nums {
		map1[v] = map1[v] + 1
	}
	for k, v := range map1 {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 回文数
func isPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	// 个位数是回文数
	if x < 10 {
		return true
	}

	//把整数转换为字符串
	str := fmt.Sprintf("%d", x)
	// 从字符串的两端开始比较字符
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// 有效的括号
func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pair[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

// 大整数加1
func plusOne(digits []int) []int {
	n := len(digits)
	// 从最低位开始加1
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] <= 9 {
			// 无进位，直接返回
			return digits
		}
		// 有进位，当前位置0，继续循环处理更高位
		digits[i] = 0
	}
	// 处理最高位仍有进位的情况，如 999 -> 1000
	return append([]int{1}, digits...)
}

// 原地删除有序数组重复元素，返回唯一元素个数
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 按起点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] { // 重叠，合并
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else { // 不重叠，直接追加
			res = append(res, intervals[i])
		}
	}
	return res
}

// 两数之和：返回和为 target 的两个元素下标
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := indexMap[target-v]; ok {
			return []int{j, i} // 找到答案
		}
		indexMap[v] = i // 记录值与下标
	}
	return nil // 题目保证有解，这里不会执行
}

func main() {
	//题目一
	arr := []int{1, 1, 2, 2, 3, 4, 3}
	number := singleNumber(arr)
	fmt.Println(number)

	//题目二
	palindrome := isPalindrome(12321)
	fmt.Println(palindrome)

	//题目三
	valid := isValid("()")
	fmt.Println(valid)

	//题目四
	strs := []string{"hello", "hero", "head"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)

	//题目五
	arr = plusOne([]int{1, 9})
	fmt.Println(arr)

	//题目六
	duplicates := removeDuplicates([]int{0, 0, 1, 2, 2, 3, 4, 4, 5, 5, 6})
	fmt.Println(duplicates)

	//题目七
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	fmt.Println(merged)

	//题目八
	nums := []int{7, 1, 2, 5}
	target := 6
	sum := twoSum(nums, target)
	fmt.Println(sum)
}
