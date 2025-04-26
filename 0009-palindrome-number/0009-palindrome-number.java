class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0 || (x % 10 == 0 && x != 0)){          // negative ints or ints ending with 0 (but not 0 itself) are not palindromes
            return false;
        }

        int reversed = 0;                               // store reversed digits of the the second half of x
        while (x > reversed) {                          // build reversed half until it matches or exceeds (for odd digits) remaining x
            reversed = (reversed * 10) + (x % 10);      // append last digit of x to reversed
            x /= 10;                                    // remove last digit from x
        }

        return x == reversed || x == reversed / 10;     // for even length: x == reversed; for odd length: x == reversed/10 (skip middle digit)
    }
}