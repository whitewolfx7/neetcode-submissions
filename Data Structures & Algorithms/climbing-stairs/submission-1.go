func climbStairs(n int) int {

    if n <= 2 {
        return n
    }

    ways := make([]int, n + 1)
    ways[1], ways[2] = 1, 2
    for i := 3 ; i <= n; i++ {
        ways[i] = ways[i-1] + ways[i-2]
    }

    return ways[n]

}
