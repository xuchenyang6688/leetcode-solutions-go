package dynamic_programming

func ClimbStairs(n int) int {
	if n<=2 {
		return n
	}
	 // Fibonacci-like sequence: ways[i] = ways[i-1] + ways[i-2]
    pre1, pre2 := 1, 2 // pre1 = ways for n=1, pre2 = ways for n=2
	for i:=3; i<=n; i++ {
		current := pre1 + pre2 //  Use := to declare and assign
		pre1 = pre2
		pre2 = current
		
	}
	return pre2
}