# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
def guess(num: int) -> int:
    if num > 6:
        return -1
    if num < 6:
        return 1
    return 0

class Solution:
    def guessNumber(self, n: int) -> int:
        l, r = 0, n
        while l < r:
            # Avoid overflow (r+l)/2
            i = int(l + (r - l) / 2)
            if guess(i) == 0:
                return i
            if guess(i) == 1:
                l = i+1
            else:
                r = i
        return l

print("2 =", Solution().guessNumber(2))
print("6 =", Solution().guessNumber(10))
print("1 =", Solution().guessNumber(1))

