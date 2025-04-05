# The isBadVersion API is already defined for you.
def isBadVersion(version: int) -> bool:
    if version >= 2:
        return True
    return False

class Solution:
    def firstBadVersion(self, n: int) -> int:
        l, r = 0, n
        while r - l > 1:
            # Avoid overflow (r+l)/2
            i = int(l + (r - l) / 2)
            if isBadVersion(i):
                r = i
            else:
                l = i
        # After narrowing down, check which of the remaining two is the bad version
        if isBadVersion(l):
            return l
        else:
            return r

print(Solution().firstBadVersion(10))
