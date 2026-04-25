func minCostClimbingStairs(cost []int) int {

    price := make([]int, len(cost))
	price[0] = cost[0]
	price[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		price[i] = cost[i] +  min(price[i-1], price[i-2])
	}

	return min(price[len(cost)-1], price[len(cost)-2])
}