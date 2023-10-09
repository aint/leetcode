from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        count = 0
        for i in range(1, len(nums)):
            s = nums[i-1] - nums[i]
            if s >= 0:
                s += 1
                nums[i] = nums[i] + s
                count += s

        return count


s = Solution()
nums = [1,1,1]
nums = [1,5,2,4,1]
print(s.minOperations(nums))
