package main

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	res := -32768
	f_n := -1
	// {-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 0; i < len(nums); i++ {
		f_n = Max(nums[i], f_n+nums[i])
		res = Max(f_n, res)
	}
	return res
}

func Max(a int, b int) int {

	if a >= b {
		return a
	}
	return b
}

// @lc code=end
