class Solution {
    public int removeDuplicates(int[] nums) {
        int l = 0;                              // slow pointer: tracks position of last unique element

        for (int r = 1; r < nums.length; r++){  // fast pointer: scans each element
            if (nums[r] != nums[r - 1]) {           // if a new unique element is found, move slow pointer forward
                l++;
                nums[l] = nums[r];              // overwrite with the new unique element
            }
        }       
        return l + 1;                           // return number of unique elements
    }
}