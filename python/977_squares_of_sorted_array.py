from typing import List

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        result: List[int] = []
        # Use a two-pointer algoritm.
        # Start from both ends of the sorted input list,
        # as the largest absolute values might be at either end.
        l, r = 0, len(nums) - 1
        for i in range(0, len(nums)):
            if abs(nums[l]) > abs(nums[r]):
                result.insert(0, nums[l]*nums[l])
                l = l + 1
            else:
                result.insert(0, nums[r]*nums[r])
                r = r - 1

        return result


s = Solution()
print("[0, 1, 9, 16, 100] =", s.sortedSquares([-4, -1, 0 ,3, 10]))
print("[4, 9, 9, 49, 121] =", s.sortedSquares([-7,-3,2,3,11]))
