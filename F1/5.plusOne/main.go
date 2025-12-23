package main

import "fmt"

// plusOne 实现数组表示的数字加一功能
func plusOne(digits []int) []int {
	n := len(digits)
	// 从最后一位开始遍历
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++
		// 取余，判断是否进位：如果是10则变为0，否则无进位直接退出循环
		digits[i] %= 10
		if digits[i] != 0 {
			// 无进位，直接返回
			return digits
		}
	}
	// 所有位都进位（如999→1000），需要新建数组
	result := make([]int, n+1)
	result[0] = 1 // 首位设为1，其余默认0
	return result
}

func main() {
	// 测试用例1：普通情况 [1,2,3] → [1,2,4]
	test1 := []int{1, 2, 3}
	fmt.Println(plusOne(test1)) // 输出 [1 2 4]

	// 测试用例2：全9情况 [9,9] → [1,0,0]
	test2 := []int{9, 9}
	fmt.Println(plusOne(test2)) // 输出 [1 0 0]

	// 测试用例3：末尾无进位 [4,3,2,1] → [4,3,2,2]
	test3 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(test3)) // 输出 [4 3 2 2]
}
