#include <iostream>

class Solution {
public:
    int reverse(int x) {
        int result = 0;
        do {
            int remainder = x % 10;
            if (result > INT_MAX / 10) return 0;
            if (result < INT_MIN / 10) return 0;
            result = result * 10 + remainder;
            x /= 10;


        } while (x != 0);
        return result;
    }
};
