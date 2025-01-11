from typing import List

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        n = len(nums)
        s1 = (n*(n+1))/2
        s2 = 0
        for i in nums:
            s2 += i
        return int(s1-s2)