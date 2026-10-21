class Solution {
    public void moveZeroes(int[] nums) {
        int l = 0;                                  // init left pointer

        for (int r = 0; r < nums.length; r++) {     // iterate over arr, increment right pointer
            if (nums[r] != 0) {                     // if right num is not 0
                int temp = nums[l];                 // swap left and right
                nums[l] = nums[r];
                nums[r] = temp;
                l++;                                // increment left pointer
            }
        }
    }
}