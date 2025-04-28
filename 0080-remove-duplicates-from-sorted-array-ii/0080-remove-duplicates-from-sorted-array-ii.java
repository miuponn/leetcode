class Solution {
    public int removeDuplicates(int[] nums) {
        int l = 2;                              // slow pointer: tracks position to place next allowed element

        for (int r = 2; r < nums.length; r++){  // fast pointer: scans each element (starting from third element)
            if (nums[r] != nums[l - 2]) {       // compare with element two steps behind
                nums[l] = nums[r];              // place current element at position of slow pointer
                l++;                            
            }
        }       
        return l;                               // return number of allowed elements
    }
}