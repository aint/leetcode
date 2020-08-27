#!/usr/bin/env python3
from typing import List

class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        result = []
        for i in range(n, n * 2):
            result.append(nums[i - n])
            result.append(nums[i])
        return result
