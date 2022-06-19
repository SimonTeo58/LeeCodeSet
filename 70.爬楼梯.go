package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// int BottomUpCutRod(const std::vector<int> &p, const int &n)
// {
//     std::vector<int> r{0};
//     for (auto j = 1; j <= n; ++j)
//     {
//         int q = INT8_MIN;
//         for (auto i = 1; i <= j; ++i)
//         {
//             q = std::max(q, p.at(i) + r.at(j - i));
//         }
//         r.push_back(q);
//     }
//     return r.at(n);
// }
// 1 1 2 3 5 8
//   1 2 3 4 5
func climbStairs(n int) int {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := []int{1, 2}
	for i := 3; i <= n; i++ {
		Sum(dp)
	}

	return dp[1]
}
func Sum(dp []int) {
	pre := dp[0]
	dp[0] = dp[1]
	dp[1] = pre + dp[1]
}

// @lc code=end
