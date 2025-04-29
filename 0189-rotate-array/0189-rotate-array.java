class Solution {
    public void rotate(int[] nums, int k) {
        k = k % nums.length;                            // handle cases where k > array length

        reverse(nums, 0, nums.length - 1);              // 1.reverse entire array 2.reverse first k elements 3.reverse remaining n - k elements
        reverse(nums, 0, k - 1);
        reverse(nums, k, nums.length - 1);
    }

    public void reverse(int[] nums, int l, int r) {     // helper function: reverses elements between indicies l and r (inclusive)
        while (l < r) {                                 
            int temp = nums[l];                         
            nums[l] = nums[r];
            nums[r] = temp;

            l++;                    
            r--;
        }
    }
}