package p1218

//https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/

//给你一个整数数组arr和一个整数difference，请你找出并返回 arr中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。
//子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。

func longestSubsequence(arr []int, difference int) (ans int) {
	dp := map[int]int{}
	for _, num := range arr {
		dp[num] = dp[num-difference] + 1
		if dp[num] > ans {
			ans = dp[num]
		}
	}
	return
}
