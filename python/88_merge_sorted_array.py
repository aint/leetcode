from typing import List

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        # Use two pointers for the both arrays and a pointer k initialized to m+n-1, which will be used
        # to keep track of the position in nums1 where we will be placing the larger element.
        # Start iterating from the end of the both arrays, and compare the respective elements. Place the larger element in nums1 at position k,
        # and decrement the corresponding pointer p1 or p2 accordingly. Continue doing this until we have iterated through all the elements in nums2.
        # If there are still elements left in nums1, we don't need to do anything because they are already in their correct place.
        p1, p2 = m-1, n-1
        k = m + n - 1
        while p2 >= 0:
            if p1 >= 0 and nums1[p1] > nums2[p2]:
                nums1[k] = nums1[p1]
                p1 -= 1
            else:
                nums1[k] = nums2[p2]
                p2 -= 1
            k -=1

        print(nums1)


s = Solution()
s.merge([1,3,5,0,0,0], 3, [1,2,3], 3)
