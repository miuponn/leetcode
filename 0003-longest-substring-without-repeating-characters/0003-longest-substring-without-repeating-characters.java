class Solution {
    public int lengthOfLongestSubstring(String s) {
        Set<Character> substr = new HashSet<>();        // dynamic sliding window (set of current chars)
        int res = 0, l = 0;                             // initialize result and left pointer

        for (int r = 0; r < s.length(); r++){           // iterate over string elements with right pointer
            char ch = s.charAt(r);                      // current char at right pointer
            while (substr.contains(ch)) {               // if duplicate found, shrink window from left
                substr.remove(s.charAt(l));             // remove leftmost char
                l++;                                    // move left pointer forward
            }
            substr.add(ch);                             // add current char to the window
            res = Math.max(res, r - l + 1);             // update result with max window size so far
        }
        return res;
    }
}