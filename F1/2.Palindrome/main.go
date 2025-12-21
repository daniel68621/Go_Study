package main

import "fmt"

// isPalindrome 判断整数是否为回文数
func isPalindrome(x int) bool {
	// 特殊情况1：负数不是回文数（负号无法对称）
	if x < 0 {
		return false
	}
    // 特殊情况2：0到9为回文数
    if 0<=x && x<10 {
		return true
	// 特殊情况3：非零且末尾为0的数不是回文数（0本身是回文数）
	if x != 0 && x%10 == 0 {
		return false
	}

	reversedHalf := 0 // 存储反转的后半部分数字
	for x > reversedHalf {
		// 取出x的最后一位，拼接到reversedHalf末尾
		reversedHalf = reversedHalf*10 + x%10
		// 去掉x的最后一位
		x = x / 10
	}

	// 两种情况：
	// 1. 偶数位（如1221）：x == reversedHalf
	// 2. 奇数位（如12321）：x == reversedHalf/10（去掉中间的3）
	return x == reversedHalf || x == reversedHalf/10
        
}

func main() {
	// 测试用例
	testCases := []int{121, -121, 10, 12321, 0, 1331, 123}
	for _, caseNum := range testCases {
		fmt.Printf("%d 是否是回文数：%v\n", caseNum, isPalindrome(caseNum))
	}
}

//  输出结果
// 121 是否是回文数：true
// -121 是否是回文数：false
// 10 是否是回文数：false
// 12321 是否是回文数：true
// 0 是否是回文数：true
// 1331 是否是回文数：true
// 123 是否是回文数：false
