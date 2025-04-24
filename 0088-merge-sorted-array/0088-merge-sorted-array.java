class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {

        int last = m + n - 1;                      // pointer to last empty element in nums1

        while (m > 0 && n > 0) {                   // iterate nums1 and nums2 starting from end of list 
            if (nums1[m - 1] > nums2[n - 1]){      // if element of nums1 > element of nums2
                nums1[last] = nums1[m - 1];        // assign element to position of pointer
                m -= 1;                            // decrement nums1
            } else {                               // if element of nums2 > element of nums1
                nums1[last] = nums2[n - 1];        // assign element to position of pointer
                n -= 1;                            // decrement nums2
            }  
            last -= 1;                             // decrement pointer
        }
        
        while (n > 0) {                            // if nums1 is done sorting but not nums2
            nums1[last] = nums2[n - 1];            // fill rest of nums1 with remaining nums2
            n -= 1;
            last -= 1;
        }
    }
}