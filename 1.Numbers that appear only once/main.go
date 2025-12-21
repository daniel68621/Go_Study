package main

import (
	"fmt"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

// 方法一：
// 解题思路
// 统计次数：创建一个 map[int]int 类型的哈希表，键是数组中的数字，值是该数字出现的次数；
// 遍历数组：用 for 循环遍历输入数组，每遇到一个数字，就将其在 map 中对应的次数加 1；
// 查找目标：再次用 for 循环遍历 map，找到值为 1 的键，这个键就是只出现一次的数字：
// 缺点：空间复杂度为 O(n)（需要存储所有数字的次数）

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
	// 按题意总有一个答案，这里做兜底返回
	return 0
}

// 方法二：使用异或运算

func singleNumber1(nums []int) int {
	// 初始化结果为0
	result := 0
	// 遍历数组中的每个数字，依次进行异或运算
	for _, num := range nums {
		result ^= num
	}
	// 最终结果就是只出现一次的数字
	return result
}

func main() {
	//1、只出现一次的数字

	// 测试用例1
	testCase1 := []int{2, 2, 1}
	fmt.Printf("数组 %v 中只出现一次的数字是：%d\n", testCase1, singleNumber(testCase1)) // 输出:数组 [2 2 1] 中只出现一次的数字是：1

	// 测试用例2
	testCase2 := []int{4, 1, 2, 1, 2}
	fmt.Printf("数组 %v 中只出现一次的数字是：%d\n", testCase2, singleNumber1(testCase2)) // 输出：数组 [4 1 2 1 2] 中只出现一次的数字是：4

}
