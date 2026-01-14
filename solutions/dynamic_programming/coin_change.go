package dynamic_programming

func CoinChange(coins []int, amount int) int {
	// Use amount+1 as "infinity" since max coins needed is amount (if coin=1 exists)
	maxVal := amount + 1
	// Initialize dp array
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = maxVal
	} 
	dp[0] = 0 // // Base case: 0 coins needed for amount 0

	for i:=1; i<=amount; i++ {
		for _, coin := range coins {
			if coin<=i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	// Check if solution exists
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
	   return a
	}
	return b   
}