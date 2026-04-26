func rob(nums []int) int {
   
	if len(nums) == 1{
		return nums[0]
	}

	find := func(nums []int) int{

		if len(nums) == 1 {
        	return nums[0]
    	}

		money_2, money_1 := nums[0], max(nums[0], nums[1])

		for i:=2; i<len(nums); i++{
			money := max(nums[i] + money_2, money_1)
			money_2, money_1 = money_1, money
		}

		return money_1
	}

	return max(find(nums[1:]), find(nums[:len(nums)-1]))
}
