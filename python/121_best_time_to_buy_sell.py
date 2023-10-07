from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        left, right, result = 0, 0, 0

        while right < len(prices):
            if prices[right] - prices[left] <= 0:
                left = right
                right += 1
                continue

            if prices[right] - prices[left] > result:
                result = prices[right] - prices[left]

            right += 1

        return result

s = Solution()
prices = [7,1,5,3,6,4]
print(s.maxProfit(prices))

# two pointer
