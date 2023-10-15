from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        result, left, right = 0, 0, len(height) - 1
        while left < right:
            r = (right - left) * min(height[left], height[right])
            if r > result:
                result = r

            if height[left] > height[right]:
                right -= 1
            else:
                left += 1

        return result


s = Solution()
print(s.maxArea([1,8,6,2,5,4,8,3,7]))
