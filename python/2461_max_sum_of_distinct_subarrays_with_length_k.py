from typing import List

# 1. Store first k elements in the map
# 2. Move the window - add the new one to the map and remove the old one
# 3. The same to compute running sum
# 4. Pick the max sum

class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        answer, cur_sum = 0, 0
        dic = {}
        for i in range(k): # store first k elements in the map
            n = nums[i]
            dic[n] = dic.get(n, 0) + 1
            cur_sum += n

        if len(dic) == k: # no duplicates in the window, so use the sum
            answer = cur_sum

        left, right = 0, k
        while right < len(nums):
            print(list(dic.keys()))

            # moving the window
            dic[nums[right]] = dic.get(nums[right], 0) + 1
            if dic[nums[left]] == 1: # removing the left element if no duplicates in the windows
                del dic[nums[left]]
            else:
                dic[nums[left]] -= 1 # there are duplicates, so subtracting 1 as we're moving to the right

            # compute running sum
            cur_sum += nums[right]
            cur_sum -= nums[left]

            left += 1
            right += 1

            if len(dic) == k: # no duplicates in the window, so use the current sum
                if cur_sum > answer:
                    answer = cur_sum

        return answer

s = Solution()
nums = [1,5,4,2,9,9,9]
print(">>>expected 15, actual answer", s.maximumSubarraySum(nums, 3))

nums = [1,1,1,7,8,9]
print(">>>expected 24, actual answer", s.maximumSubarraySum(nums, 3))

nums = [5,3,3,1,1]
print(">>>expected 0, actual answer", s.maximumSubarraySum(nums, 3))
