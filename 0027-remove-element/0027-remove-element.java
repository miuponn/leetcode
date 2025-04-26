class Solution {
    public int removeElement(int[] nums, int val) {
        int k = 0;                                    // counter for number of elements not equal to val

        for (int i = 0; i < nums.length; i++){        // iterate through the array
            if (nums[i] != val){                      // if current element is not equal to val
                nums[k] = nums[i];                    // overwrite element at index k with current element
                k++;                                  // increment count of valid elements
            } 
        }
        return k;                                     // return number of elements not equal to val
    }
}
