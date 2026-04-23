class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        i, j = 0, len(nums) - 1
        map = {}

        for i, num in enumerate(nums):
            if (target - num) in map:
                return [map[target-num], i]
            map[num] = i
            
        return [0, 0]