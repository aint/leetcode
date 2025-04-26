class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = 'aeiou'
        s = list(s)
        l = 0
        r = len(s) - 1
        while l < r:
            if s[l].lower() not in vowels:
                l += 1
                continue
            if s[r].lower() not in vowels:
                r -= 1
                continue
            tmp = s[l]
            s[l] = s[r]
            s[r] = tmp
            r -= 1
            l += 1
        return ''.join(s)
