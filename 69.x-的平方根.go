package main

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	//求6的平方根

	a := x / 2
	diff := 0
	b := 0
	for true {
		b = (a + x/a) / 2.0
		diff = a - b
		a = b
		if diff < 0 {
			break
		}
	}
	return a

}

// @lc code=end
