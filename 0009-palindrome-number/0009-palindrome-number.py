class Solution:
    def isPalindrome(self, x: int) -> bool:
        # negative nums or nums ending with 0 (but not 0 itself) cannot be palindromes
        if x < 0 or (x % 10 == 0 and x != 0):
            return False
        
        # build reverted half with modulo and division
        reverted_half = 0
        while x > reverted_half:
            reverted_half = reverted_half * 10 + x % 10
            x //= 10
        
        # for even length: x == reverted_half
        # for odd length: x == reverted_half // 10 (skip middle digit)
        return x == reverted_half or x == reverted_half // 10
        