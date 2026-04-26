func rob(nums []int) int {
	var money []int
	var n int = len(nums)
	if n == 1{
		return nums[0]
	}
    money = make([]int, n)
	money[0], money[1] = nums[0], max(nums[0], nums[1])
	for i:=2; i<n;i++ {
		money[i] = max(nums[i] + money[i-2], money[i-1])
	}

	return money[n-1] 
}
