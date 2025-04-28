class Solution {
    public int majorityElement(int[] nums) {
        int m = 0;                                 // majority candidate element
        int c = 0;                                 // counter for current candidates

        for (int i = 0; i < nums.length; i++){
            if (c == 0) {                          // if count drops to 0, reassign candidate
                m = nums[i];
            }
            if (nums[i] == m) {                    // if current element matches, increment counter
                c++;
            } else {                               // otherwise, decrement count
                c--;
            }
        }
        return m;
    }
}