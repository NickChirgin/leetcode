package graphs

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
			dp[i] = make([]int, len(matrix[0]))
	}
	copy(dp[0], matrix[0])

	for i := 1; i < len(dp); i++ {
			for j := 0; j < len(dp[0]); j++ {
					var minPathAbove int

					switch j {
					case 0:
							minPathAbove = min([]int{dp[i-1][j], dp[i-1][j+1]})
					case len(dp[0])-1:
							minPathAbove = min([]int{dp[i-1][j-1], dp[i-1][j]})
					default:
							minPathAbove = min([]int{dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]})
					}
					
					dp[i][j] = minPathAbove + matrix[i][j]
			}
	}

	return min(dp[len(dp)-1])
}

func min(a []int) int {
	curMin := a[0]
	for _, elem := range a {
			if elem < curMin {
					curMin = elem
			}
	}
	return curMin
}