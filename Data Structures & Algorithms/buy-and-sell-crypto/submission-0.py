class Solution:
    def maxProfit(self, prices: List[int]) -> int:

        curMin = prices[0]
        res = 0
        for i in range(1, len(prices)):
            if prices[i] < curMin:
                curMin = prices[i]
            else:
                res = max(prices[i] - curMin, res)

        return res
