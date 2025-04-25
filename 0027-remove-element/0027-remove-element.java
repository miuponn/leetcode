class Solution {
    public int removeElement(int[] nums, int val) {
        int k = 0;                                    // initialize counter k = elements that are not val

        for (int i = 0; i < nums.length; i++){        // iterate over nums array
            if (nums[i] != val){                     // if nums[i] != val
                nums[k] = nums[i];                    // override k index in nums with new val
                k++;                                  // increment k
            } 
        }
        return k;
    }
}