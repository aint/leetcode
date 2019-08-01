class Solution {

    public String longestPalindrome(String s) {
        if (s.length() == 1) {
            return s;
        }
        String result = "";
        for (int i = 0; i < s.length(); i++) {
            for (int j = i + 1; j <= s.length(); j++) {
                String substring = s.substring(i, j);
                if (substring.length() <= result.length()) {
                    continue;
                }
                if (isPalindromic(substring)) {
                    if (substring.length() > result.length()) {
                        result = substring;
                    }
                }
            }
        }

        return result;
    }

    private boolean isPalindromic(String s) {
        for (int i = 0; i <= s.length() / 2; i++) {
            if (s.charAt(i) != s.charAt(s.length() - i - 1)) {
                return false;
            }
        }
        return true;
    }
}