package main

import "fmt"

// m-1<- m*n 0.0

func solution(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	fmt.Println(dp)
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	fmt.Println(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fmt.Println("i, j", i, j)
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
	// 1.边界初始
	// 2.dp公式
	// 返回结果

}

func main() {
	fmt.Println(solution(3, 2))
}
